package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYaccRwConst(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '=', data: nil},
				{token: Integer_literal, data: int64(123)},
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
	})

	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: RWstring, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '=', data: nil},
				{token: String_literal, data: "ABCDEF"},
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
	})

	t.Run("Two constant", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: RWstring, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '=', data: nil},
				{token: String_literal, data: "ABCDEF"},
				{token: ';', data: nil},
				{token: RWconst, data: nil},
				{token: RWstring, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '=', data: nil},
				{token: String_literal, data: "ABCDEF"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		node := lexer.GetSpecification()
		assert.NotNil(t, node)
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IConstDcl)(nil), node) {
			return
		}

		node = node.GetNextNode()
		assert.NotNil(t, node)
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IConstDcl)(nil), node) {
			return
		}
	})

	t.Run("Two constant", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: RWstring, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '=', data: nil},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}

		node := lexer.GetSpecification()
		assert.NotNil(t, node)
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IConstDcl)(nil), node) {
			return
		}
		constDcl := node.(IConstDcl)
		if !assert.NotNil(t, constDcl) {
			return
		}
		if !assert.Implements(t, (*IScopeNameValue)(nil), constDcl.Value()) {
			return
		}
		scopeNameValue := constDcl.Value().(IScopeNameValue)
		assert.Equal(t, "DEF", scopeNameValue.ScopeName())

	})

	t.Run("Two constant", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: Identifier, data: "type"},
				{token: Identifier, data: "value"},
				{token: '=', data: nil},
				{token: Identifier, data: "12"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}

		node := lexer.GetSpecification()
		assert.NotNil(t, node)
		assert.Equal(t, "value", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IConstDcl)(nil), node) {
			return
		}
		constDcl := node.(IConstDcl)
		if !assert.NotNil(t, constDcl) {
			return
		}

		if !assert.Implements(t, (*IScopeNameDcl)(nil), constDcl.TypeDef()) {
			return
		}

		if !assert.Implements(t, (*IScopeNameValue)(nil), constDcl.Value()) {
			return
		}
		scopeNameValue := constDcl.Value().(IScopeNameValue)
		assert.Equal(t, "12", scopeNameValue.ScopeName())

	})

	t.Run("Two constant", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: Identifier, data: "::type"},
				{token: Identifier, data: "value"},
				{token: '=', data: nil},
				{token: Identifier, data: "12"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}

		node := lexer.GetSpecification()
		assert.NotNil(t, node)
		assert.Equal(t, "value", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IConstDcl)(nil), node) {
			return
		}
		constDcl := node.(IConstDcl)
		if !assert.NotNil(t, constDcl) {
			return
		}

		if !assert.Implements(t, (*IScopeNameDcl)(nil), constDcl.TypeDef()) {
			return
		}

		if !assert.Implements(t, (*IScopeNameValue)(nil), constDcl.Value()) {
			return
		}
		scopeNameValue := constDcl.Value().(IScopeNameValue)
		assert.Equal(t, "12", scopeNameValue.ScopeName())

	})

	t.Run("Two constant", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWconst, data: nil},
				{token: Identifier, data: "AA::type"},
				{token: Identifier, data: "value"},
				{token: '=', data: nil},
				{token: Identifier, data: "12"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}

		node := lexer.GetSpecification()
		assert.NotNil(t, node)
		assert.Equal(t, "value", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IConstDcl)(nil), node) {
			return
		}
		constDcl := node.(IConstDcl)
		if !assert.NotNil(t, constDcl) {
			return
		}

		if !assert.Implements(t, (*IScopeNameDcl)(nil), constDcl.TypeDef()) {
			return
		}

		if !assert.Implements(t, (*IScopeNameValue)(nil), constDcl.Value()) {
			return
		}
		scopeNameValue := constDcl.Value().(IScopeNameValue)
		assert.Equal(t, "12", scopeNameValue.ScopeName())

	})

}
