package yaccidl

import "fmt"

type IElementSpec interface {
	TypeSpec() IYaccNode
	Declarator() IYaccNode
}

type DefaultValue struct {
	YaccNode
}

func (d *DefaultValue) SetNextConstValue(v IBaseConstantValue) error {
	return d.SetNextNode(v)
}

func (d *DefaultValue) GetNextConstValue() (IBaseConstantValue, error) {
	if v, ok := d.GetNextNode().(IBaseConstantValue); ok {
		return v, nil
	}
	return nil, fmt.Errorf("net is not a IBaseConstantValue")
}

func (d DefaultValue) IsIBaseConstantValue() bool {
	return true
}

func (d DefaultValue) IsDefaultValue() bool {
	return true
}

type IBaseConstantValue interface {
	IYaccNode
	IsIBaseConstantValue() bool
	SetNextConstValue(IBaseConstantValue) error
	GetNextConstValue() (IBaseConstantValue, error)
}

type IDefaultValue interface {
	IBaseConstantValue
	IsDefaultValue() bool
}

type IConstValue interface {
	IBaseConstantValue
	IsConstValue() bool
	Value() IValue
}

type ConstValue struct {
	YaccNode
	value IValue
}

func (c ConstValue) IsIBaseConstantValue() bool {
	return true
}

func (c *ConstValue) SetNextConstValue(v IBaseConstantValue) error {
	return c.SetNextNode(v)
}

func (c *ConstValue) GetNextConstValue() (IBaseConstantValue, error) {
	if v, ok := c.GetNextNode().(IBaseConstantValue); ok {
		return v, nil
	}
	return nil, fmt.Errorf("net is not a IBaseConstantValue")
}

func (c *ConstValue) IsConstValue() bool {
	return true
}

func (c ConstValue) Value() IValue {
	return c.value
}

func NewConstValue(id string, lexemData *LexemValue, value IValue) IConstValue {
	return &ConstValue{
		YaccNode: *NewYaccNode(id, lexemData),
		value:    value,
	}
}

func NewDefaultValue(id string, lexemData *LexemValue) IDefaultValue {
	return &DefaultValue{
		YaccNode: *NewYaccNode(id, lexemData),
	}
}

type IUnionDcl interface {
	IYaccNode
	Forward() bool
	IsUnionDcl() bool
	SwitchType() IYaccNode
	UnionBody() IUnionBody
}

type IUnionBody interface {
	IYaccNode
	IsUnionBody() bool
	CaseLabels() IBaseConstantValue
	ElementSpec() IElementSpec
	SetNextUnionBody(IUnionBody) error
	GetNextUnionBody() (IUnionBody, error)
}

type UnionBody struct {
	YaccNode
	caseLabels  IBaseConstantValue
	elementSpec IElementSpec
}

func (u *UnionBody) SetNextUnionBody(v IUnionBody) error {
	return u.SetNextNode(v)
}

func (u *UnionBody) GetNextUnionBody() (IUnionBody, error) {
	if v, ok := u.GetNextNode().(IUnionBody); ok {
		return v, nil
	}
	return nil, fmt.Errorf("net is not a IUnionBody")
}

func (u UnionBody) CaseLabels() IBaseConstantValue {
	return u.caseLabels
}

func (u UnionBody) ElementSpec() IElementSpec {
	return u.elementSpec
}

func (u UnionBody) IsUnionBody() bool {
	return true
}

func NewUnionBody(lexemData *LexemValue, identifier string, caseLabels IBaseConstantValue, elementSpec IElementSpec) IUnionBody {
	return &UnionBody{
		YaccNode:    *NewYaccNode(identifier, lexemData),
		caseLabels:  caseLabels,
		elementSpec: elementSpec,
	}
}

type ElementSpec struct {
	typeSpec   IYaccNode
	declarator IYaccNode
}

func (e ElementSpec) TypeSpec() IYaccNode {
	return e.typeSpec
}

func (e ElementSpec) Declarator() IYaccNode {
	return e.declarator
}

func NewElementSpec(typeSpec, declarator IYaccNode) IElementSpec {
	return &ElementSpec{
		typeSpec:   typeSpec,
		declarator: declarator,
	}
}
