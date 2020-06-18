package yaccidl

type primitiveType struct {
	YaccNode
	primitiveType PrimitiveType
	isNative      bool
	nativeMapping string
}

func (t *primitiveType) NativeMapping() string {
	return t.nativeMapping
}

func (t *primitiveType) IsNative() bool {
	return t.isNative
}

func (t *primitiveType) PrimitiveType() PrimitiveType {
	return t.primitiveType
}

func (t *primitiveType) IsPrimitive() bool {
	return true
}

func NewPrimitiveNode(
	idlReference string,
	PrimitiveType PrimitiveType,
	lexemData *LexemValue) IPrimitiveTypeDcl {
	return NewPrimitiveNodeFull(
		idlReference,
		PrimitiveType,
		false,
		"",
		lexemData)
}

func NewPrimitiveNodeFull(
	idlReference string,
	PrimitiveType PrimitiveType,
	isNative bool,
	nativeMapping string,
	lexemData *LexemValue) IPrimitiveTypeDcl {
	return &primitiveType{
		YaccNode:      *NewYaccNode(idlReference, lexemData),
		primitiveType: PrimitiveType,
		isNative:      isNative,
		nativeMapping: nativeMapping,
	}
}
