package yaccidl

import "fmt"

type IValue interface {
	fmt.Stringer
	Value() interface{}
	SetNextNode(IValue) error
	GetNextNode() IValue
}

type ValueBase struct {
	nextNode IValue
}

func (v *ValueBase) SetNextNode(value IValue) error {
	v.nextNode = value
	return nil
}

func (v *ValueBase) GetNextNode() IValue {
	return v.nextNode
}
