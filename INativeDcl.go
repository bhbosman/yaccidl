package yaccidl

type INativeDcl interface {
	IYaccNode
	IsNativeDcl() bool
}
