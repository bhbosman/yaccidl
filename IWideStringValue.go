package yaccidl

type IWideStringValue interface {
	IValue
	WideStringValue() string
}
