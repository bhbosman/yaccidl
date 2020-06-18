package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimitives(t *testing.T) {
	t.Run("short", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWshort, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlShort", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("int16", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWint16, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlInt16", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("int32", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWint32, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlInt32", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("int64", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWint64, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlInt64", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("unsigned short", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWunsigned, data: nil},
				{token: RWshort, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlUnsignedShort", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("uint16", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWuint16, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlUInt16", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("uint32", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWuint32, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlUInt32", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("uint64", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWuint64, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlUint64", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("charType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWchar, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlChar", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("wideCharType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWwchar, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlWideChar", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("stringType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWstring, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlString", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("wideStringType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWwstring, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlWideString", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("BooleanType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWboolean, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlBoolean", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("FloatType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWfloat, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlFloatType", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("DoubleType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWdouble, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlDoubleType", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("LongDoubleType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWlong, data: nil},
				{token: RWdouble, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlLongDoubleType", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("Int8Type", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWint8, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlInt8", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("UInt8Type", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWuint8, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlUInt8", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("AnyType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWany, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlAny", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("ObjectType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWObject, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlObject", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("ValueBaseType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWValueBase, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlValueBase", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})
	t.Run("OctetType", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWoctet, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}

		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlOctet", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

}
