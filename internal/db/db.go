package db

import (
	"context"
	"database/sql"
	"github.com/HarvinRaj/goldshop/configs"
)

type DB struct {
	db *sql.DB
}

// NewDBConnection ...
func NewDBConnection(ctx context.Context, config *configs.Config) (*DB, error) {
	return NewMySQL(ctx, config)
}

// GetConnection ...
func (db *DB) GetConnection() *sql.DB {
	return db.db
}

// Close ...
func (db *DB) Close() error {
	return db.db.Close()
}
