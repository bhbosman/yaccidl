package yaccidl

type IWideCharValue interface {
	IValue
	WideCharValue() rune
}
