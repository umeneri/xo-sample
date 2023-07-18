package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// Product represents a row from 'mydb.products'.
type Product struct {
	ID          int       `json:"id"`          // id
	Name        string    `json:"name"`        // name
	Price       float64   `json:"price"`       // price
	Description string    `json:"description"` // description
	CreatedAt   time.Time `json:"created_at"`  // created_at
	UpdatedAt   time.Time `json:"updated_at"`  // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Product] exists in the database.
func (p *Product) Exists() bool {
	return p._exists
}

// Deleted returns true when the [Product] has been marked for deletion
// from the database.
func (p *Product) Deleted() bool {
	return p._deleted
}

// Insert inserts the [Product] to the database.
func (p *Product) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO mydb.products (` +
		`name, price, description, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, p.Name, p.Price, p.Description, p.CreatedAt, p.UpdatedAt)
	res, err := db.ExecContext(ctx, sqlstr, p.Name, p.Price, p.Description, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	p.ID = int(id)
	// set exists
	p._exists = true
	return nil
}

// Update updates a [Product] in the database.
func (p *Product) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE mydb.products SET ` +
		`name = ?, price = ?, description = ?, created_at = ?, updated_at = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, p.Name, p.Price, p.Description, p.CreatedAt, p.UpdatedAt, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.Name, p.Price, p.Description, p.CreatedAt, p.UpdatedAt, p.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Product] to the database.
func (p *Product) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for [Product].
func (p *Product) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO mydb.products (` +
		`id, name, price, description, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`name = VALUES(name), price = VALUES(price), description = VALUES(description), created_at = VALUES(created_at), updated_at = VALUES(updated_at)`
	// run
	logf(sqlstr, p.ID, p.Name, p.Price, p.Description, p.CreatedAt, p.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.Name, p.Price, p.Description, p.CreatedAt, p.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the [Product] from the database.
func (p *Product) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM mydb.products ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// ProductByID retrieves a row from 'mydb.products' as a [Product].
//
// Generated from index 'products_id_pkey'.
func ProductByID(ctx context.Context, db DB, id int) (*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, price, description, created_at, updated_at ` +
		`FROM mydb.products ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	p := Product{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}