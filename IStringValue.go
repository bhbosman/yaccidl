package yaccidl

type IStringValue interface {
	IValue
	StringValue() string
}
