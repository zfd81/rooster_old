package xsql

import (
	"database/sql"
)

type DB struct {
	*sql.DB
	driverName string
	unsafe     bool
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{DB: db, driverName: driverName}, err
}

func (db *DB) Execute(query string, arg interface{}) (sql.Result, error) {
	var sql string
	var arglist []interface{}
	var err error
	if maparg, ok := arg.(map[string]interface{}); ok {
		sql, arglist, err = bindMap(query, maparg)
	} else {
		// sql, arglist, err = bindStruct(query, maparg)
	}
	if err != nil {
		return nil, err
	}
	return db.Exec(sql, arglist...)
}
