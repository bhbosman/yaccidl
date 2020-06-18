package yaccidl

type interfaceDcl struct {
	YaccNode
	interfaceHeader IInterfaceHeader
	body            IYaccNode
	forward         bool
}

func (i interfaceDcl) IsInterfaceDcl() bool {
	return true
}

func (i interfaceDcl) Forward() bool {
	return i.forward
}

func (i interfaceDcl) InterfaceHeader() IInterfaceHeader {
	return i.interfaceHeader
}

func (i interfaceDcl) Body() IYaccNode {
	return i.body
}

func newInterfaceDcl(lexemData *LexemValue, identifier string, interfaceHeader IInterfaceHeader, body IYaccNode, forward bool) (IInterfaceDcl, error) {
	return &interfaceDcl{
		YaccNode:        *NewYaccNode(identifier, lexemData),
		interfaceHeader: interfaceHeader,
		body:            body,
		forward:         forward,
	}, nil
}
