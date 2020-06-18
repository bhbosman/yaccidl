package yaccidl

import "fmt"

type float64Value struct {
	ValueBase
	value float64
}

func (i float64Value) String() string {
	return fmt.Sprintf("%f", i.FloatValue())
}

func (i float64Value) Value() interface{} {
	return i.value
}

func (i float64Value) FloatValue() float64 {
	return i.value
}

func newFloat64Value(value float64) (IFloat64Value, error) {
	return &float64Value{
		value: value,
	}, nil
}
