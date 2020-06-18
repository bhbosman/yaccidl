package yaccidl

type module struct {
	YaccNode
	childDecls IYaccNode
}

func (m module) IsModuleDcl() bool {
	return true
}

func (m module) ChildDecls() IYaccNode {
	return m.childDecls
}

func newModuleDcl(lexemData *LexemValue, identifier string, childDecls IYaccNode) (IModuleDcl, error) {
	return &module{
		YaccNode:   *NewYaccNode(identifier, lexemData),
		childDecls: childDecls,
	}, nil
}
