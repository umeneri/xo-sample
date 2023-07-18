package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// PrismaMigration represents a row from 'mydb._prisma_migrations'.
type PrismaMigration struct {
	ID                string         `json:"id"`                  // id
	Checksum          string         `json:"checksum"`            // checksum
	FinishedAt        sql.NullTime   `json:"finished_at"`         // finished_at
	MigrationName     string         `json:"migration_name"`      // migration_name
	Logs              sql.NullString `json:"logs"`                // logs
	RolledBackAt      sql.NullTime   `json:"rolled_back_at"`      // rolled_back_at
	StartedAt         time.Time      `json:"started_at"`          // started_at
	AppliedStepsCount uint           `json:"applied_steps_count"` // applied_steps_count
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [PrismaMigration] exists in the database.
func (pm *PrismaMigration) Exists() bool {
	return pm._exists
}

// Deleted returns true when the [PrismaMigration] has been marked for deletion
// from the database.
func (pm *PrismaMigration) Deleted() bool {
	return pm._deleted
}

// Insert inserts the [PrismaMigration] to the database.
func (pm *PrismaMigration) Insert(ctx context.Context, db DB) error {
	switch {
	case pm._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case pm._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO mydb._prisma_migrations (` +
		`id, checksum, finished_at, migration_name, logs, rolled_back_at, started_at, applied_steps_count` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, pm.ID, pm.Checksum, pm.FinishedAt, pm.MigrationName, pm.Logs, pm.RolledBackAt, pm.StartedAt, pm.AppliedStepsCount)
	if _, err := db.ExecContext(ctx, sqlstr, pm.ID, pm.Checksum, pm.FinishedAt, pm.MigrationName, pm.Logs, pm.RolledBackAt, pm.StartedAt, pm.AppliedStepsCount); err != nil {
		return logerror(err)
	}
	// set exists
	pm._exists = true
	return nil
}

// Update updates a [PrismaMigration] in the database.
func (pm *PrismaMigration) Update(ctx context.Context, db DB) error {
	switch {
	case !pm._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case pm._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE mydb._prisma_migrations SET ` +
		`checksum = ?, finished_at = ?, migration_name = ?, logs = ?, rolled_back_at = ?, started_at = ?, applied_steps_count = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, pm.Checksum, pm.FinishedAt, pm.MigrationName, pm.Logs, pm.RolledBackAt, pm.StartedAt, pm.AppliedStepsCount, pm.ID)
	if _, err := db.ExecContext(ctx, sqlstr, pm.Checksum, pm.FinishedAt, pm.MigrationName, pm.Logs, pm.RolledBackAt, pm.StartedAt, pm.AppliedStepsCount, pm.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [PrismaMigration] to the database.
func (pm *PrismaMigration) Save(ctx context.Context, db DB) error {
	if pm.Exists() {
		return pm.Update(ctx, db)
	}
	return pm.Insert(ctx, db)
}

// Upsert performs an upsert for [PrismaMigration].
func (pm *PrismaMigration) Upsert(ctx context.Context, db DB) error {
	switch {
	case pm._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO mydb._prisma_migrations (` +
		`id, checksum, finished_at, migration_name, logs, rolled_back_at, started_at, applied_steps_count` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id = VALUES(id), checksum = VALUES(checksum), finished_at = VALUES(finished_at), migration_name = VALUES(migration_name), logs = VALUES(logs), rolled_back_at = VALUES(rolled_back_at), started_at = VALUES(started_at), applied_steps_count = VALUES(applied_steps_count)`
	// run
	logf(sqlstr, pm.ID, pm.Checksum, pm.FinishedAt, pm.MigrationName, pm.Logs, pm.RolledBackAt, pm.StartedAt, pm.AppliedStepsCount)
	if _, err := db.ExecContext(ctx, sqlstr, pm.ID, pm.Checksum, pm.FinishedAt, pm.MigrationName, pm.Logs, pm.RolledBackAt, pm.StartedAt, pm.AppliedStepsCount); err != nil {
		return logerror(err)
	}
	// set exists
	pm._exists = true
	return nil
}

// Delete deletes the [PrismaMigration] from the database.
func (pm *PrismaMigration) Delete(ctx context.Context, db DB) error {
	switch {
	case !pm._exists: // doesn't exist
		return nil
	case pm._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM mydb._prisma_migrations ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, pm.ID)
	if _, err := db.ExecContext(ctx, sqlstr, pm.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	pm._deleted = true
	return nil
}

// PrismaMigrationByID retrieves a row from 'mydb._prisma_migrations' as a [PrismaMigration].
//
// Generated from index '_prisma_migrations_id_pkey'.
func PrismaMigrationByID(ctx context.Context, db DB, id string) (*PrismaMigration, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, checksum, finished_at, migration_name, logs, rolled_back_at, started_at, applied_steps_count ` +
		`FROM mydb._prisma_migrations ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	pm := PrismaMigration{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&pm.ID, &pm.Checksum, &pm.FinishedAt, &pm.MigrationName, &pm.Logs, &pm.RolledBackAt, &pm.StartedAt, &pm.AppliedStepsCount); err != nil {
		return nil, logerror(err)
	}
	return &pm, nil
}