package yaccidl

type nativeDcl struct {
	YaccNode
}

func (s nativeDcl) IsNativeDcl() bool {
	return true
}

func newNativeDcl(lexemData *LexemValue, identifier string) (INativeDcl, error) {
	return &nativeDcl{
		YaccNode: *NewYaccNode(identifier, lexemData),
	}, nil
}
