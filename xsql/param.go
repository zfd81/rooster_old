package xsql

import (
	"github.com/zfd81/rooster"
	"reflect"
)

type Params struct {
	params map[string]interface{}
}

func (p *Params) Get(name string) (interface{}, bool) {
	val, ok := p.params[name]
	return val, ok
}

func (p *Params) Add(name string, value interface{}) *Params {
	_, ok := p.params[name]
	if ok {
		delete(p.params, name)
	}
	p.params[name] = value
	return p
}

func (p *Params) Remove(name string) *Params {
	_, ok := p.params[name]
	if ok {
		delete(p.params, name)
	}
	return p
}
func (p *Params) Names() []string {
	if len(p.params) < 1 {
		return nil
	}
	names := make([]string, 5, 10)
	for k := range p.params {
		names = append(names, k)
	}
	return names
}

func (p *Params) Size(name string) int {
	return len(p.params)
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
	return &Params{params}, nil
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
	for i := 0; i < valueOfParams.NumField(); i++ {
		newParams[typeOfParams.Field(i).Name] = valueOfParams.Field(i).Interface()
	}
	return &Params{newParams}, nil
}
