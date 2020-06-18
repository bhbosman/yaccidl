package yaccidl

import "fmt"

type typeDeclarator struct {
	YaccNode
	simpleTypeSpec IYaccNode
	anyDeclarator  IYaccNode
}

func (t typeDeclarator) String() string {
	ss := fmt.Sprintf("%v", t.anyDeclarator)
	return fmt.Sprintf("typeDeclarator: (simpleTypeSpec: %v, anyDeclarator: %v)", t.simpleTypeSpec, ss)
}

func (t typeDeclarator) TypeSpec() IYaccNode {
	return t.simpleTypeSpec
}

func (t typeDeclarator) Declarator() IYaccNode {
	return t.anyDeclarator
}

func newTypeDeclarator(lexemData *LexemValue, identifier string, simpleTypeSpec IYaccNode, anyDeclarator IYaccNode) (ITypeDeclaratorDcl, error) {
	return &typeDeclarator{
		YaccNode:       *NewYaccNode(identifier, lexemData),
		simpleTypeSpec: simpleTypeSpec,
		anyDeclarator:  anyDeclarator,
	}, nil
}
