package drivers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"errors"
)

type base struct {
	Type string
	dsn  string
	db   *sql.DB
}

func (b *base) Open() error {
	db, err := sql.Open(b.Type, b.dsn)
	if err != nil {
		return err
	}
	b.db = db
	return nil
}

func (b base) Query(q string) ([]map[string]interface{}) {
	if b.db == nil {
		panic(errors.New("need call Open() before Query()"))
	}

	rows, err := b.db.Query(q)

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//result
	var rs = make([]map[string]interface{}, 1)

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var row = map[string]interface{}{}
		for i, col := range values {
			if col == nil {
				row[columns[i]] = nil
			} else {
				row[columns[i]] = string(col)
			}

		}
		rs = append(rs, row)
	}

	return rs
}

func (b *base) Close() {
	if b.db == nil {
		panic(errors.New("need call Open() before Close()"))
	}
	b.db.Close()
}

func newBase(t string, dsn string) *base {
	return &base{t, dsn, nil}
}
