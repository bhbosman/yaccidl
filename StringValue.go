package yaccidl

type StringValue struct {
	ValueBase
	value string
}

func (i StringValue) String() string {
	return i.StringValue()
}

func (i StringValue) Value() interface{} {
	return i.value
}

func NewStringValue(value string) (IStringValue, error) {
	return &StringValue{
		value: value,
	}, nil
}

func (i StringValue) StringValue() string {
	return i.value
}
