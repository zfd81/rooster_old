package xsql

import (
	"database/sql"
	"reflect"
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

func NewMapParam(param map[string]interface{}) (*MapParam, error) {
	if param == nil {
		return nil, ErrParamNotNil
	}
	model := &MapParam{param}
	return model, nil
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

type StructParam struct {
	param interface{}
}

func NewStructParam(param interface{}) (*StructParam, error) {
	if param == nil {
		return nil, ErrParamNotNil
	}
	typeOfParam := reflect.TypeOf(param)
	if typeOfParam.Kind() != reflect.Struct {
		if typeOfParam.Kind() != reflect.Ptr {
			return nil, ErrParamType
		} else {
			if typeOfParam.Elem().Kind() != reflect.Struct {
				return nil, ErrParamType
			} else {
				// if reflect.ValueOf(param).Elem() == nil {
				// 	return nil, ErrParamNotNil
				// }
			}
		}
	}
	model := &StructParam{param}
	return model, nil
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
