package xsql

import (
	"github.com/zfd81/rooster/util"
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
	names := make([]string, 0, 10)
	for k := range p.params {
		names = append(names, k)
	}
	return names
}

func (p *Params) Size() int {
	return len(p.params)
}

func (p *Params) Iterator(handler func(key string, value interface{})) {
	if len(p.params) < 1 {
		return
	}
	for k, v := range p.params {
		handler(k, v)
	}
}

func NewParams() *Params {
	model := &Params{make(map[string]interface{})}
	return model
}

func NewMapParams(params map[string]interface{}) *Params {
	if params == nil || len(params) < 1 {
		return &Params{make(map[string]interface{})}
	}
	return &Params{params}
}

func NewStructParams(params interface{}) *Params {
	newParams := make(map[string]interface{})
	err := util.StructIterator(params, func(index int, key string, value interface{}) {
		newParams[key] = value
	})
	if err != nil {
		return &Params{make(map[string]interface{})}
	}
	return &Params{newParams}
}
