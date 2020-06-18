package yaccidl

type IValueInheritanceSpec interface {
	IYaccNode
	SupportedInterfaceName() IYaccNode
	ValueName() IYaccNode
}

type valueInheritanceSpec struct {
	YaccNode
	supportedInterfaceName IYaccNode
	valueName              IYaccNode
}

func (v valueInheritanceSpec) SupportedInterfaceName() IYaccNode {
	return v.supportedInterfaceName
}

func (v valueInheritanceSpec) ValueName() IYaccNode {
	return v.valueName
}

func newValueInheritanceSpec(supportedInterfaceName IYaccNode, valueName IYaccNode) (IValueInheritanceSpec, error) {
	return &valueInheritanceSpec{
		YaccNode:               *NewYaccNode("", nil),
		supportedInterfaceName: supportedInterfaceName,
		valueName:              valueName,
	}, nil
}
