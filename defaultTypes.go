package yaccidl

type IIdlShortType interface {
	ShortType() IYaccNode
}

type IIdlInt16Type interface {
	Int16Type() IYaccNode
}

type IIdlInt32Type interface {
	Int32Type() IYaccNode
}

type IIdlInt64Type interface {
	Int64Type() IYaccNode
}

type IIdlUnsignedShortType interface {
	UnsignedShortType() IYaccNode
}

type IIdlUint16Type interface {
	Uint16Type() IYaccNode
}

type IIdlUint32Type interface {
	Uint32Type() IYaccNode
}

type IIdlUint64Type interface {
	Uint64Type() IYaccNode
}

type IIdlLongType interface {
	LongType() IYaccNode
}

type IIdlLongLongType interface {
	LongLongType() IYaccNode
}

type IIdlUnsignedLongType interface {
	UnsignedLongType() IYaccNode
}

type IIdlUnsignedLongLongType interface {
	UnsignedLongLongType() IYaccNode
}

type IIdlCharType interface {
	CharType() IYaccNode
}

type IIdlWideCharType interface {
	WideCharType() IYaccNode
}

type IIdlStringType interface {
	StringType() IYaccNode
}

type IIdlWideStringType interface {
	WideStringType() IYaccNode
}

type IIdlBooleanType interface {
	BooleanType() IYaccNode
}

type IIdlFloatType interface {
	FloatType() IYaccNode
}

type IIdlDoubleType interface {
	DoubleType() IYaccNode
}

type IIdlLongDoubleType interface {
	LongDoubleType() IYaccNode
}

type IIdlVoidType interface {
	VoidType() IYaccNode
}

type IIdlInt8Type interface {
	Int8Type() IYaccNode
}

type IIdlUInt8Type interface {
	UInt8Type() IYaccNode
}

type IIdlAnyType interface {
	AnyType() IYaccNode
}

type IIdlObjectType interface {
	ObjectType() IYaccNode
}

type IIdlValueBaseType interface {
	ValueBaseType() IYaccNode
}

type IIdlOctetType interface {
	OctetType() IYaccNode
}
