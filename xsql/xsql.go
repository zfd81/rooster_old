package xsql

import (
	"database/sql"
)

type Paramer interface {
	Get(name string) (interface{}, bool)
}
type MapParam struct {
	param map[string]interface{}
}

func NewParam() *MapParam {
	model := &MapParam{make(map[string]interface{})}
	return model
}

func NewMapParam(param map[string]interface{}) *MapParam {
	model := &MapParam{param}
	return model
}

func (m *MapParam) Get(name string) (interface{}, bool) {
	val, ok := m.param[name]
	return val, ok
}

func (m *MapParam) Add(name string, value interface{}) *MapParam {
	_, ok := m.param[name]
	if ok {
		delete(m.param, name)
	}
	m.param[name] = value
	return m
}

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

// func (db *DB) Execute(query string, arg interface{}) (sql.Result, error) {
// 	var sql string
// 	var arglist []interface{}
// 	var err error
// 	if maparg, ok := arg.(map[string]interface{}); ok {
// 		sql, arglist, err = bindMap(query, maparg)
// 	} else {
// 		// sql, arglist, err = bindStruct(query, maparg)
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db.Exec(sql, arglist...)
// }
