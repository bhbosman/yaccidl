package yaccidl

type IAssignSpecification interface {
	AssignSpecification(node IYaccNode)
	GetSpecification() IYaccNode
}
