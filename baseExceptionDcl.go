package yaccidl

type baseExceptionDcl struct {
	YaccNode
	member IYaccNode
}

type exceptionDcl struct {
	baseExceptionDcl
}

func (e exceptionDcl) IsExceptionDcl() bool {
	return true
}

func (e exceptionDcl) GetMember() IYaccNode {
	return e.member
}

func newExceptionDcl(identifier IYaccNode, member IYaccNode) (IExceptionDcl, error) {
	return &exceptionDcl{
		baseExceptionDcl: baseExceptionDcl{
			YaccNode: *NewYaccNode(identifier.Identifier(), identifier.LexemData()),
			member:   member,
		},
	}, nil
}
