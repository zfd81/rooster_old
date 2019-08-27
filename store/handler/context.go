package handler

import (
	"github.com/zfd81/rooster/store"
	"github.com/zfd81/rooster/store/types"
	"strconv"
)

type Context struct {
	env []Variable
	row *store.Row
}

//func (c *Context) Add(val Variable) int {
//	index := len(c.env)
//	val.Index = index
//	c.env = append(c.env, val)
//	return index
//}

func (c *Context) SetRow(row *store.Row) *Context {
	if row != nil {
		c.row = row
	}
	return c
}

//func (c *Context) Get(param *Parameter) []byte {
//	if param != nil {
//		if param.Type == types.Field {
//			return *c.row.GetField(param.Index)
//		} else if param.Type == types.Variable {
//			return c.env[param.Index].Value
//		}
//	}
//	return nil
//}

func NewContext() *Context {
	return &Context{env: make([]Variable, 0, 20)}
}

type Parameter interface {
	Val(row *store.Row) []byte
}

type Constant struct {
	Type  types.ParamType
	value []byte
}

func (c *Constant) Val(row *store.Row) []byte {
	return c.value
}

type Variable struct {
	Type  types.ParamType
	index int
}

func (v *Variable) Val(row *store.Row) []byte {
	return *row.GetField(v.index)
}

func NewConst(val interface{}) *Constant {
	switch val.(type) {
	case string:
		return &Constant{types.Constant, []byte(val.(string))}
	case int:
		return &Constant{types.Constant, []byte(strconv.Itoa(val.(int)))}
	case float64:
		return &Constant{types.Constant, []byte(strconv.FormatFloat(val.(float64), 'f', -1, 64))}
	case []byte:
		return &Constant{types.Constant, val.([]byte)}
	default:
		return &Constant{types.Constant, nil}
	}
}

func NewVar(index int) *Variable {
	return &Variable{types.Field, index}
}
