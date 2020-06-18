package yaccidl

type IWideStringType interface {
	IYaccNode
	WideStringLength() int64
}
