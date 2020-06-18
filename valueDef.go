package yaccidl

type valueDef struct {
	YaccNode
	valueHeader IValueHeader
	body        IYaccNode
	forward     bool
}

func (v valueDef) Forward() bool {
	return v.forward
}

func (v valueDef) ValueHeader() IValueHeader {
	return v.valueHeader
}

func (v valueDef) Body() IYaccNode {
	return v.body
}

func newValueDef(lexemData *LexemValue, identifier string, valueHeader IValueHeader, valueElement IYaccNode, forward bool) (IValueDefDcl, error) {
	return &valueDef{
		YaccNode:    *NewYaccNode(identifier, lexemData),
		valueHeader: valueHeader,
		body:        valueElement,
		forward:     forward,
	}, nil
}
