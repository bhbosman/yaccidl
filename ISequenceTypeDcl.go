package yaccidl

import "fmt"

type ISequenceTypeDcl interface {
	IYaccNode
	TypeSpec() IYaccNode
	ArraySize() int64
}

type sequenceType struct {
	YaccNode
	typeSpec  IYaccNode
	arraySize int64
}

func (s sequenceType) String() string {
	return fmt.Sprintf("%v,%v", s.typeSpec, s.arraySize)
}

func (s sequenceType) TypeSpec() IYaccNode {
	return s.typeSpec
}

func (s sequenceType) ArraySize() int64 {
	return s.arraySize
}

func newSequenceType(identifier string, lexemData *LexemValue, typeSpec IYaccNode, arraySize int64) (ISequenceTypeDcl, error) {
	return &sequenceType{
		YaccNode:  *NewYaccNode(identifier, lexemData),
		typeSpec:  typeSpec,
		arraySize: arraySize,
	}, nil
}
