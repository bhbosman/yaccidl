package yaccidl

import (
	"os"
	"path/filepath"
)

type lexem struct {
	token int
	data  interface{}
}

type YaccIdlLexerImpl struct {
	pos  int
	data []lexem
	//t            *testing.T
	error                bool
	errorMessage         string
	specificationNode    IYaccNode
	shortType            IYaccNode
	int16Type            IYaccNode
	int32Type            IYaccNode
	int64Type            IYaccNode
	unsignedShortType    IYaccNode
	uint16Type           IYaccNode
	uint32Type           IYaccNode
	uint64Type           IYaccNode
	longType             IYaccNode
	longLongType         IYaccNode
	unsignedLongType     IYaccNode
	unsignedLongLongType IYaccNode
	charType             IYaccNode
	wideCharType         IYaccNode
	stringType           IYaccNode
	wideStringType       IYaccNode
	booleanType          IYaccNode
	floatType            IYaccNode
	doubleType           IYaccNode
	longDoubleType       IYaccNode
	voidType             IYaccNode
	int8Type             IYaccNode
	uInt8Type            IYaccNode
	anyType              IYaccNode
	objectType           IYaccNode
	valueBaseType        IYaccNode
	octetType            IYaccNode
	definition           IYaccNode
}

func (y *YaccIdlLexerImpl) AssignDefinition(node IYaccNode) {
	y.definition = node
}

func (y *YaccIdlLexerImpl) GetDefinition() IYaccNode {
	return y.definition
}

func (y *YaccIdlLexerImpl) BooleanType() IYaccNode {
	return y.booleanType
}

func (y *YaccIdlLexerImpl) FloatType() IYaccNode {
	return y.floatType
}

func (y *YaccIdlLexerImpl) DoubleType() IYaccNode {
	return y.doubleType
}

func (y *YaccIdlLexerImpl) LongDoubleType() IYaccNode {
	return y.longDoubleType
}

func (y *YaccIdlLexerImpl) VoidType() IYaccNode {
	return y.voidType
}

func (y *YaccIdlLexerImpl) Int8Type() IYaccNode {
	return y.int8Type
}

func (y *YaccIdlLexerImpl) UInt8Type() IYaccNode {
	return y.uInt8Type
}

func (y *YaccIdlLexerImpl) AnyType() IYaccNode {
	return y.anyType
}

func (y *YaccIdlLexerImpl) ObjectType() IYaccNode {
	return y.objectType
}

func (y *YaccIdlLexerImpl) ValueBaseType() IYaccNode {
	return y.valueBaseType
}

func (y *YaccIdlLexerImpl) OctetType() IYaccNode {
	return y.octetType
}

func (y *YaccIdlLexerImpl) CharType() IYaccNode {
	return y.charType
}

func (y *YaccIdlLexerImpl) WideCharType() IYaccNode {
	return y.wideCharType
}

func (y *YaccIdlLexerImpl) StringType() IYaccNode {
	return y.stringType
}

func (y *YaccIdlLexerImpl) WideStringType() IYaccNode {
	return y.wideStringType
}

func (y *YaccIdlLexerImpl) LongType() IYaccNode {
	return y.longType
}

func (y *YaccIdlLexerImpl) LongLongType() IYaccNode {
	return y.longLongType
}

func (y *YaccIdlLexerImpl) UnsignedLongType() IYaccNode {
	return y.unsignedLongType
}

func (y *YaccIdlLexerImpl) UnsignedLongLongType() IYaccNode {
	return y.unsignedLongLongType
}

func (y *YaccIdlLexerImpl) UnsignedShortType() IYaccNode {
	return y.unsignedShortType
}

func (y *YaccIdlLexerImpl) Uint16Type() IYaccNode {
	return y.uint16Type
}

func (y *YaccIdlLexerImpl) Uint32Type() IYaccNode {
	return y.uint32Type
}

func (y *YaccIdlLexerImpl) Uint64Type() IYaccNode {
	return y.uint64Type
}

func (y *YaccIdlLexerImpl) Int16Type() IYaccNode {
	return y.int16Type
}

func (y *YaccIdlLexerImpl) Int32Type() IYaccNode {
	return y.int32Type
}

func (y *YaccIdlLexerImpl) Int64Type() IYaccNode {
	return y.int64Type
}

func (y *YaccIdlLexerImpl) ShortType() IYaccNode {
	return y.shortType
}

func NewYaccIdlLexerImpl(pos int, data []lexem) *YaccIdlLexerImpl {
	goSrcPath := filepath.Join(os.Getenv("GOPATH"), "./src")
	LexemDataForPrimitives := NewLexemDataForPrimitives("", "", filepath.Join(goSrcPath, "github.com/bhbosman/orb/src/omg"), 0, 0, "omg")

	return &YaccIdlLexerImpl{
		pos:                  pos,
		data:                 data,
		shortType:            NewPrimitiveNode("IdlShort", ShortType, LexemDataForPrimitives),
		int16Type:            NewPrimitiveNode("IdlInt16", Int16Type, LexemDataForPrimitives),
		int32Type:            NewPrimitiveNode("IdlInt32", Int32Type, LexemDataForPrimitives),
		int64Type:            NewPrimitiveNode("IdlInt64", Int64Type, LexemDataForPrimitives),
		unsignedShortType:    NewPrimitiveNode("IdlUnsignedShort", UnsignedShortType, LexemDataForPrimitives),
		uint16Type:           NewPrimitiveNode("IdlUInt16", Uint16Type, LexemDataForPrimitives),
		uint32Type:           NewPrimitiveNode("IdlUInt32", Uint32Type, LexemDataForPrimitives),
		uint64Type:           NewPrimitiveNode("IdlUint64", Uint64Type, LexemDataForPrimitives),
		longType:             NewPrimitiveNode("IdlLong", LongType, LexemDataForPrimitives),
		longLongType:         NewPrimitiveNode("IdlLongLong", LongLongType, LexemDataForPrimitives),
		unsignedLongType:     NewPrimitiveNode("IdlUnsignedLong", UnsignedLongType, LexemDataForPrimitives),
		unsignedLongLongType: NewPrimitiveNode("IdlUnsignedLongLong", UnsignedLongLongType, LexemDataForPrimitives),
		charType:             NewPrimitiveNode("IdlChar", CharType, LexemDataForPrimitives),
		wideCharType:         NewPrimitiveNode("IdlWideChar", WideCharType, LexemDataForPrimitives),
		stringType:           NewPrimitiveNode("IdlString", StringType, LexemDataForPrimitives),
		wideStringType:       NewPrimitiveNode("IdlWideString", WideStringType, LexemDataForPrimitives),
		booleanType:          NewPrimitiveNode("IdlBoolean", BooleanType, LexemDataForPrimitives),
		floatType:            NewPrimitiveNode("IdlFloatType", FloatType, LexemDataForPrimitives),
		doubleType:           NewPrimitiveNode("IdlDoubleType", DoubleType, LexemDataForPrimitives),
		longDoubleType:       NewPrimitiveNode("IdlLongDoubleType", LongDoubleType, LexemDataForPrimitives),
		voidType:             NewPrimitiveNode("IdlVoidType", VoidType, LexemDataForPrimitives),
		int8Type:             NewPrimitiveNode("IdlInt8", Int8Type, LexemDataForPrimitives),
		uInt8Type:            NewPrimitiveNode("IdlUInt8", UInt8Type, LexemDataForPrimitives),
		anyType:              NewPrimitiveNode("IdlAny", AnyType, LexemDataForPrimitives),
		objectType:           NewPrimitiveNode("IdlObject", IdlObjectKind, LexemDataForPrimitives),
		valueBaseType:        NewPrimitiveNode("IdlValueBase", IdlValueBaseKind, LexemDataForPrimitives),
		octetType:            NewPrimitiveNode("IdlOctet", IdlOctetKind, LexemDataForPrimitives),
	}
}

func (y *YaccIdlLexerImpl) AssignSpecification(node IYaccNode) {
	y.specificationNode = node
}

func (y YaccIdlLexerImpl) GetSpecification() IYaccNode {
	return y.specificationNode
}

func (y *YaccIdlLexerImpl) Lex(lval *YaccIdlSymType) int {
	for {
		y.pos++
		if y.pos == len(y.data)+1 {
			return 0
		}
		switch y.data[y.pos-1].token {
		case HashPragma:
			if v, ok := y.data[y.pos-1].data.(IYaccNode); ok {
				y.AssignDefinition(v)
			}
			continue

		case Identifier, String_literal, Wide_String_literal:
			if s, ok := y.data[y.pos-1].data.(string); ok {
				lval.Node = NewYaccNode(s, nil)
			}
			return y.data[y.pos-1].token
		case Integer_literal:
			if v, ok := y.data[y.pos-1].data.(int64); ok {
				lval.Value, _ = NewInt64Value(v)
			} else if v, ok := y.data[y.pos-1].data.(int); ok {
				lval.Value, _ = NewInt64Value(int64(v))
			} else if v, ok := y.data[y.pos-1].data.(int32); ok {
				lval.Value, _ = NewInt64Value(int64(v))
			} else if v, ok := y.data[y.pos-1].data.(int16); ok {
				lval.Value, _ = NewInt64Value(int64(v))
			} else if v, ok := y.data[y.pos-1].data.(int8); ok {
				lval.Value, _ = NewInt64Value(int64(v))
			} else if v, ok := y.data[y.pos-1].data.(uint64); ok {
				lval.Value, _ = NewInt64Value(int64(v))
			} else if v, ok := y.data[y.pos-1].data.(uint32); ok {
				lval.Value, _ = NewInt64Value(int64(v))
			} else if v, ok := y.data[y.pos-1].data.(uint16); ok {
				lval.Value, _ = NewInt64Value(int64(v))
			} else if v, ok := y.data[y.pos-1].data.(uint8); ok {
				lval.Value, _ = NewInt64Value(int64(v))
			}
			return y.data[y.pos-1].token
		case Floating_pt_literal:
			if v, ok := y.data[y.pos-1].data.(float64); ok {
				lval.Value, _ = newFloat64Value(v)
			}
			return y.data[y.pos-1].token

		}
		lval.Node = NewYaccNode("", &LexemValue{})
		return y.data[y.pos-1].token
	}

}

func (y *YaccIdlLexerImpl) Error(s string) {
	y.error = true
	y.errorMessage = s
}
