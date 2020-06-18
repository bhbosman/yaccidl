package yaccidl

type readonlyAttrDeclarator struct {
	simpleDeclarator IYaccNode
	raisesExpr       IYaccNode
}

func (rad readonlyAttrDeclarator) SimpleDeclarator() IYaccNode {
	return rad.simpleDeclarator
}

func (rad readonlyAttrDeclarator) RaisesExpr() IYaccNode {
	return rad.raisesExpr
}

func newReadonlyAttrDeclarator(simpleDeclarator, raisesExpr IYaccNode) (IAttrDeclarator, error) {
	return &readonlyAttrDeclarator{
		simpleDeclarator: simpleDeclarator,
		raisesExpr:       raisesExpr,
	}, nil
}
