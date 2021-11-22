package repository

const (
	insertCustomerQuery = `INSERT INTO customer (name, email, msisdn, address, created_at) 
					VALUES ($1, $2, $3, $4, now()) 
					RETURNING id`

	updateCustomerQuery = `UPDATE customer 
					SET name = COALESCE(NULLIF($1, ''), name),
						email = COALESCE(NULLIF($2, ''), email), 
					    msisdn = COALESCE(NULLIF($3, ''), msisdn), 
					    address = COALESCE(NULLIF($4, ''), address), 
					    updated_at = now() 
					WHERE id = $5`

	getCustomerByIDsQuery = `SELECT COALESCE(id, 0),
	COALESCE(name, ''),
	COALESCE(email, ''),
	COALESCE(msisdn, ''),
	COALESCE(address, '')
FROM customer
WHERE id in ($1)`

	deleteCustomerByIDQuery = `DELETE FROM customer WHERE id = $1`
)
