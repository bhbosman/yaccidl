package yaccidl

type attributeDcl struct {
	YaccNode
	typeSpec               IYaccNode
	readonlyAttrDeclarator IAttrDeclarator
	readonly               bool
}

func (a attributeDcl) Readonly() bool {
	return a.readonly
}

func (a attributeDcl) TypeSpec() IYaccNode {
	return a.typeSpec
}

func (a attributeDcl) AttrDeclarator() IAttrDeclarator {
	return a.readonlyAttrDeclarator
}

func newAttributeDcl(lexemData *LexemValue, identifier string, typeSpec IYaccNode, readonlyAttrDeclarator IAttrDeclarator, readonly bool) (IAttributeDcl, error) {
	return &attributeDcl{
		YaccNode:               *NewYaccNode(identifier, lexemData),
		typeSpec:               typeSpec,
		readonlyAttrDeclarator: readonlyAttrDeclarator,
		readonly:               readonly,
	}, nil
}
