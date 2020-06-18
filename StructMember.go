package yaccidl

type StructMember struct {
	YaccNode
	typeDcl    IYaccNode
	declarator IYaccNode
}

func (s StructMember) IsStructMember() bool {
	return true
}

func (s StructMember) TypeDcl() IYaccNode {
	return s.typeDcl
}

func (s StructMember) Declarator() IYaccNode {
	return s.declarator
}

func NewStructMember(typeDcl IYaccNode, declarators IYaccNode) (IStructMemberDcl, error) {
	return &StructMember{
		YaccNode:   *NewYaccNode("(struct member)", typeDcl.LexemData()),
		typeDcl:    typeDcl,
		declarator: declarators,
	}, nil
}
