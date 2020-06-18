package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplyOperator(t *testing.T) {
	t.Run("int64(123) * int64(2)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '=', data: nil},
				{token: Integer_literal, data: int64(123)},
				{token: '*', data: nil},
				{token: Integer_literal, data: int64(2)},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IConstDcl)(nil), lexer.GetSpecification()) {
			return
		}
		constDcl, ok1 := lexer.GetSpecification().(IConstDcl)
		assert.True(t, ok1)
		int64Value, ok2 := constDcl.Value().(IInt64Value)
		assert.True(t, ok2)
		assert.Equal(t, int64(246), int64Value.Int64Value())
	})

	t.Run("\"123\" * \"123\"", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '=', data: nil},
				{token: String_literal, data: "123"},
				{token: '*', data: nil},
				{token: String_literal, data: "123"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.True(t, lexer.error) {
			return
		}
	})

	t.Run("float64(123) * float64(2);", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '=', data: nil},
				{token: Floating_pt_literal, data: float64(123)},
				{token: '*', data: nil},
				{token: Floating_pt_literal, data: float64(2)},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IConstDcl)(nil), lexer.GetSpecification()) {
			return
		}
		constDcl, ok1 := lexer.GetSpecification().(IConstDcl)
		assert.True(t, ok1)
		float64Value, ok2 := constDcl.Value().(IFloat64Value)
		assert.True(t, ok2)
		assert.Equal(t, float64(246), float64Value.FloatValue())
	})
}
