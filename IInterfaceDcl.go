package yaccidl

type IInterfaceDcl interface {
	IYaccNode
	InterfaceHeader() IInterfaceHeader
	Body() IYaccNode
	Forward() bool
	IsInterfaceDcl() bool
}
