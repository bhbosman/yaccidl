package yaccidl

import (
	"encoding/json"
	"strconv"
	"strings"
)

type PrimitiveType byte

func (p *PrimitiveType) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	if v, ok := reversedPrimitiveLookup[s]; ok {
		*p = v
		return nil
	}
	*p = IdlInvalid
	return nil
}

func (p *PrimitiveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p PrimitiveType) String() string {
	if int(p) < len(PrimitiveTypeNames) {
		return PrimitiveTypeNames[p]
	}
	return "kind" + strconv.Itoa(int(p))
}

const (
	IdlInvalid PrimitiveType = iota
	ShortType
	Int16Type
	Int32Type
	Int64Type
	UnsignedShortType
	Uint16Type
	Uint32Type
	Uint64Type
	LongType
	LongLongType
	UnsignedLongType
	UnsignedLongLongType
	CharType
	WideCharType
	StringType
	WideStringType
	BooleanType
	FloatType
	DoubleType
	LongDoubleType
	VoidType
	Int8Type
	UInt8Type
	AnyType
	IdlObjectKind
	IdlValueBaseKind
	IdlOctetKind
	IdlConst
	IdlEnum
	IdlInterface
	IdlInterfaceContainer
	IdlModule
	IdlNative
	IdlScope
	IdlSequence
	IdlStruct
	IdlTypedef
	IdlTypePrefix
	IdlUnion
)

var PrimitiveTypeNames = []string{
	ShortType:            "ShortType",
	Int16Type:            "Int16Type",
	Int32Type:            "Int32Type",
	Int64Type:            "Int64Type",
	UnsignedShortType:    "UnsignedShortType",
	Uint16Type:           "Uint16Type",
	Uint32Type:           "Uint32Type",
	Uint64Type:           "Uint64Type",
	LongType:             "LongType",
	LongLongType:         "LongLongType",
	UnsignedLongType:     "UnsignedLongType",
	UnsignedLongLongType: "UnsignedLongLongType",
	CharType:             "CharType",
	WideCharType:         "WideCharType",
	StringType:           "StringType",
	WideStringType:       "WideStringType",
	BooleanType:          "BooleanType",
	FloatType:            "FloatType",
	DoubleType:           "DoubleType",
	LongDoubleType:       "LongDoubleType",
	VoidType:             "VoidType",
	Int8Type:             "Int8Type",
	UInt8Type:            "UInt8Type",
	AnyType:              "AnyType",
	IdlObjectKind:        "IdlObjectKind",
	IdlValueBaseKind:     "IdlValueBaseKind",
	IdlOctetKind:         "IdlOctetKind",
	IdlInvalid:           "IdlInvalid",
	IdlConst:             "IdlConst",
	IdlEnum:              "IdlEnum",
	IdlInterface:         "IdlInterface",
	IdlModule:            "IdlModule",
	IdlNative:            "IdlNative",
	IdlScope:             "IdlScope",
	IdlSequence:          "IdlSequence",
	IdlStruct:            "IdlStruct",
	IdlTypedef:           "IdlTypedef",
	IdlTypePrefix:        "IdlTypePrefix",
	IdlUnion:             "IdlUnion",
}

type IPrimitiveTypeDcl interface {
	IYaccNode
	PrimitiveType() PrimitiveType
}

var reversedPrimitiveLookup map[string]PrimitiveType

func init() {
	reversedPrimitiveLookup = make(map[string]PrimitiveType)
	for i, v := range PrimitiveTypeNames {
		reversedPrimitiveLookup[v] = PrimitiveType(i)

	}

}
