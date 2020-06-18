package yaccidl

type IStructDcl interface {
	IBaseExceptionDcl
	Forward() bool
	InheritsFrom() IYaccNode
	IsStructDcl() bool
}
