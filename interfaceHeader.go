package yaccidl

type interfaceHeader struct {
	lexemData     *LexemValue
	identifier    string
	interfaceKind IInterfaceKind
	inheritance   IYaccNode
}

func (i interfaceHeader) IsInterfaceHeader() bool {
	return true
}

func (i interfaceHeader) LexemData() *LexemValue {
	return i.lexemData
}

func (i interfaceHeader) Inheritance() IYaccNode {
	return i.inheritance
}

func (i interfaceHeader) Identifier() string {
	return i.identifier
}

func (i interfaceHeader) InterfaceKind() IInterfaceKind {
	return i.interfaceKind
}

func newInterfaceHeader(lexemData *LexemValue, identifier string, interfaceKind IInterfaceKind, inheritance IYaccNode) (IInterfaceHeader, error) {
	return &interfaceHeader{
		lexemData:     lexemData,
		identifier:    identifier,
		interfaceKind: interfaceKind,
		inheritance:   inheritance,
	}, nil
}
