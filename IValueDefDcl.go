package yaccidl

type IValueDefDcl interface {
	IYaccNode
	ValueHeader() IValueHeader
	Body() IYaccNode
	Forward() bool
}
