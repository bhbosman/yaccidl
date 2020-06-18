package yaccidl

type enumDcl struct {
	YaccNode
	enumerators IYaccNode
}

func (e enumDcl) IsEnumDcl() bool {
	return true
}

func (e enumDcl) Enumerator() IYaccNode {
	return e.enumerators
}

func newEnumDcl(identifier IYaccNode, enumerators IYaccNode) (IEnumDcl, error) {
	return &enumDcl{
		YaccNode:    *NewYaccNode(identifier.Identifier(), identifier.LexemData()),
		enumerators: enumerators,
	}, nil
}
