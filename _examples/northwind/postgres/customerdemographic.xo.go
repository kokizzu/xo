package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// CustomerDemographic represents a row from 'public.customer_demographics'.
type CustomerDemographic struct {
	CustomerTypeID string         `json:"customer_type_id"` // customer_type_id
	CustomerDesc   sql.NullString `json:"customer_desc"`    // customer_desc
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [CustomerDemographic] exists in the database.
func (cd *CustomerDemographic) Exists() bool {
	return cd._exists
}

// Deleted returns true when the [CustomerDemographic] has been marked for deletion
// from the database.
func (cd *CustomerDemographic) Deleted() bool {
	return cd._deleted
}

// Insert inserts the [CustomerDemographic] to the database.
func (cd *CustomerDemographic) Insert(ctx context.Context, db DB) error {
	switch {
	case cd._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case cd._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.customer_demographics (` +
		`customer_type_id, customer_desc` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	logf(sqlstr, cd.CustomerTypeID, cd.CustomerDesc)
	if _, err := db.ExecContext(ctx, sqlstr, cd.CustomerTypeID, cd.CustomerDesc); err != nil {
		return logerror(err)
	}
	// set exists
	cd._exists = true
	return nil
}

// Update updates a [CustomerDemographic] in the database.
func (cd *CustomerDemographic) Update(ctx context.Context, db DB) error {
	switch {
	case !cd._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case cd._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.customer_demographics SET ` +
		`customer_desc = $1 ` +
		`WHERE customer_type_id = $2`
	// run
	logf(sqlstr, cd.CustomerDesc, cd.CustomerTypeID)
	if _, err := db.ExecContext(ctx, sqlstr, cd.CustomerDesc, cd.CustomerTypeID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [CustomerDemographic] to the database.
func (cd *CustomerDemographic) Save(ctx context.Context, db DB) error {
	if cd.Exists() {
		return cd.Update(ctx, db)
	}
	return cd.Insert(ctx, db)
}

// Upsert performs an upsert for [CustomerDemographic].
func (cd *CustomerDemographic) Upsert(ctx context.Context, db DB) error {
	switch {
	case cd._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.customer_demographics (` +
		`customer_type_id, customer_desc` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (customer_type_id) DO ` +
		`UPDATE SET ` +
		`customer_desc = EXCLUDED.customer_desc `
	// run
	logf(sqlstr, cd.CustomerTypeID, cd.CustomerDesc)
	if _, err := db.ExecContext(ctx, sqlstr, cd.CustomerTypeID, cd.CustomerDesc); err != nil {
		return logerror(err)
	}
	// set exists
	cd._exists = true
	return nil
}

// Delete deletes the [CustomerDemographic] from the database.
func (cd *CustomerDemographic) Delete(ctx context.Context, db DB) error {
	switch {
	case !cd._exists: // doesn't exist
		return nil
	case cd._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.customer_demographics ` +
		`WHERE customer_type_id = $1`
	// run
	logf(sqlstr, cd.CustomerTypeID)
	if _, err := db.ExecContext(ctx, sqlstr, cd.CustomerTypeID); err != nil {
		return logerror(err)
	}
	// set deleted
	cd._deleted = true
	return nil
}

// CustomerDemographicByCustomerTypeID retrieves a row from 'public.customer_demographics' as a [CustomerDemographic].
//
// Generated from index 'customer_demographics_pkey'.
func CustomerDemographicByCustomerTypeID(ctx context.Context, db DB, customerTypeID string) (*CustomerDemographic, error) {
	// query
	const sqlstr = `SELECT ` +
		`customer_type_id, customer_desc ` +
		`FROM public.customer_demographics ` +
		`WHERE customer_type_id = $1`
	// run
	logf(sqlstr, customerTypeID)
	cd := CustomerDemographic{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, customerTypeID).Scan(&cd.CustomerTypeID, &cd.CustomerDesc); err != nil {
		return nil, logerror(err)
	}
	return &cd, nil
}
