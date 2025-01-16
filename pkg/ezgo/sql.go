package ezgo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewLocalPostgresDB(
	userName string,
	password string,
	port uint16,
	dbName string,
) (*PostgresDB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@localhost:%d/%s?sslmode=disable",
		userName,
		password,
		port,
		dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if IsErr(err) {
		return nil, NewCause(err, "Open Postgresql DB")
	}

	return &PostgresDB{db: db}, nil
}

func (d *PostgresDB) Close() {
	d.db.Close()
}

func (d *PostgresDB) Insert(table string, cols map[string]*sqlCol) error {
	insertSql := BuildInsertSql(table, cols)
	_, err := d.db.Exec(insertSql)
	if IsErr(err) {
		return NewCausef(err, "Insert(%s)", insertSql)
	}
	return nil
}
