package yaccidl

type IInterfaceKind interface {
	IYaccNode
	Local() bool
	Abstract() bool
	Custom() bool
	IsnterfaceKind() bool
}
