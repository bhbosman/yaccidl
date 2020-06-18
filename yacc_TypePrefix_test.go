package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypeprefix, data: nil},
				{token: Identifier, data: "CORBA"},
				{token: String_literal, data: "omg.org"},
				{token: ';', data: nil},
			})
		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "CORBA", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*ITypePrefix)(nil), lexer.GetSpecification()) {
			return
		}
	})
}
