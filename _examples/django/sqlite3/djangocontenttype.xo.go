package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// DjangoContentType represents a row from 'django_content_type'.
type DjangoContentType struct {
	ID       int    `json:"id"`        // id
	AppLabel string `json:"app_label"` // app_label
	Model    string `json:"model"`     // model
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [DjangoContentType] exists in the database.
func (dct *DjangoContentType) Exists() bool {
	return dct._exists
}

// Deleted returns true when the [DjangoContentType] has been marked for deletion
// from the database.
func (dct *DjangoContentType) Deleted() bool {
	return dct._deleted
}

// Insert inserts the [DjangoContentType] to the database.
func (dct *DjangoContentType) Insert(ctx context.Context, db DB) error {
	switch {
	case dct._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dct._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django_content_type (` +
		`id, app_label, model` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`
	// run
	logf(sqlstr, dct.AppLabel, dct.Model)
	res, err := db.ExecContext(ctx, sqlstr, dct.ID, dct.AppLabel, dct.Model)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	dct.ID = int(id)
	// set exists
	dct._exists = true
	return nil
}

// Update updates a [DjangoContentType] in the database.
func (dct *DjangoContentType) Update(ctx context.Context, db DB) error {
	switch {
	case !dct._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dct._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django_content_type SET ` +
		`app_label = $1, model = $2 ` +
		`WHERE id = $3`
	// run
	logf(sqlstr, dct.AppLabel, dct.Model, dct.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dct.AppLabel, dct.Model, dct.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [DjangoContentType] to the database.
func (dct *DjangoContentType) Save(ctx context.Context, db DB) error {
	if dct.Exists() {
		return dct.Update(ctx, db)
	}
	return dct.Insert(ctx, db)
}

// Upsert performs an upsert for [DjangoContentType].
func (dct *DjangoContentType) Upsert(ctx context.Context, db DB) error {
	switch {
	case dct._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO django_content_type (` +
		`id, app_label, model` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`app_label = EXCLUDED.app_label, model = EXCLUDED.model `
	// run
	logf(sqlstr, dct.ID, dct.AppLabel, dct.Model)
	if _, err := db.ExecContext(ctx, sqlstr, dct.ID, dct.AppLabel, dct.Model); err != nil {
		return logerror(err)
	}
	// set exists
	dct._exists = true
	return nil
}

// Delete deletes the [DjangoContentType] from the database.
func (dct *DjangoContentType) Delete(ctx context.Context, db DB) error {
	switch {
	case !dct._exists: // doesn't exist
		return nil
	case dct._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django_content_type ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, dct.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dct.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dct._deleted = true
	return nil
}

// DjangoContentTypeByAppLabelModel retrieves a row from 'django_content_type' as a [DjangoContentType].
//
// Generated from index 'django_content_type_app_label_model_76bd3d3b_uniq'.
func DjangoContentTypeByAppLabelModel(ctx context.Context, db DB, appLabel, model string) (*DjangoContentType, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, app_label, model ` +
		`FROM django_content_type ` +
		`WHERE app_label = $1 AND model = $2`
	// run
	logf(sqlstr, appLabel, model)
	dct := DjangoContentType{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, appLabel, model).Scan(&dct.ID, &dct.AppLabel, &dct.Model); err != nil {
		return nil, logerror(err)
	}
	return &dct, nil
}

// DjangoContentTypeByID retrieves a row from 'django_content_type' as a [DjangoContentType].
//
// Generated from index 'django_content_type_id_pkey'.
func DjangoContentTypeByID(ctx context.Context, db DB, id int) (*DjangoContentType, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, app_label, model ` +
		`FROM django_content_type ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	dct := DjangoContentType{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dct.ID, &dct.AppLabel, &dct.Model); err != nil {
		return nil, logerror(err)
	}
	return &dct, nil
}
