package handler

import (
	"github.com/zfd81/rooster/store"
	"github.com/zfd81/rooster/store/types"
	"strconv"
)

type Parameter interface {
	Val(row *store.Line) []byte
}

type Constant struct {
	Type  types.ParamType
	value []byte
}

func (c *Constant) Val(row *store.Line) []byte {
	return c.value
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

type Variable struct {
	Type  types.ParamType
	index int
}

func (v *Variable) Val(row *store.Line) []byte {
	return *row.GetField(v.index)
}

func NewVar(index int) *Variable {
	return &Variable{types.Field, index}
}
