package xsql

import "errors"

var (
	ErrParamNotNil = errors.New("Parameters cannot be nil")
	ErrParamType   = errors.New("Parameter type error")
)
