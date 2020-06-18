package yaccidl

type wideStringType struct {
	YaccNode
	length int64
}

func (w wideStringType) WideStringLength() int64 {
	return w.length
}

func newWideStringType(length int64) (IWideStringType, error) {
	return &wideStringType{
		YaccNode: *NewYaccNode("", nil),
		length:   length,
	}, nil
}
