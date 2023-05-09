package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// Book represents a row from 'public.books'.
type Book struct {
	BookID            int64     `json:"book_id"`              // book_id
	ISBN              string    `json:"isbn"`                 // isbn
	BookType          int       `json:"book_type"`            // book_type
	Title             string    `json:"title"`                // title
	Year              int       `json:"year"`                 // year
	Available         time.Time `json:"available"`            // available
	BooksAuthorIDFkey int64     `json:"books_author_id_fkey"` // books_author_id_fkey
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Book] exists in the database.
func (b *Book) Exists() bool {
	return b._exists
}

// Deleted returns true when the [Book] has been marked for deletion
// from the database.
func (b *Book) Deleted() bool {
	return b._deleted
}

// Insert inserts the [Book] to the database.
func (b *Book) Insert(ctx context.Context, db DB) error {
	switch {
	case b._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case b._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.books (` +
		`isbn, book_type, title, year, available, books_author_id_fkey` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING book_id`
	// run
	logf(sqlstr, b.ISBN, b.BookType, b.Title, b.Year, b.Available, b.BooksAuthorIDFkey)
	if err := db.QueryRowContext(ctx, sqlstr, b.ISBN, b.BookType, b.Title, b.Year, b.Available, b.BooksAuthorIDFkey).Scan(&b.BookID); err != nil {
		return logerror(err)
	}
	// set exists
	b._exists = true
	return nil
}

// Update updates a [Book] in the database.
func (b *Book) Update(ctx context.Context, db DB) error {
	switch {
	case !b._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case b._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.books SET ` +
		`isbn = $1, book_type = $2, title = $3, year = $4, available = $5, books_author_id_fkey = $6 ` +
		`WHERE book_id = $7`
	// run
	logf(sqlstr, b.ISBN, b.BookType, b.Title, b.Year, b.Available, b.BooksAuthorIDFkey, b.BookID)
	if _, err := db.ExecContext(ctx, sqlstr, b.ISBN, b.BookType, b.Title, b.Year, b.Available, b.BooksAuthorIDFkey, b.BookID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Book] to the database.
func (b *Book) Save(ctx context.Context, db DB) error {
	if b.Exists() {
		return b.Update(ctx, db)
	}
	return b.Insert(ctx, db)
}

// Upsert performs an upsert for [Book].
func (b *Book) Upsert(ctx context.Context, db DB) error {
	switch {
	case b._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.books (` +
		`book_id, isbn, book_type, title, year, available, books_author_id_fkey` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)` +
		` ON CONFLICT (book_id) DO ` +
		`UPDATE SET ` +
		`isbn = EXCLUDED.isbn, book_type = EXCLUDED.book_type, title = EXCLUDED.title, year = EXCLUDED.year, available = EXCLUDED.available, books_author_id_fkey = EXCLUDED.books_author_id_fkey `
	// run
	logf(sqlstr, b.BookID, b.ISBN, b.BookType, b.Title, b.Year, b.Available, b.BooksAuthorIDFkey)
	if _, err := db.ExecContext(ctx, sqlstr, b.BookID, b.ISBN, b.BookType, b.Title, b.Year, b.Available, b.BooksAuthorIDFkey); err != nil {
		return logerror(err)
	}
	// set exists
	b._exists = true
	return nil
}

// Delete deletes the [Book] from the database.
func (b *Book) Delete(ctx context.Context, db DB) error {
	switch {
	case !b._exists: // doesn't exist
		return nil
	case b._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.books ` +
		`WHERE book_id = $1`
	// run
	logf(sqlstr, b.BookID)
	if _, err := db.ExecContext(ctx, sqlstr, b.BookID); err != nil {
		return logerror(err)
	}
	// set deleted
	b._deleted = true
	return nil
}

// BooksByBooksAuthorIDFkey retrieves a row from 'public.books' as a [Book].
//
// Generated from index 'books_books_author_id_fkey_73ac0c26'.
func BooksByBooksAuthorIDFkey(ctx context.Context, db DB, booksAuthorIDFkey int64) ([]*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, isbn, book_type, title, year, available, books_author_id_fkey ` +
		`FROM public.books ` +
		`WHERE books_author_id_fkey = $1`
	// run
	logf(sqlstr, booksAuthorIDFkey)
	rows, err := db.QueryContext(ctx, sqlstr, booksAuthorIDFkey)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Book
	for rows.Next() {
		b := Book{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&b.BookID, &b.ISBN, &b.BookType, &b.Title, &b.Year, &b.Available, &b.BooksAuthorIDFkey); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// BookByBookID retrieves a row from 'public.books' as a [Book].
//
// Generated from index 'books_pkey'.
func BookByBookID(ctx context.Context, db DB, bookID int64) (*Book, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, isbn, book_type, title, year, available, books_author_id_fkey ` +
		`FROM public.books ` +
		`WHERE book_id = $1`
	// run
	logf(sqlstr, bookID)
	b := Book{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, bookID).Scan(&b.BookID, &b.ISBN, &b.BookType, &b.Title, &b.Year, &b.Available, &b.BooksAuthorIDFkey); err != nil {
		return nil, logerror(err)
	}
	return &b, nil
}

// Author returns the Author associated with the [Book]'s (BooksAuthorIDFkey).
//
// Generated from foreign key 'books_books_author_id_fkey_73ac0c26_fk_authors_author_id'.
func (b *Book) Author(ctx context.Context, db DB) (*Author, error) {
	return AuthorByAuthorID(ctx, db, b.BooksAuthorIDFkey)
}
