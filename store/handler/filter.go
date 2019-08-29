package handler

import (
	"bytes"
	"github.com/zfd81/rooster/store"
	"github.com/zfd81/rooster/util"
)

const (
	defaultFilterChainCapacity = 128
)

type FilterFunc func(line *store.Line) bool

func (f *FilterFunc) Filter(line *store.Line) bool {
	return (*f)(line)
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

func (fc *FilterChain) Filter(line *store.Line) bool {
	for _, filter := range *fc {
		if !filter(line) {
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
	return func(line *store.Line) bool {
		for _, v := range filters {
			if !v(line) {
				return false
			}
		}
		return true
	}
}

func OR(filters ...FilterFunc) FilterFunc {
	return func(line *store.Line) bool {
		for _, v := range filters {
			if v(line) {
				return true
			}
		}
		return false
	}
}

func Equal(a, b Expression) FilterFunc {
	return func(line *store.Line) bool {
		return util.Equal(a.Bytes(line), b.Bytes(line))
	}
}

func NotEqual(a, b Expression) FilterFunc {
	return func(line *store.Line) bool {
		return util.NotEqual(a.Bytes(line), b.Bytes(line))
	}
}

func Greater(a, b Expression) FilterFunc {
	return func(line *store.Line) bool {
		return util.Greater(a.Bytes(line), b.Bytes(line))
	}
}

func GreaterOrEqual(a, b Expression) FilterFunc {
	return func(line *store.Line) bool {
		return util.GreaterOrEqual(a.Bytes(line), b.Bytes(line))
	}
}

func Less(a, b Expression) FilterFunc {
	return func(line *store.Line) bool {
		return util.Less(a.Bytes(line), b.Bytes(line))
	}
}

func LessOrEqual(a, b Expression) FilterFunc {
	return func(line *store.Line) bool {
		return util.LessOrEqual(a.Bytes(line), b.Bytes(line))
	}
}

func In(a Expression, b [][]byte) FilterFunc {
	return func(line *store.Line) bool {
		return util.In(a.Bytes(line), b)
	}
}

func NotIn(a Expression, b [][]byte) FilterFunc {
	return func(line *store.Line) bool {
		return util.NotIn(a.Bytes(line), b)
	}
}

func Empty(val Expression) FilterFunc {
	return func(line *store.Line) bool {
		return util.Empty(val.Bytes(line))
	}
}

func NotEmpty(val Expression) FilterFunc {
	return func(line *store.Line) bool {
		return util.NotEmpty(val.Bytes(line))
	}
}

func HasPrefix(s Expression, prefix []byte) FilterFunc {
	return func(line *store.Line) bool {
		return bytes.HasPrefix(s.Bytes(line), prefix)
	}
}

func HasSuffix(s Expression, suffix []byte) FilterFunc {
	return func(line *store.Line) bool {
		return bytes.HasSuffix(s.Bytes(line), suffix)
	}
}
