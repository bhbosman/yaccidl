package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStruct(t *testing.T) {
	t.Run("empty struct", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
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
	})

	t.Run("forward struct decl", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABCDEF"},
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
		structNode, ok := node.(IStructDcl)
		assert.True(t, ok)
		assert.True(t, structNode.Forward())
	})

	t.Run("struct with one type and one field", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: RWlong, data: nil},
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
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
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
		assert.Equal(t, "IdlLong", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("struct with one type and one field", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
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
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
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

	t.Run("struct with one type and 4 fields", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "A"},
				{token: ',', data: nil},
				{token: Identifier, data: "B"},
				{token: ',', data: nil},
				{token: Identifier, data: "C"},
				{token: ',', data: nil},
				{token: Identifier, data: "D"},
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
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
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
		assert.Equal(t, "IdlLong", member.TypeDcl().Identifier())

		decl01 := member.Declarator()
		if !assert.NotNil(t, decl01) {
			return
		}
		assert.Equal(t, "A", decl01.Identifier())
		assert.NotNil(t, decl01.GetNextNode())

		decl02 := decl01.GetNextNode()
		if !assert.NotNil(t, decl02) {
			return
		}
		assert.Equal(t, "B", decl02.Identifier())
		assert.NotNil(t, decl02.GetNextNode())

		decl03 := decl02.GetNextNode()
		if !assert.NotNil(t, decl03) {
			return
		}
		assert.Equal(t, "C", decl03.Identifier())
		assert.NotNil(t, decl03.GetNextNode())

		decl04 := decl03.GetNextNode()
		if !assert.NotNil(t, decl04) {
			return
		}
		assert.Equal(t, "D", decl04.Identifier())
		assert.Nil(t, decl04.GetNextNode())
	})

	t.Run("struct with two( type and one field)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: RWdouble, data: nil},
				{token: Identifier, data: "GHI"},
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
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}
		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member := structNode.GetMember()
		member01, ok := member.(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlLong", member01.TypeDcl().Identifier())
		assert.NotNil(t, member01.GetNextNode())

		decl01 := member01.Declarator()
		assert.NotNil(t, decl01)
		assert.Equal(t, "DEF", decl01.Identifier())
		assert.Nil(t, decl01.GetNextNode())

		member = member.GetNextNode()
		if !assert.NotNil(t, member) {
			return
		}
		member02, ok := member.(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlDoubleType", member02.TypeDcl().Identifier())
		assert.Nil(t, member02.GetNextNode())

		decl02 := member02.Declarator()
		assert.NotNil(t, decl02)
		assert.Equal(t, "GHI", decl02.Identifier())
		assert.Nil(t, decl01.GetNextNode())
	})

	t.Run("struct with two( type and two field)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "A"},
				{token: ',', data: nil},
				{token: Identifier, data: "B"},
				{token: ';', data: nil},
				{token: RWdouble, data: nil},
				{token: Identifier, data: "C"},
				{token: ',', data: nil},
				{token: Identifier, data: "D"},
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
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}
		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member := structNode.GetMember()
		member01, ok := member.(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlLong", member01.TypeDcl().Identifier())
		assert.NotNil(t, member01.GetNextNode())

		decl01 := member01.Declarator()
		assert.NotNil(t, decl01)
		assert.Equal(t, "A", decl01.Identifier())
		assert.NotNil(t, decl01.GetNextNode())

		decl02 := decl01.GetNextNode()
		assert.NotNil(t, decl02)
		assert.Equal(t, "B", decl02.Identifier())
		assert.Nil(t, decl02.GetNextNode())

		member = member.GetNextNode()
		if !assert.NotNil(t, member) {
			return
		}
		member02, ok := member.(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlDoubleType", member02.TypeDcl().Identifier())
		assert.Nil(t, member02.GetNextNode())

		decl03 := member02.Declarator()
		assert.NotNil(t, decl03)
		assert.Equal(t, "C", decl03.Identifier())
		assert.NotNil(t, decl03.GetNextNode())

		decl04 := decl03.GetNextNode()
		assert.NotNil(t, decl04)
		assert.Equal(t, "D", decl04.Identifier())
		assert.Nil(t, decl04.GetNextNode())
	})

	t.Run("inherited struct, no fields", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: ':', data: nil},
				{token: Identifier, data: "DEF"},
				{token: '{', data: nil},
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
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}
		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.Equal(t, "DEF", structNode.InheritsFrom().Identifier())

	})

	t.Run("inherited struct, with fields", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: ':', data: nil},
				{token: Identifier, data: "DEF"},
				{token: '{', data: nil},
				{token: RWlong, data: nil},
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
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}
		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.Equal(t, "DEF", structNode.InheritsFrom().Identifier())

		member, ok := structNode.GetMember().(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlLong", member.TypeDcl().Identifier())

		decl := member.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "DEF", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())

	})

	t.Run("Three", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
				{token: RWstruct, data: nil},
				{token: Identifier, data: "DEF"},
				{token: '{', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
				{token: RWstruct, data: nil},
				{token: Identifier, data: "GHI"},
				{token: '{', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
				{token: RWstruct, data: nil},
				{token: Identifier, data: "JKL"},
				{token: '{', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}

		node01 := lexer.GetSpecification()
		if !assert.NotNil(t, node01) {
			return
		}
		assert.Equal(t, "ABC", node01.Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node01) {
			return
		}
		assert.NotNil(t, node01.GetNextNode())

		node02 := node01.GetNextNode()
		if !assert.NotNil(t, node01) {
			return
		}
		assert.Equal(t, "DEF", node02.Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node02) {
			return
		}
		assert.NotNil(t, node02.GetNextNode())

		node03 := node02.GetNextNode()
		if !assert.NotNil(t, node03) {
			return
		}
		assert.Equal(t, "GHI", node03.Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node03) {
			return
		}
		assert.NotNil(t, node03.GetNextNode())

		node04 := node03.GetNextNode()
		if !assert.NotNil(t, node04) {
			return
		}
		assert.Equal(t, "JKL", node04.Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node04) {
			return
		}
		assert.Nil(t, node04.GetNextNode())

	})

	t.Run("struct with two( type and two field)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "A"},
				{token: '[', data: nil},
				{token: Integer_literal, data: 2},
				{token: ']', data: nil},

				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		if !assert.Equal(t, 0, parser.Parse(lexer)) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		if !assert.NotNil(t, node) {
			return
		}
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IStructDcl)(nil), node) {
			return
		}
		structNode := node.(IStructDcl)
		assert.Nil(t, structNode.GetNextNode())
		assert.NotNil(t, structNode.GetMember())
		if !assert.Implements(t, (*IStructMemberDcl)(nil), structNode.GetMember()) {
			return
		}
		member := structNode.GetMember()
		member01, ok := member.(IStructMemberDcl)
		assert.True(t, ok)
		assert.Equal(t, "IdlLong", member01.TypeDcl().Identifier())
		assert.Nil(t, member01.GetNextNode())

		decl01 := member01.Declarator()
		assert.NotNil(t, decl01)
		assert.Equal(t, "A", decl01.Identifier())
		assert.Nil(t, decl01.GetNextNode())

	})
}
