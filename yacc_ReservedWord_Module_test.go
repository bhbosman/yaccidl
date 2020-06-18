package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModule(t *testing.T) {
	t.Run("Empty module", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWmodule, data: nil},
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
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		node := lexer.GetSpecification()
		if !assert.Implements(t, (*IModuleDcl)(nil), node) {
			return
		}
		module, ok1 := lexer.GetSpecification().(IModuleDcl)
		assert.True(t, ok1)
		assert.Nil(t, module.ChildDecls())
	})

	t.Run("module with two const decls", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWmodule, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: RWconst, data: nil},
				{token: RWstring, data: nil},
				{token: Identifier, data: "DEF"},
				{token: '=', data: nil},
				{token: String_literal, data: "GHI"},
				{token: ';', data: nil},
				{token: RWconst, data: nil},
				{token: RWstring, data: nil},
				{token: Identifier, data: "JKL"},
				{token: '=', data: nil},
				{token: String_literal, data: "MNO"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())

		node := lexer.GetSpecification()
		if !assert.Implements(t, (*IModuleDcl)(nil), node) {
			return
		}

		module, ok1 := lexer.GetSpecification().(IModuleDcl)
		assert.True(t, ok1)
		assert.NotNil(t, module.ChildDecls())
		assert.Nil(t, module.GetNextNode())

		node = module.ChildDecls()
		if !assert.Implements(t, (*IConstDcl)(nil), node) {
			return
		}

		node = node.GetNextNode()
		if !assert.Implements(t, (*IConstDcl)(nil), node) {
			return
		}
	})

	t.Run("Empty module, with struct with scope in", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWmodule, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: ScopeOp, data: nil},
				{token: Identifier, data: "ZZZ"},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},

				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		node := lexer.GetSpecification()
		if !assert.Implements(t, (*IModuleDcl)(nil), node) {
			return
		}
		_, ok1 := lexer.GetSpecification().(IModuleDcl)
		assert.True(t, ok1)
	})

	t.Run("Empty module, with struct with scope in", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWmodule, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWstruct, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},

				{token: ScopeOp, data: nil},
				{token: Identifier, data: "ZZZ"},
				{token: Identifier, data: "DEF"},
				{token: ';', data: nil},
				{token: '}', data: nil},
				{token: ';', data: nil},

				{token: '}', data: nil},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		node := lexer.GetSpecification()
		if !assert.Implements(t, (*IModuleDcl)(nil), node) {
			return
		}
		_, ok1 := lexer.GetSpecification().(IModuleDcl)
		assert.True(t, ok1)
	})

}
