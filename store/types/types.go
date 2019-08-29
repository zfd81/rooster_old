package types

type DataType byte

const (
	Int DataType = iota
	Decimal
	String
	Datetime
	Timestamp
	Bytes
)

type ParamType byte

const (
	Field ParamType = iota
	Constant
	Variable
	Expression
)
