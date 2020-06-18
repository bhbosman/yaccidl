package yaccidl

type IOperationalParameter interface {
	IYaccNode
	Direction() ParamDirection
	ParamType() IYaccNode
}
