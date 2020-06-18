package yaccidl

type UnionDcl struct {
	YaccNode
	forward    bool
	switchType IYaccNode
	unionBody  IUnionBody
}

func (u UnionDcl) UnionBody() IUnionBody {
	return u.unionBody
}

func (u UnionDcl) SwitchType() IYaccNode {
	return u.switchType
}

func (u UnionDcl) Forward() bool {
	return u.forward
}

func (u UnionDcl) IsUnionDcl() bool {
	return true
}

func NewUnionDcl(identifier IYaccNode, switchType IYaccNode, unionBody IUnionBody, forward bool) IUnionDcl {
	return &UnionDcl{
		YaccNode:   *NewYaccNode(identifier.Identifier(), identifier.LexemData()),
		forward:    forward,
		switchType: switchType,
		unionBody:  unionBody,
	}
}
