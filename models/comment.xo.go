// Package models contains generated code for schema 'mydb'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// Comment represents a row from 'mydb.comments'.
type Comment struct {
	ID        int       `json:"id"`         // id
	UserID    int       `json:"user_id"`    // user_id
	ProductID int       `json:"product_id"` // product_id
	Text      string    `json:"text"`       // text
	CreatedAt time.Time `json:"created_at"` // created_at
	UpdatedAt time.Time `json:"updated_at"` // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Comment] exists in the database.
func (c *Comment) Exists() bool {
	return c._exists
}

// Deleted returns true when the [Comment] has been marked for deletion
// from the database.
func (c *Comment) Deleted() bool {
	return c._deleted
}

// Insert inserts the [Comment] to the database.
func (c *Comment) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO mydb.comments (` +
		`user_id, product_id, text, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, c.UserID, c.ProductID, c.Text, c.CreatedAt, c.UpdatedAt)
	res, err := db.ExecContext(ctx, sqlstr, c.UserID, c.ProductID, c.Text, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	c.ID = int(id)
	// set exists
	c._exists = true
	return nil
}

// Update updates a [Comment] in the database.
func (c *Comment) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE mydb.comments SET ` +
		`user_id = ?, product_id = ?, text = ?, created_at = ?, updated_at = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, c.UserID, c.ProductID, c.Text, c.CreatedAt, c.UpdatedAt, c.ID)
	if _, err := db.ExecContext(ctx, sqlstr, c.UserID, c.ProductID, c.Text, c.CreatedAt, c.UpdatedAt, c.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Comment] to the database.
func (c *Comment) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Upsert performs an upsert for [Comment].
func (c *Comment) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO mydb.comments (` +
		`id, user_id, product_id, text, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`user_id = VALUES(user_id), product_id = VALUES(product_id), text = VALUES(text), created_at = VALUES(created_at), updated_at = VALUES(updated_at)`
	// run
	logf(sqlstr, c.ID, c.UserID, c.ProductID, c.Text, c.CreatedAt, c.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, c.ID, c.UserID, c.ProductID, c.Text, c.CreatedAt, c.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Delete deletes the [Comment] from the database.
func (c *Comment) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM mydb.comments ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, c.ID)
	if _, err := db.ExecContext(ctx, sqlstr, c.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CommentByID retrieves a row from 'mydb.comments' as a [Comment].
//
// Generated from index 'comments_id_pkey'.
func CommentByID(ctx context.Context, db DB, id int) (*Comment, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, product_id, text, created_at, updated_at ` +
		`FROM mydb.comments ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	c := Comment{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&c.ID, &c.UserID, &c.ProductID, &c.Text, &c.CreatedAt, &c.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}