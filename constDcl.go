package yaccidl

type constDcl struct {
	YaccNode
	value   IValue
	typeDef IYaccNode
}

func (c constDcl) IsConstDcl() bool {
	return true
}

func (c constDcl) TypeDef() IYaccNode {
	return c.typeDef
}

func (c constDcl) Value() IValue {
	return c.value
}

func NewConstDcl(lexemData *LexemValue, typeDef IYaccNode, identifier string, value IValue) (IConstDcl, error) {
	return &constDcl{
		YaccNode: *NewYaccNode(identifier, lexemData),
		value:    value,
		typeDef:  typeDef,
	}, nil
}
