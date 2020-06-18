package yaccidl

type scopeName struct {
	YaccNode
}

func (s scopeName) IsScopeName() bool {
	return true
}

func newScopeName(lexemData *LexemValue, identifier string) (IScopeNameDcl, error) {
	return &scopeName{
		YaccNode: *NewYaccNode(identifier, lexemData),
	}, nil
}
