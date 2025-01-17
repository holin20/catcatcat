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
		// hmmm localhost won't work if connecting from container.
		// is there a better way?
		"postgres://%s:%s@host.docker.internal:%d/%s?sslmode=disable",
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

func (d *PostgresDB) Insert(table string, cols map[string]*SqlCol) error {
	insertSql := BuildInsertSql(table, cols)
	_, err := d.db.Exec(insertSql)
	if IsErr(err) {
		return NewCausef(err, "Insert(%s)", insertSql)
	}
	return nil
}

func (d *PostgresDB) Query(sqlString string) ([]string, [][]any, error) {
	rows, err := d.db.Query(sqlString)
	if IsErr(err) {
		return nil, nil, NewCausef(err, "Query(%s)", sqlString)
	}
	if rows == nil {
		return nil, nil, NewCausef(fmt.Errorf("nil rows on nil err by Query"), "Query(%s)", sqlString)
	}
	defer rows.Close()

	colNames, err := rows.Columns()
	if IsErr(err) {
		return nil, nil, NewCause(err, "Columns")
	}
	colTypes, err := rows.ColumnTypes()
	if IsErr(err) {
		return nil, nil, NewCause(err, "ColumnTypes")
	}

	var resultSet [][]any
	for rows.Next() {
		colsInPointer := make([]any, len(colTypes))
		for i := range colTypes {
			var col any
			colsInPointer[i] = &col
		}
		err := rows.Scan(colsInPointer...)
		if IsErr(err) {
			return nil, nil, NewCause(err, "PostgresDB_Scan")
		}
		colsInValue := SliceApply(colsInPointer, func(i int, p any) any {
			return *(p.(*any))
		})

		resultSet = append(resultSet, colsInValue)
	}
	if IsErr(err) {
		return nil, nil, NewCause(err, "ColumnNames")
	}

	return colNames, resultSet, nil
}
