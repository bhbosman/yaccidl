package yaccidl

type WideCharValue struct {
	ValueBase
	value rune
}

func (i WideCharValue) String() string {
	return string(i.value)
}

func (i WideCharValue) Value() interface{} {
	return i.value
}

func (i WideCharValue) WideCharValue() rune {
	return i.value
}

func NewWideCharValue(value rune) (IWideCharValue, error) {
	return &WideCharValue{value: value}, nil
}
