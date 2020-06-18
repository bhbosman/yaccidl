package yaccidl

type StructDcl struct {
	baseExceptionDcl
	scopeName IYaccNode
	forward   bool
}

func (s StructDcl) GetMember() IYaccNode {
	return s.member
}

func (s StructDcl) IsStructDcl() bool {
	return true
}

func (s StructDcl) InheritsFrom() IYaccNode {
	return s.scopeName
}

func (s StructDcl) Forward() bool {
	return s.forward
}

func NewStructType(identifier IYaccNode, scopeName IYaccNode, member IYaccNode, forward bool) (IStructDcl, error) {
	return &StructDcl{
		baseExceptionDcl: baseExceptionDcl{
			YaccNode: *NewYaccNode(identifier.Identifier(), identifier.LexemData()),
			member:   member,
		},
		scopeName: scopeName,
		forward:   forward,
	}, nil

}
