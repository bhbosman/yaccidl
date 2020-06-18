package yaccidl

type IFloat64Value interface {
	IValue
	FloatValue() float64
}
