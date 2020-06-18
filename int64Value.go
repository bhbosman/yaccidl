package yaccidl

import "fmt"

type Int64Value struct {
	ValueBase
	value int64
}

func (i Int64Value) String() string {
	return fmt.Sprintf("%d", i.Int64Value())
}

func (i Int64Value) Value() interface{} {
	return i.value
}

func (i Int64Value) Int64Value() int64 {
	return i.value
}

func NewInt64Value(value int64) (IInt64Value, error) {
	return &Int64Value{
		value: value,
	}, nil
}
