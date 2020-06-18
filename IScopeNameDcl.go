package yaccidl

type IScopeNameDcl interface {
	IYaccNode
	IsScopeName() bool
}
