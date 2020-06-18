package yaccidl

type IEnumDcl interface {
	IYaccNode
	Enumerator() IYaccNode
	IsEnumDcl() bool
}
