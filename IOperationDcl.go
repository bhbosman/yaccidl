package yaccidl

type IOperationDcl interface {
	IYaccNode
	ParamDcl() IYaccNode
	ReturnType() IYaccNode
	ExceptionDcl() IYaccNode
}
