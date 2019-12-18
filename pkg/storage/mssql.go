package storage

import (
	"fmt"
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type mssql struct {
	Openned bool
	Config *Connection
	Db *sql.DB
}

func (m *mssql) getConnString() (string, error) {
	return "?", nil
}

func (m *mssql) Open() (error) {
	if m.Openned {
		return errors.New("Db already openned")
	}

	connString, err := m.getConnString()
	if err != nil {
		return err
	}

	db, err := sql.Open("mssql", connString)
	if err != nil {
		return err
	}

	m.Db = db
	m.Openned = true
	return nil
}

func (m *mssql) Close() (error) {
	if !m.Openned {
		return errors.New("Db not openned, cannot close")
	}
	return nil
}

func (m *mssql) Query(query string, operator func([]string) (), args ...interface{}) (error) {
	if !m.Openned {
		return errors.New("Db not openned")
	}

	if query == "" {
		return errors.New("Missing query")
	}

	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}

	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	for rows.Next() {
		rawVals := make([][]byte, len(cols))
		rawValsPtrs := make([]interface{}, len(cols))

		for i, _ := range cols {
			rawValsPtrs[i] = &rawVals[i]
		}

		err = rows.Scan(rawValsPtrs...)
		if err != nil {
			return err
		}

		var data []string
		for _, val := range rawVals {
			data = append(data, fmt.Sprintf("%v", string(val)))
		}

		operator(data)
	}

	return nil
}