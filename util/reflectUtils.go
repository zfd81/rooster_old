package util

import (
	"github.com/zfd81/rooster/xsql"
	"reflect"
)

type IteratorFunc func(index int, key string, value interface{})

func StructIterator(arg interface{}, iterator IteratorFunc) error {
	if arg == nil {
		return xsql.ErrParamNotNil
	}
	typeOfArg := reflect.TypeOf(arg)
	valueOfArg := reflect.ValueOf(arg)
	if valueOfArg.Kind() == reflect.Ptr {
		typeOfArg = typeOfArg.Elem()
		valueOfArg = valueOfArg.Elem()
	}
	if valueOfArg.Kind() != reflect.Struct || !valueOfArg.IsValid() {
		return xsql.ErrParamType
	}
	for i := 0; i < valueOfArg.NumField(); i++ {
		iterator(i, typeOfArg.Field(i).Name, valueOfArg.Field(i).Interface())
	}
	return nil
}
