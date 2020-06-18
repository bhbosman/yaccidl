package yaccidl

type IInt64Value interface {
	IValue
	Int64Value() int64
}
