package yaccidl

type IInterfaceHeader interface {
	Identifier() string
	InterfaceKind() IInterfaceKind
	Inheritance() IYaccNode
	LexemData() *LexemValue
	IsInterfaceHeader() bool
}
