package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeDef(t *testing.T) {
	t.Run("simple typedef, primitive to 1 decl", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWint64, data: nil},
				{token: Identifier, data: "AA"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}

		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")
		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), lexer.GetSpecification()) {
			return
		}

		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		assert.Equal(t, "IdlInt64", typeDcl.TypeSpec().Identifier())
		assert.Nil(t, typeDcl.GetNextNode())

		decl := typeDcl.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "AA", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("simple typedef, primitive to 3 decls", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWint64, data: nil},
				{token: Identifier, data: "AA"},
				{token: ',', data: nil},
				{token: Identifier, data: "BB"},
				{token: ',', data: nil},
				{token: Identifier, data: "CC"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}

		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")
		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), lexer.GetSpecification()) {
			return
		}

		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		assert.Equal(t, "IdlInt64", typeDcl.TypeSpec().Identifier())
		assert.Nil(t, typeDcl.GetNextNode())

		decl := typeDcl.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "AA", decl.Identifier())
		assert.NotNil(t, decl.GetNextNode())

		decl = decl.GetNextNode()
		assert.NotNil(t, decl)
		assert.Equal(t, "BB", decl.Identifier())
		assert.NotNil(t, decl.GetNextNode())

		decl = decl.GetNextNode()
		assert.NotNil(t, decl)
		assert.Equal(t, "CC", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("struct typedef, scope to 1 new type ", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: Identifier, data: "SomeStruct"},
				{token: Identifier, data: "NewType"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")

		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), lexer.GetSpecification()) {
			return
		}

		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		assert.Equal(t, "SomeStruct", typeDcl.TypeSpec().Identifier())
		assert.Nil(t, typeDcl.GetNextNode())

		decl := typeDcl.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "NewType", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("struct typedef, scope to 3 new type ", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: Identifier, data: "SomeStruct"},
				{token: Identifier, data: "NewType"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")

		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), lexer.GetSpecification()) {
			return
		}

		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		assert.Equal(t, "SomeStruct", typeDcl.TypeSpec().Identifier())
		assert.Nil(t, typeDcl.GetNextNode())

		decl := typeDcl.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "NewType", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("struct typedef", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWint64, data: nil},
				{token: Identifier, data: "AA"},
				{token: ',', data: nil},
				{token: Identifier, data: "BB"},
				{token: ',', data: nil},
				{token: Identifier, data: "CC"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")
		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), lexer.GetSpecification()) {
			return
		}

		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		assert.Equal(t, "IdlInt64", typeDcl.TypeSpec().Identifier())
		assert.Nil(t, typeDcl.GetNextNode())

		decl := typeDcl.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "AA", decl.Identifier())
		assert.NotNil(t, decl.GetNextNode())

		decl = decl.GetNextNode()
		assert.NotNil(t, decl)
		assert.Equal(t, "BB", decl.Identifier())
		assert.NotNil(t, decl.GetNextNode())

		decl = decl.GetNextNode()
		assert.NotNil(t, decl)
		assert.Equal(t, "CC", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("typedef sequence<int64> AA;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWsequence, data: nil},
				{token: '<', data: nil},
				{token: RWint64, data: nil},
				{token: '>', data: nil},
				{token: Identifier, data: "AA"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		if !assert.NotNil(t, lexer.GetSpecification(), "No typedef impl") {
			return
		}
		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), lexer.GetSpecification()) {
			return
		}

		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		if !assert.Implements(t, (*ISequenceTypeDcl)(nil), typeDcl.TypeSpec()) {
			return
		}
		sequenceType, ok := typeDcl.TypeSpec().(ISequenceTypeDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, sequenceType) {
			return
		}
		assert.Equal(t, "IdlInt64", sequenceType.TypeSpec().Identifier())
		assert.Equal(t, int64(0), sequenceType.ArraySize())

		decl := typeDcl.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "AA", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())
	})

	t.Run("typedef sequence<int64> AA, BB, CC;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWsequence, data: nil},
				{token: '<', data: nil},
				{token: RWint64, data: nil},
				{token: '>', data: nil},
				{token: Identifier, data: "AA"},
				{token: ',', data: nil},
				{token: Identifier, data: "BB"},
				{token: ',', data: nil},
				{token: Identifier, data: "CC"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		if !assert.NotNil(t, lexer.GetSpecification(), "No typedef impl") {
			return
		}
		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), lexer.GetSpecification()) {
			return
		}

		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		if !assert.Implements(t, (*ISequenceTypeDcl)(nil), typeDcl.TypeSpec()) {
			return
		}
		sequenceType, ok := typeDcl.TypeSpec().(ISequenceTypeDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, sequenceType) {
			return
		}
		assert.Equal(t, "IdlInt64", sequenceType.TypeSpec().Identifier())
		assert.Equal(t, int64(0), sequenceType.ArraySize())

		decl := typeDcl.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "AA", decl.Identifier())
		assert.NotNil(t, decl.GetNextNode())

		decl = decl.GetNextNode()
		assert.NotNil(t, decl)
		assert.Equal(t, "BB", decl.Identifier())
		assert.NotNil(t, decl.GetNextNode())

		decl = decl.GetNextNode()
		assert.NotNil(t, decl)
		assert.Equal(t, "CC", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())

	})

	t.Run("typedef sequence<int64,10> AA;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWsequence, data: nil},
				{token: '<', data: nil},
				{token: RWint64, data: nil},
				{token: ',', data: nil},
				{token: Integer_literal, data: 10},
				{token: '>', data: nil},
				{token: Identifier, data: "AA"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")

		if !assert.Implements(t, (*ITypeDeclaratorDcl)(nil), lexer.GetSpecification()) {
			return
		}

		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		if !assert.Implements(t, (*ISequenceTypeDcl)(nil), typeDcl.TypeSpec()) {
			return
		}

		decl := typeDcl.Declarator()
		assert.NotNil(t, decl)
		assert.Equal(t, "AA", decl.Identifier())
		assert.Nil(t, decl.GetNextNode())

	})

	t.Run("typedef string<10> AA;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWstring, data: nil},
				{token: '<', data: nil},
				{token: Integer_literal, data: 10},
				{token: '>', data: nil},
				{token: Identifier, data: "AA"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")
		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		if !assert.Implements(t, (*IStringType)(nil), typeDcl.TypeSpec()) {
			return
		}
		stringType, ok := typeDcl.TypeSpec().(IStringType)
		assert.True(t, ok)
		if !assert.NotNil(t, stringType) {
			return
		}
		assert.Equal(t, int64(10), stringType.StringLength())

	})

	t.Run("typedef wstring<10> AA;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWwstring, data: nil},
				{token: '<', data: nil},
				{token: Integer_literal, data: 10},
				{token: '>', data: nil},
				{token: Identifier, data: "AA"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")
		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		assert.True(t, ok)
		if !assert.NotNil(t, typeDcl) {
			return
		}
		if !assert.Implements(t, (*IWideStringType)(nil), typeDcl.TypeSpec()) {
			return
		}
		stringType, ok := typeDcl.TypeSpec().(IWideStringType)
		assert.True(t, ok)
		if !assert.NotNil(t, stringType) {
			return
		}
		assert.Equal(t, int64(10), stringType.WideStringLength())
	})

	t.Run("typedef int64 BB, CC, DD;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWint64, data: nil},
				{token: Identifier, data: "BB"},
				{token: ',', data: nil},
				{token: Identifier, data: "CC"},
				{token: ',', data: nil},
				{token: Identifier, data: "DD"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")
		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		if !assert.True(t, ok) {
			return
		}
		if !assert.NotNil(t, typeDcl) {
			return
		}
		assert.Equal(t, "IdlInt64", typeDcl.TypeSpec().Identifier())
		decl := typeDcl.Declarator()
		decl = decl.GetNextNode()
		decl = decl.GetNextNode()
		decl = decl.GetNextNode()
		assert.Nil(t, decl)
	})

	t.Run("typedef struct AA {} int64 BB, CC, DD;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWstruct, data: nil},
				{token: Identifier, data: "AA"},
				{token: '{', data: nil},
				{token: '}', data: nil},
				{token: Identifier, data: "BB"},
				{token: ',', data: nil},
				{token: Identifier, data: "CC"},
				{token: ',', data: nil},
				{token: Identifier, data: "DD"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")
		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		if !assert.True(t, ok) {
			return
		}
		if !assert.NotNil(t, typeDcl) {
			return
		}

		if !assert.Implements(t, (*IStructDcl)(nil), typeDcl.TypeSpec()) {
			return
		}

		decl := typeDcl.Declarator()
		assert.Equal(t, "BB", decl.Identifier())
		decl = decl.GetNextNode()
		assert.Equal(t, "CC", decl.Identifier())
		decl = decl.GetNextNode()
		assert.Equal(t, "DD", decl.Identifier())
		decl = decl.GetNextNode()
		assert.Nil(t, decl)
	})

	t.Run("typedef sequence<sequence<AA>> int64 BB, CC, DD;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWtypedef, data: nil},
				{token: RWsequence, data: nil},
				{token: '<', data: nil},
				{token: RWsequence, data: nil},
				{token: '<', data: nil},
				{token: Identifier, data: "AA"},
				{token: '>', data: nil},
				{token: '>', data: nil},
				{token: Identifier, data: "BB"},
				{token: ',', data: nil},
				{token: Identifier, data: "CC"},
				{token: ',', data: nil},
				{token: Identifier, data: "DD"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification(), "No typedef impl")
		typeDcl, ok := lexer.GetSpecification().(ITypeDeclaratorDcl)
		if !assert.True(t, ok) {
			return
		}
		if !assert.NotNil(t, typeDcl) {
			return
		}

		if !assert.Implements(t, (*ISequenceTypeDcl)(nil), typeDcl.TypeSpec()) {
			return
		}

		decl := typeDcl.Declarator()
		assert.Equal(t, "BB", decl.Identifier())
		decl = decl.GetNextNode()
		assert.Equal(t, "CC", decl.Identifier())
		decl = decl.GetNextNode()
		assert.Equal(t, "DD", decl.Identifier())
		decl = decl.GetNextNode()
		assert.Nil(t, decl)
	})

}
