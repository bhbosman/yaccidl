package yaccidl

import "github.com/bhbosman/yaccpragma"

type ITypePrefix interface {
	IYaccNode
	IsTypePrefix() bool
	Value() string
}

type TypePrefix struct {
	YaccNode
	value string
}

func (t TypePrefix) Value() string {
	return t.value
}

func (t TypePrefix) IsTypePrefix() bool {
	return true
}

func NewTypePrefix(lexemData *LexemValue, scopeName string, value string) ITypePrefix {
	return &TypePrefix{
		YaccNode: *NewYaccNode(scopeName, lexemData),
		value:    value,
	}
}

type IPragmaValue interface {
	IYaccNode
	IsIPragmaValue() bool
	Value() yaccpragma.IPragmaNode
}
