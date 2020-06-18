package yaccidl

type valueHeader struct {
	YaccNode
	valueKind            IInterfaceKind
	valueInheritanceSpec IValueInheritanceSpec
}

func (v valueHeader) ValueKind() IInterfaceKind {
	return v.valueKind
}

func (v valueHeader) ValueInheritanceSpec() IValueInheritanceSpec {
	return v.valueInheritanceSpec
}

func newValueHeader(valueKind IInterfaceKind, identifier IYaccNode, valueInheritanceSpec IValueInheritanceSpec) (IValueHeader, error) {
	return &valueHeader{
		YaccNode:             *NewYaccNode(identifier.Identifier(), identifier.LexemData()),
		valueKind:            valueKind,
		valueInheritanceSpec: valueInheritanceSpec,
	}, nil
}
