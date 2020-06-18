package yaccidl

type IYaccNode interface {
	Identifier() string
	LexemData() *LexemValue
	SetNextNode(IYaccNode) error
	GetNextNode() IYaccNode
	AddArray(array []int64)
	IsArray() bool
	Array() []int64
}
