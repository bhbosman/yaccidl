package yaccidl

type IModuleDcl interface {
	IYaccNode
	ChildDecls() IYaccNode
	IsModuleDcl() bool
}
