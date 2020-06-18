package yaccidl

type IValueHeader interface {
	IYaccNode
	ValueKind() IInterfaceKind
	ValueInheritanceSpec() IValueInheritanceSpec
}
