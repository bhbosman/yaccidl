package yaccidl

type ICharValue interface {
	IValue
	CharValue() rune
}
