package yaccidl

type operationalParameter struct {
	YaccNode
	direction ParamDirection
	paramType IYaccNode
}

func (o operationalParameter) Direction() ParamDirection {
	return o.direction
}

func (o operationalParameter) ParamType() IYaccNode {
	return o.paramType
}

func newOperationalParameter(identifier string, lexemData *LexemValue, direction ParamDirection, paramType IYaccNode) (IOperationalParameter, error) {
	return &operationalParameter{
		YaccNode:  *NewYaccNode(identifier, lexemData),
		direction: direction,
		paramType: paramType,
	}, nil
}
