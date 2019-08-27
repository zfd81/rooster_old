package handler

import (
	"bytes"
	"github.com/zfd81/rooster/store"
	"github.com/zfd81/rooster/util"
)

const (
	defaultFilterChainCapacity = 128
)

type FilterFunc func(row *store.Row) bool

func (f *FilterFunc) Filter(row *store.Row) bool {
	return (*f)(row)
}

type FilterChain []FilterFunc

func (fc *FilterChain) Add(filter ...FilterFunc) *FilterChain {
	if filter != nil {
		*fc = append(*fc, filter...)
	}
	return fc
}

func (fc *FilterChain) AddOR(filters ...FilterFunc) *FilterChain {
	if filters != nil {
		*fc = append(*fc, OR(filters...))
	}
	return fc
}

func (fc *FilterChain) Clear() *FilterChain {
	*fc = (*fc)[0:0]
	return fc
}

func (fc *FilterChain) Filter(row *store.Row) bool {
	for _, filter := range *fc {
		if !filter(row) {
			return false
		}
	}
	return true
}

func NewFilterChain() *FilterChain {
	var chain FilterChain = make([]FilterFunc, 0, defaultFilterChainCapacity)
	return &chain
}

func AND(filters ...FilterFunc) FilterFunc {
	return func(row *store.Row) bool {
		for _, v := range filters {
			if !v(row) {
				return false
			}
		}
		return true
	}
}

func OR(filters ...FilterFunc) FilterFunc {
	return func(row *store.Row) bool {
		for _, v := range filters {
			if v(row) {
				return true
			}
		}
		return false
	}
}

func Equal(a, b Parameter) FilterFunc {
	return func(row *store.Row) bool {
		return util.Equal(a.Val(row), b.Val(row))
	}
}

func NotEqual(a, b Parameter) FilterFunc {
	return func(row *store.Row) bool {
		return util.NotEqual(a.Val(row), b.Val(row))
	}
}

func Greater(a, b Parameter) FilterFunc {
	return func(row *store.Row) bool {
		return util.Greater(a.Val(row), b.Val(row))
	}
}

func GreaterOrEqual(a, b Parameter) FilterFunc {
	return func(row *store.Row) bool {
		return util.GreaterOrEqual(a.Val(row), b.Val(row))
	}
}

func Less(a, b Parameter) FilterFunc {
	return func(row *store.Row) bool {
		return util.Less(a.Val(row), b.Val(row))
	}
}

func LessOrEqual(a, b Parameter) FilterFunc {
	return func(row *store.Row) bool {
		return util.LessOrEqual(a.Val(row), b.Val(row))
	}
}

func In(a Parameter, b [][]byte) FilterFunc {
	return func(row *store.Row) bool {
		return util.In(a.Val(row), b)
	}
}

func NotIn(a Parameter, b [][]byte) FilterFunc {
	return func(row *store.Row) bool {
		return util.NotIn(a.Val(row), b)
	}
}

func Empty(val Parameter) FilterFunc {
	return func(row *store.Row) bool {
		return util.Empty(val.Val(row))
	}
}

func NotEmpty(val Parameter) FilterFunc {
	return func(row *store.Row) bool {
		return util.NotEmpty(val.Val(row))
	}
}

func HasPrefix(s Parameter, prefix []byte) FilterFunc {
	return func(row *store.Row) bool {
		return bytes.HasPrefix(s.Val(row), prefix)
	}
}

func HasSuffix(s Parameter, suffix []byte) FilterFunc {
	return func(row *store.Row) bool {
		return bytes.HasSuffix(s.Val(row), suffix)
	}
}
