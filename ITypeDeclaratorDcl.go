package yaccidl

type ITypeDeclaratorDcl interface {
	IYaccNode
	TypeSpec() IYaccNode
	Declarator() IYaccNode
}
