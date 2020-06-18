package yaccidl

type IAttributeDcl interface {
	IYaccNode
	TypeSpec() IYaccNode
	AttrDeclarator() IAttrDeclarator
	Readonly() bool
}
