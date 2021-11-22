package usecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/zahraayafni/go-restful-api/config"
	"github.com/zahraayafni/go-restful-api/internal/customer"
	"github.com/zahraayafni/go-restful-api/internal/models"
	"github.com/zahraayafni/go-restful-api/pkg/httpErrors"
)

// Customer UseCase constructor
func InitCustomerUseCase(cfg *config.Config, customerRepo customer.CustomerRepository, customerRedisRepo customer.CustomerRedisRepository) customer.CustomerUseCase {
	return &CustomerUC{cfg: cfg, customerRepo: customerRepo, customerRedisRepo: customerRedisRepo}
}

// Insert customer
func (u *CustomerUC) InsertCustomer(ctx context.Context, customer *models.Customer) (uuid.UUID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CustomerUC.InsertCustomer")
	defer span.Finish()

	if err := PreValidateInsertCustomer(customer); err != nil {
		return uuid.UUID{}, httpErrors.NewBadRequestError(errors.WithMessage(err, "CustomerUC.InsertCustomer.PreValidateInsertCustomer"))
	}

	customerID, err := u.customerRepo.InsertCustomerDB(ctx, customer)
	if err != nil {
		return uuid.UUID{}, errors.WithStack(err)
	}

	return customerID, nil
}

func PreValidateInsertCustomer(customer *models.Customer) error {
	if customer.Email == "" || customer.Msisdn == "" {
		return errors.New("No email or msisdn data")
	}
	return nil
}

func PreValidateUpdateCustomer(customer *models.Customer) error {
	if customer.ID.String() == "" {
		return errors.New("ID is mandatory for data updates")
	}
	return nil
}

// Update customer data
func (u *CustomerUC) UpdateCustomer(ctx context.Context, customer *models.Customer) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CustomerUC.UpdateCustomer")
	defer span.Finish()

	if err := PreValidateUpdateCustomer(customer); err != nil {
		return httpErrors.NewBadRequestError(errors.WithMessage(err, "CustomerUC.UpdateCustomer.PreValidateUpdateCustomer"))
	}

	err := u.customerRepo.UpdateCustomerDB(ctx, customer)
	if err != nil {
		return err
	}

	if err = u.customerRedisRepo.ClearCustomersCache(ctx, []uuid.UUID{customer.ID}); err != nil {
		log.Printf("CustomerUC.UpdateCustomer.ClearCustomersCache: %v", err)
	}

	return nil
}

// Get customer by ids
func (u *CustomerUC) GetCustomerByIDs(ctx context.Context, customerIDs []uuid.UUID) ([]models.Customer, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CustomerUC.GetCustomerByIDs")
	defer span.Finish()

	result := make([]models.Customer, 0, len(customerIDs))
	// retrieve data from redis
	customersCache, err := u.customerRedisRepo.GetCustomerByIDsCache(ctx, customerIDs)
	if err != nil {
		// todo: add bypass cache nil error return from redis
		log.Printf("CustomerUC.GetCustomerByIDs.GetCustomerByIDsCache: %v", err)
	}

	// filter not cached data
	notCachedIDs := make([]uuid.UUID, 0, len(customerIDs))
	for _, cid := range customerIDs {
		if _, ok := customersCache[cid]; !ok {
			notCachedIDs = append(notCachedIDs, cid)
		}
	}

	// get data from db based on not cached ids
	var customersDB []models.Customer
	if len(notCachedIDs) > 0 {
		customersDB, err = u.customerRepo.GetCustomerByIDsDB(ctx, notCachedIDs)
		if err != nil {
			return result, errors.WithStack(err)
		}

		defer func() {
			if err = u.customerRedisRepo.SetCustomersCache(ctx, customersDB, cacheDuration); err != nil {
				log.Printf("CustomerUC.GetCustomerByIDs.SetCustomersCache: %s", err)
			}
		}()
	}

	// combine cache and db result
	result = append(result, customersDB...)
	for _, cc := range customersCache {
		result = append(result, cc)
	}

	return result, nil
}

// Delete customer result
func (u *CustomerUC) DeleteCustomer(ctx context.Context, customerID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CustomerUC.DeleteCustomer")
	defer span.Finish()

	//delete cache data
	if err := u.customerRedisRepo.ClearCustomersCache(ctx, []uuid.UUID{customerID}); err != nil {
		return errors.WithStack(err)
	}

	// delete db data
	err := u.customerRepo.DeleteCustomerByIDDB(ctx, customerID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
