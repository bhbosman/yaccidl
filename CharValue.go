package yaccidl

type CharValue struct {
	ValueBase
	value rune
}

func (i CharValue) String() string {
	return string(i.value)
}

func (i CharValue) Value() interface{} {
	return i.value
}

func (i CharValue) CharValue() rune {
	return i.value
}

func NewCharValue(value rune) (ICharValue, error) {
	return &CharValue{value: value}, nil
}
