package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportDcl_IsImportDcl(t *testing.T) {

	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWchar, data: nil},
				{token: Identifier, data: "a"},
				{token: '[', data: nil},
				{token: Integer_literal, data: 4},
				{token: ']', data: nil},
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

		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), node) {
			return
		}
		typeDeclaratorDcl, _ := node.(ITypeDeclaratorDcl)
		decl := typeDeclaratorDcl.Declarator()
		if !assert.Equal(t, "a", decl.Identifier()) {
			return
		}

		if !assert.True(t, decl.IsArray()) {
			return
		}
		if !assert.Len(t, decl.Array(), 1) {
			return
		}
		if !assert.Equal(t, int64(4), decl.Array()[0]) {
			return
		}
	})

	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWchar, data: nil},
				{token: Identifier, data: "a"},
				{token: '[', data: nil},
				{token: Integer_literal, data: 4},
				{token: ']', data: nil},
				{token: '[', data: nil},
				{token: Integer_literal, data: 4},
				{token: ']', data: nil},
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

		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), node) {
			return
		}
		typeDeclaratorDcl, _ := node.(ITypeDeclaratorDcl)
		decl := typeDeclaratorDcl.Declarator()
		if !assert.Equal(t, "a", decl.Identifier()) {
			return
		}

		if !assert.True(t, decl.IsArray()) {
			return
		}
		if !assert.Len(t, decl.Array(), 2) {
			return
		}
		if !assert.Equal(t, int64(4), decl.Array()[0]) {
			return
		}
		if !assert.Equal(t, int64(4), decl.Array()[1]) {
			return
		}
	})

	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWchar, data: nil},
				{token: Identifier, data: "a"},
				{token: '[', data: nil},
				{token: Integer_literal, data: 1},
				{token: ']', data: nil},
				{token: '[', data: nil},
				{token: Integer_literal, data: 2},
				{token: ']', data: nil},
				{token: '[', data: nil},
				{token: Integer_literal, data: 3},
				{token: ']', data: nil},
				{token: '[', data: nil},
				{token: Integer_literal, data: 4},
				{token: ']', data: nil},
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

		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), node) {
			return
		}
		typeDeclaratorDcl, _ := node.(ITypeDeclaratorDcl)
		decl := typeDeclaratorDcl.Declarator()
		if !assert.Equal(t, "a", decl.Identifier()) {
			return
		}

		if !assert.True(t, decl.IsArray()) {
			return
		}
		if !assert.Len(t, decl.Array(), 4) {
			return
		}
		if !assert.Equal(t, int64(1), decl.Array()[0]) {
			return
		}
		if !assert.Equal(t, int64(2), decl.Array()[1]) {
			return
		}
		if !assert.Equal(t, int64(3), decl.Array()[2]) {
			return
		}
		if !assert.Equal(t, int64(4), decl.Array()[3]) {
			return
		}
	})

	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWimport, data: nil},
				{token: Identifier, data: "::CORBA"},
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

		if !assert.Implements(t, (*IImportDcl)(nil), node) {
			return
		}
		importDcl, _ := node.(IImportDcl)

		if !assert.Equal(t, "::CORBA", importDcl.Identifier()) {
			return
		}
		if !assert.Implements(t, (*IScopeNameDcl)(nil), importDcl.ImportedScope()) {
			return
		}
		scopeNameDcl, _ := importDcl.ImportedScope().(IScopeNameDcl)
		if !assert.Equal(t, "::CORBA", scopeNameDcl.Identifier()) {
			return
		}
	})

	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWimport, data: nil},
				{token: String_literal, data: "::CORBA"},
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

		if !assert.Implements(t, (*IImportDcl)(nil), node) {
			return
		}
		importDcl, _ := node.(IImportDcl)

		if !assert.Equal(t, "::CORBA", importDcl.Identifier()) {
			return
		}

		if !assert.Equal(t, "::CORBA", importDcl.ImportedScope().Identifier()) {
			return
		}
	})
}
