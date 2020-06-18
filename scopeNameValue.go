package yaccidl

type scopeNameValue struct {
	ValueBase
	value string
}

func (s scopeNameValue) String() string {
	return s.value
}

func (s scopeNameValue) Value() interface{} {
	return s.value
}

func (s scopeNameValue) ScopeName() string {
	return s.value
}

func newScopeNameValue(identifier string) IScopeNameValue {
	return &scopeNameValue{
		value: identifier,
	}
}
