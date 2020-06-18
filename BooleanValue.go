package yaccidl

type BooleanValue struct {
	ValueBase
	value bool
}

func (i BooleanValue) String() string {
	if i.BooleanValue() {
		return "true"
	}
	return "false"
}

func (i BooleanValue) Value() interface{} {
	return i.value
}

func (i BooleanValue) BooleanValue() bool {
	return i.value
}

func NewBooleanValue(value bool) (IBooleanValue, error) {
	return &BooleanValue{value: value}, nil
}
