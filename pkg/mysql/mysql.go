package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adrianozp/go-plateful/pkg/config"
)

type MySQLDB struct {
	DB *sql.DB
}

func connString(cfg config.MySQL) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

// NewMySQLDB creates a new MySQLDB instance.
func NewMySQLDB(cfg config.MySQL) (*MySQLDB, error) {
	db, err := sql.Open("mysql", connString(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}
	return &MySQLDB{DB: db}, nil
}

func (db *MySQLDB) Close() error {
	return db.DB.Close()
}

func (db *MySQLDB) Execute(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	return result, nil
}

func (db *MySQLDB) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRowContext(ctx, query, args...)
}

func (db *MySQLDB) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.QueryContext(ctx, query, args...)
}
