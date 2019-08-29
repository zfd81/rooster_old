package handler

import (
	"github.com/shopspring/decimal"
	"github.com/zfd81/rooster/store"
	"github.com/zfd81/rooster/store/types"
	"strconv"
)

type Expression interface {
	Bytes(line *store.Line) []byte
	Int(line *store.Line) int
	Decimal(line *store.Line) decimal.Decimal
}

type Constant struct {
	Type     types.ParamType
	DataType types.DataType
	value    []byte
}

func (c *Constant) Bytes(line *store.Line) []byte {
	return c.value
}

func (c *Constant) Int(line *store.Line) int {
	num, _ := strconv.Atoi(string(c.value))
	return num
}

func (c *Constant) Decimal(line *store.Line) decimal.Decimal {
	num, _ := decimal.NewFromString(string(c.value))
	return num
}

func NewConst(val interface{}) *Constant {
	switch val.(type) {
	case string:
		return &Constant{types.Constant, types.String, []byte(val.(string))}
	case int:
		return &Constant{types.Constant, types.Int, []byte(strconv.Itoa(val.(int)))}
	case float64:
		return &Constant{types.Constant, types.Decimal, []byte(strconv.FormatFloat(val.(float64), 'f', -1, 64))}
	case []byte:
		return &Constant{types.Constant, types.Bytes, val.([]byte)}
	default:
		return &Constant{types.Constant, types.Bytes, nil}
	}
}

type Variable struct {
	Type  types.ParamType
	index int
}

func (v *Variable) Bytes(line *store.Line) []byte {
	return *line.GetField(v.index)
}

func (v *Variable) Int(line *store.Line) int {
	num, _ := strconv.Atoi((*line).GetField(v.index).ToString())
	return num
}

func (v *Variable) Decimal(line *store.Line) decimal.Decimal {
	num, _ := decimal.NewFromString((*line).GetField(v.index).ToString())
	return num
}

func NewVar(index int) *Variable {
	return &Variable{types.Field, index}
}

type StrFunc func(line *store.Line) []byte

type NumFunc func(line *store.Line) decimal.Decimal
