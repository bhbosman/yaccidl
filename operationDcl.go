package yaccidl

type operationDcl struct {
	YaccNode
	paramDcl     IYaccNode
	returnType   IYaccNode
	exceptionDcl IYaccNode
}

func (o operationDcl) ParamDcl() IYaccNode {
	return o.paramDcl
}

func (o operationDcl) ReturnType() IYaccNode {
	return o.returnType
}

func (o operationDcl) ExceptionDcl() IYaccNode {
	return o.exceptionDcl
}

func newOperationDcl(identifier IYaccNode, paramdDcl IYaccNode, returnType IYaccNode, exceptionDcl IYaccNode) (IOperationDcl, error) {
	return &operationDcl{
		YaccNode:     *NewYaccNode(identifier.Identifier(), identifier.LexemData()),
		paramDcl:     paramdDcl,
		returnType:   returnType,
		exceptionDcl: exceptionDcl,
	}, nil
}
