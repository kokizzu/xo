package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// AuthPermission represents a row from 'public.auth_permission'.
type AuthPermission struct {
	ID            int    `json:"id"`              // id
	Name          string `json:"name"`            // name
	ContentTypeID int    `json:"content_type_id"` // content_type_id
	Codename      string `json:"codename"`        // codename
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [AuthPermission] exists in the database.
func (ap *AuthPermission) Exists() bool {
	return ap._exists
}

// Deleted returns true when the [AuthPermission] has been marked for deletion
// from the database.
func (ap *AuthPermission) Deleted() bool {
	return ap._deleted
}

// Insert inserts the [AuthPermission] to the database.
func (ap *AuthPermission) Insert(ctx context.Context, db DB) error {
	switch {
	case ap._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ap._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.auth_permission (` +
		`name, content_type_id, codename` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`
	// run
	logf(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename)
	if err := db.QueryRowContext(ctx, sqlstr, ap.Name, ap.ContentTypeID, ap.Codename).Scan(&ap.ID); err != nil {
		return logerror(err)
	}
	// set exists
	ap._exists = true
	return nil
}

// Update updates a [AuthPermission] in the database.
func (ap *AuthPermission) Update(ctx context.Context, db DB) error {
	switch {
	case !ap._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ap._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.auth_permission SET ` +
		`name = $1, content_type_id = $2, codename = $3 ` +
		`WHERE id = $4`
	// run
	logf(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename, ap.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ap.Name, ap.ContentTypeID, ap.Codename, ap.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [AuthPermission] to the database.
func (ap *AuthPermission) Save(ctx context.Context, db DB) error {
	if ap.Exists() {
		return ap.Update(ctx, db)
	}
	return ap.Insert(ctx, db)
}

// Upsert performs an upsert for [AuthPermission].
func (ap *AuthPermission) Upsert(ctx context.Context, db DB) error {
	switch {
	case ap._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.auth_permission (` +
		`id, name, content_type_id, codename` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, content_type_id = EXCLUDED.content_type_id, codename = EXCLUDED.codename `
	// run
	logf(sqlstr, ap.ID, ap.Name, ap.ContentTypeID, ap.Codename)
	if _, err := db.ExecContext(ctx, sqlstr, ap.ID, ap.Name, ap.ContentTypeID, ap.Codename); err != nil {
		return logerror(err)
	}
	// set exists
	ap._exists = true
	return nil
}

// Delete deletes the [AuthPermission] from the database.
func (ap *AuthPermission) Delete(ctx context.Context, db DB) error {
	switch {
	case !ap._exists: // doesn't exist
		return nil
	case ap._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.auth_permission ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, ap.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ap.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	ap._deleted = true
	return nil
}

// AuthPermissionByContentTypeID retrieves a row from 'public.auth_permission' as a [AuthPermission].
//
// Generated from index 'auth_permission_content_type_id_2f476e4b'.
func AuthPermissionByContentTypeID(ctx context.Context, db DB, contentTypeID int) ([]*AuthPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM public.auth_permission ` +
		`WHERE content_type_id = $1`
	// run
	logf(sqlstr, contentTypeID)
	rows, err := db.QueryContext(ctx, sqlstr, contentTypeID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AuthPermission
	for rows.Next() {
		ap := AuthPermission{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ap)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AuthPermissionByContentTypeIDCodename retrieves a row from 'public.auth_permission' as a [AuthPermission].
//
// Generated from index 'auth_permission_content_type_id_codename_01ab375a_uniq'.
func AuthPermissionByContentTypeIDCodename(ctx context.Context, db DB, contentTypeID int, codename string) (*AuthPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM public.auth_permission ` +
		`WHERE content_type_id = $1 AND codename = $2`
	// run
	logf(sqlstr, contentTypeID, codename)
	ap := AuthPermission{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, contentTypeID, codename).Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename); err != nil {
		return nil, logerror(err)
	}
	return &ap, nil
}

// AuthPermissionByID retrieves a row from 'public.auth_permission' as a [AuthPermission].
//
// Generated from index 'auth_permission_pkey'.
func AuthPermissionByID(ctx context.Context, db DB, id int) (*AuthPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM public.auth_permission ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	ap := AuthPermission{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename); err != nil {
		return nil, logerror(err)
	}
	return &ap, nil
}

// DjangoContentType returns the DjangoContentType associated with the [AuthPermission]'s (ContentTypeID).
//
// Generated from foreign key 'auth_permission_content_type_id_2f476e4b_fk_django_co'.
func (ap *AuthPermission) DjangoContentType(ctx context.Context, db DB) (*DjangoContentType, error) {
	return DjangoContentTypeByID(ctx, db, ap.ContentTypeID)
}
