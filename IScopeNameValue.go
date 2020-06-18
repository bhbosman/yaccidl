package yaccidl

type IScopeNameValue interface {
	IValue
	ScopeName() string
}
