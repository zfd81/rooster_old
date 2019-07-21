package xsql

import (
	"bytes"
	"fmt"
	"github.com/zfd81/rooster/util"
	"reflect"
)

func bindParams(sql string, arg *Params) (string, []interface{}, error) {
	newSql, err := util.ReplaceBetween(sql, "{", "}", func(index int, start int, end int, content string) (string, error) {
		ignore := false
		fragment, err := util.ReplaceByKeyword(content, ':', func(i int, s int, e int, c string) (string, error) {
			_, ok := arg.Get(c)
			if !ok {
				ignore = true
				return "", nil
			}
			return fmt.Sprintf(":%s", c), nil
		})
		if ignore {
			return "", err
		}
		return fragment, err
	})
	if err != nil {
		return "", nil, err
	}
	params := make([]interface{}, 0, 20)
	newSql, err = util.ReplaceByKeyword(newSql, ':', func(index int, start int, end int, content string) (string, error) {
		val, ok := arg.Get(content)
		if !ok {
			return "?", fmt.Errorf("could not find name %s in %#v", content, arg)
		}
		params = append(params, val)
		return "?", nil
	})
	return newSql, params, err
}

func insertByMap(table string, object map[string]interface{}) (string, []interface{}, error) {
	if table == "" || object == nil {
		return "", nil, ErrParamNotNil
	}
	size := len(object)
	if size < 1 {
		return "", nil, ErrParamEmpty
	}
	var sql bytes.Buffer
	var sql2 bytes.Buffer
	sql.WriteString("insert into ")
	sql.WriteString(table)
	sql.WriteString(" (")
	sql2.WriteString(") values (")
	params := make([]interface{}, 0, 20)
	index := 0
	for k, v := range object {
		if index == 0 {
			index++
		} else {
			sql.WriteString(",")
			sql2.WriteString(",")
		}
		sql.WriteString(k)
		sql2.WriteString("?")
		params = append(params, v)
	}
	sql.WriteString(sql2.String())
	sql.WriteString(")")
	return sql.String(), params, nil
}

func insertByStruct(table string, object interface{}) (string, []interface{}, error) {
	if table == "" || object == nil {
		return "", nil, ErrParamNotNil
	}
	typeOfObject := reflect.TypeOf(object)
	valueOfObject := reflect.ValueOf(object)
	if valueOfObject.Kind() == reflect.Ptr {
		typeOfObject = typeOfObject.Elem()
		valueOfObject = valueOfObject.Elem()
	}
	if valueOfObject.Kind() != reflect.Struct || !valueOfObject.IsValid() {
		return "", nil, ErrParamType
	}
	var sql bytes.Buffer
	var sql2 bytes.Buffer
	sql.WriteString("insert into ")
	sql.WriteString(table)
	sql.WriteString(" (")
	sql2.WriteString(") values (")
	params := make([]interface{}, 0, 20)
	for i := 0; i < valueOfObject.NumField(); i++ {
		if i > 0 {
			sql.WriteString(",")
			sql2.WriteString(",")
		}
		sql.WriteString(typeOfObject.Field(i).Name)
		sql2.WriteString("?")
		params = append(params, valueOfObject.Field(i).Interface())
	}
	sql.WriteString(sql2.String())
	sql.WriteString(")")
	return sql.String(), params, nil
}
