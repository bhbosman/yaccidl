package yaccidl

type IStringType interface {
	IYaccNode
	StringLength() int64
}

type stringType struct {
	YaccNode
	length int64
}

func (s stringType) StringLength() int64 {
	return s.length
}

func newStringType(length int64) (IStringType, error) {
	return &stringType{
		YaccNode: *NewYaccNode("", nil),
		length:   length,
	}, nil
}
