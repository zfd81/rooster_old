package handler

import (
	"github.com/shopspring/decimal"
	"github.com/zfd81/rooster/store"
	"github.com/zfd81/rooster/store/types"
)

type Computer struct {
	Type     types.ParamType
	function NumFunc
}

func (c *Computer) Bytes(line *store.Line) []byte {
	return []byte(c.function(line).String())
}

func (c *Computer) Int(line *store.Line) int {
	num := c.function(line).IntPart()
	return int(num)
}

func (c *Computer) Decimal(line *store.Line) decimal.Decimal {
	return c.function(line)
}

func NewComputer(function NumFunc) *Computer {
	return &Computer{types.Expression, function}
}

func Add(a, b Expression) Expression {
	return NewComputer(func(line *store.Line) decimal.Decimal {
		return a.Decimal(line).Add(b.Decimal(line))
	})
}

func Sub(a, b Expression) Expression {
	return NewComputer(func(line *store.Line) decimal.Decimal {
		return a.Decimal(line).Sub(b.Decimal(line))
	})
}

func Mul(a, b Expression) Expression {
	return NewComputer(func(line *store.Line) decimal.Decimal {
		return a.Decimal(line).Mul(b.Decimal(line))
	})
}

func Div(a, b Expression) Expression {
	return NewComputer(func(line *store.Line) decimal.Decimal {
		return a.Decimal(line).Div(b.Decimal(line))
	})
}

func Mod(a, b Expression) Expression {
	return NewComputer(func(line *store.Line) decimal.Decimal {
		return a.Decimal(line).Mod(b.Decimal(line))
	})
}
