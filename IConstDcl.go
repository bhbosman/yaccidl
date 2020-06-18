package yaccidl

type IConstDcl interface {
	IYaccNode
	Value() IValue
	TypeDef() IYaccNode
	IsConstDcl() bool
}
