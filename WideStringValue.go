package yaccidl

import "fmt"

type WideStringValue struct {
	ValueBase
	value IYaccNode
}

func (i WideStringValue) String() string {
	return fmt.Sprintf("%v", i.WideStringValue())
}

func (i WideStringValue) Value() interface{} {
	return i.value
}

func NewWideStringValue(value IYaccNode) (IWideStringValue, error) {
	return &WideStringValue{value: value}, nil
}

func (i WideStringValue) WideStringValue() string {
	return i.value.Identifier()
}
