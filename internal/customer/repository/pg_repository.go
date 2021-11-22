package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/zahraayafni/go-restful-api/internal/customer"
	"github.com/zahraayafni/go-restful-api/internal/models"
)

// Customer Repository
type customerRepository struct {
	db *sqlx.DB
}

// Customer repository constructor
func InitCustomerRepository(db *sqlx.DB) customer.CustomerRepository {
	return &customerRepository{db: db}
}

// Create customer
func (c *customerRepository) InsertCustomerDB(ctx context.Context, customer *models.Customer) (uuid.UUID, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "customerRepository.insertCustomerDB")
	defer span.Finish()

	var customerID uuid.UUID
	if err := c.db.QueryRowxContext(
		ctx,
		insertCustomerQuery,
		&customer.Name,
		&customer.Email,
		&customer.Msisdn,
		&customer.Address,
	).StructScan(&customerID); err != nil {
		return uuid.UUID{}, errors.Wrap(err, "customerRepository.InsertCustomerDB.QueryRowxContext")
	}

	return customerID, nil
}

// Update customer item
func (c *customerRepository) UpdateCustomerDB(ctx context.Context, customer *models.Customer) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "customerRepository.UpdateCustomerDB")
	defer span.Finish()

	if err := c.db.QueryRowxContext(
		ctx,
		updateCustomerQuery,
		&customer.Name,
		&customer.Email,
		&customer.Msisdn,
		&customer.Address,
		&customer.ID,
	); err != nil {
		return errors.Wrap(err.Err(), "customerRepository.UpdateCustomerDB.QueryRowxContext")
	}

	return nil
}

// Get customer by ids
func (c *customerRepository) GetCustomerByIDsDB(ctx context.Context, customerIDs []uuid.UUID) ([]models.Customer, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "customerRepository.GetCustomerByIDsDB")
	defer span.Finish()

	cutomers := make([]models.Customer, 0, len(customerIDs))
	rows, err := c.db.QueryxContext(ctx, getCustomerByIDsQuery, customerIDs)
	if err != nil {
		return nil, errors.Wrap(err, "customerRepository.GetCustomerByIDsDB.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		cRaw := models.Customer{}
		if err = rows.StructScan(cRaw); err != nil {
			return nil, errors.Wrap(err, "customerRepository.GetCustomerByIDsDB.StructScan")
		}
		cutomers = append(cutomers, cRaw)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "customerRepository.GetCustomerByIDsDB.rows.Err")
	}

	return cutomers, nil
}

// Delete customer by id
func (c *customerRepository) DeleteCustomerByIDDB(ctx context.Context, customerID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "customerRepository.DeleteCustomerByIDDB")
	defer span.Finish()

	result, err := c.db.ExecContext(ctx, deleteCustomerByIDQuery, customerID)
	if err != nil {
		return errors.Wrap(err, "customerRepository.DeleteCustomerByIDDB.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "customerRepository.DeleteCustomerByIDDB.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "customerRepository.DeleteCustomerByIDDB.rowsAffected")
	}

	return nil
}
