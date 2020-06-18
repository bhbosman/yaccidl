package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExceptions(t *testing.T) {
	t.Run("exception with one type and one field", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWexception, data: nil},
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
		if !assert.Implements(t, (*IExceptionDcl)(nil), node) {
			return
		}
		structNode := node.(IExceptionDcl)
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
}
