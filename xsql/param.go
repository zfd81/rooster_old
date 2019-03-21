package xsql

import (
	"reflect"
	"strings"
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
	newParam := make(map[string]interface{})
	for k, v := range param {
		newParam[strings.ToLower(k)] = v
	}
	model := &MapParam{newParam}
	return model, nil
}

func (m *MapParam) Get(name string) (interface{}, bool) {
	val, ok := m.param[strings.ToLower(name)]
	return val, ok
}

func (m *MapParam) Add(name string, value interface{}) *MapParam {
	name = strings.ToLower(name)
	_, ok := m.param[name]
	if ok {
		delete(m.param, name)
	}
	m.param[name] = value
	return m
}

type StructParam struct {
	param         interface{}
	valueOfParam  reflect.Value
	fieldIndexMap map[string]int
}

func NewStructParam(param interface{}) (*StructParam, error) {
	if param == nil {
		return nil, ErrParamNotNil
	}
	typeOfParam := reflect.TypeOf(param)
	valueOfParam := reflect.ValueOf(param)
	if valueOfParam.Kind() == reflect.Ptr {
		typeOfParam = typeOfParam.Elem()
		valueOfParam = valueOfParam.Elem()
	}
	if valueOfParam.Kind() != reflect.Struct || !valueOfParam.IsValid() {
		return nil, ErrParamType
	}
	fieldNum := typeOfParam.NumField()
	fieldIndexMap := make(map[string]int)
	for i := 0; i < fieldNum; i++ {
		fieldIndexMap[strings.ToLower(typeOfParam.Field(i).Name)] = i + 1
	}
	model := &StructParam{param, valueOfParam, fieldIndexMap}
	return model, nil
}

func (s *StructParam) Get(name string) (interface{}, bool) {
	index := s.fieldIndexMap[strings.ToLower(name)]
	if index < 1 {
		return nil, false
	}
	valueOf := s.valueOfParam.Field(index - 1)
	if valueOf.IsValid() {
		return valueOf.Interface(), true
	}
	return nil, false
}
