package handler

import (
	"bytes"
	"github.com/zfd81/rooster/util"
)

const (
	defaultFilterChainCapacity = 128
)

type FilterFunc func() bool

type FilterChain []FilterFunc

func (fc *FilterChain) Add(filter FilterFunc) *FilterChain {
	if fc != nil {
		*fc = append(*fc, filter)
	}
	return fc
}

func (fc *FilterChain) Clear() *FilterChain {
	*fc = (*fc)[0:0]
	return fc
}

func NewFilterChain() *FilterChain {
	var chain FilterChain = make([]FilterFunc, 0, defaultFilterChainCapacity)
	return &chain
}

func AND(filters ...FilterFunc) FilterFunc {
	return func() bool {
		for _, v := range filters {
			if !v() {
				return false
			}
		}
		return true
	}
}

func OR(filters ...FilterFunc) FilterFunc {
	return func() bool {
		for _, v := range filters {
			if v() {
				return true
			}
		}
		return false
	}
}

func Equal(a, b []byte) FilterFunc {
	return func() bool {
		return util.Equal(a, b)
	}
}

func NotEqual(a, b []byte) FilterFunc {
	return func() bool {
		return util.NotEqual(a, b)
	}
}

func Greater(a, b []byte) FilterFunc {
	return func() bool {
		return util.Greater(a, b)
	}
}

func GreaterOrEqual(a, b []byte) FilterFunc {
	return func() bool {
		return util.GreaterOrEqual(a, b)
	}
}

func Less(a, b []byte) FilterFunc {
	return func() bool {
		return util.Less(a, b)
	}
}

func LessOrEqual(a, b []byte) FilterFunc {
	return func() bool {
		return util.LessOrEqual(a, b)
	}
}

func In(a []byte, b [][]byte) FilterFunc {
	return func() bool {
		return util.In(a, b)
	}
}

func NotIn(a []byte, b [][]byte) FilterFunc {
	return func() bool {
		return util.NotIn(a, b)
	}
}

func Empty(val []byte) FilterFunc {
	return func() bool {
		return util.Empty(val)
	}
}

func NotEmpty(val []byte) FilterFunc {
	return func() bool {
		return util.NotEmpty(val)
	}
}

func HasPrefix(s, prefix []byte) FilterFunc {
	return func() bool {
		return bytes.HasPrefix(s, prefix)
	}
}

func HasSuffix(s, suffix []byte) FilterFunc {
	return func() bool {
		return bytes.HasSuffix(s, suffix)
	}
}
