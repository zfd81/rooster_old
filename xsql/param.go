package xsql

import (
	"github.com/zfd81/rooster"
	"reflect"
	"strings"
)

//忽略列名大小写
var ignoreCase bool = true

func setIgnoreCase(ignore bool) {
	ignoreCase = ignore
}

type Params struct {
	params map[string]interface{}
}

func (p *Params) Get(name string) (interface{}, bool) {
	if ignoreCase {
		name = strings.ToLower(name)
	}
	val, ok := p.params[name]
	return val, ok
}

func (p *Params) Add(name string, value interface{}) *Params {
	if ignoreCase {
		name = strings.ToLower(name)
	}
	_, ok := p.params[name]
	if ok {
		delete(p.params, name)
	}
	p.params[name] = value
	return p
}

func NewParams() *Params {
	model := &Params{make(map[string]interface{})}
	return model
}

func NewMapParams(params map[string]interface{}) (*Params, error) {
	if params == nil {
		return nil, rooster.ErrParamNotNil
	}
	size := len(params)
	if size < 1 {
		return nil, rooster.ErrParamEmpty
	}
	var model Params
	if ignoreCase {
		newParams := make(map[string]interface{})
		for k, v := range params {
			newParams[strings.ToLower(k)] = v
		}
		model.params = newParams
	} else {
		model.params = params
	}
	return &model, nil
}

func NewStructParams(params interface{}) (*Params, error) {
	if params == nil {
		return nil, rooster.ErrParamNotNil
	}
	typeOfParams := reflect.TypeOf(params)
	valueOfParams := reflect.ValueOf(params)
	if valueOfParams.Kind() == reflect.Ptr {
		typeOfParams = typeOfParams.Elem()
		valueOfParams = valueOfParams.Elem()
	}
	if valueOfParams.Kind() != reflect.Struct || !valueOfParams.IsValid() {
		return nil, rooster.ErrParamType
	}
	newParams := make(map[string]interface{})
	if ignoreCase {
		for i := 0; i < valueOfParams.NumField(); i++ {
			newParams[strings.ToLower(typeOfParams.Field(i).Name)] = valueOfParams.Field(i).Interface()
		}
	} else {
		for i := 0; i < valueOfParams.NumField(); i++ {
			newParams[typeOfParams.Field(i).Name] = valueOfParams.Field(i).Interface()
		}
	}
	model := &Params{newParams}
	return model, nil
}

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
		return nil, rooster.ErrParamNotNil
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
		return nil, rooster.ErrParamNotNil
	}
	typeOfParam := reflect.TypeOf(param)
	valueOfParam := reflect.ValueOf(param)
	if valueOfParam.Kind() == reflect.Ptr {
		typeOfParam = typeOfParam.Elem()
		valueOfParam = valueOfParam.Elem()
	}
	if valueOfParam.Kind() != reflect.Struct || !valueOfParam.IsValid() {
		return nil, rooster.ErrParamType
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
