package yaccidl

type IStructMemberDcl interface {
	IYaccNode
	TypeDcl() IYaccNode
	Declarator() IYaccNode
	IsStructMember() bool
}
