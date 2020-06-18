package yaccidl

type IBooleanValue interface {
	IValue
	BooleanValue() bool
}
