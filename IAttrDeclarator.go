package yaccidl

type IAttrDeclarator interface {
	SimpleDeclarator() IYaccNode
	RaisesExpr() IYaccNode
}
