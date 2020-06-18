package yaccidl

type interfaceKind struct {
	YaccNode
	local    bool
	abstract bool
	custom   bool
}

func (i interfaceKind) IsnterfaceKind() bool {
	return true
}

func (i interfaceKind) Custom() bool {
	return i.custom
}

func (i interfaceKind) Local() bool {
	return i.local
}

func (i interfaceKind) Abstract() bool {
	return i.abstract
}

func newInterfaceKind(lexemData *LexemValue, local bool, abstract bool, custom bool) (IInterfaceKind, error) {
	return &interfaceKind{
		YaccNode: *NewYaccNode("(Interface Kind)", lexemData),
		local:    local,
		abstract: abstract,
		custom:   custom,
	}, nil
}
