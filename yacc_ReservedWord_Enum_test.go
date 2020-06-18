package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnums(t *testing.T) {
	t.Run("enum with one value", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWenum, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: Identifier, data: "AA"},
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
		if !assert.Implements(t, (*IEnumDcl)(nil), lexer.GetSpecification()) {
			return
		}
		enumDcl, ok := lexer.GetSpecification().(IEnumDcl)
		assert.True(t, ok)
		assert.NotNil(t, enumDcl.Enumerator())
		enum01 := enumDcl.Enumerator()
		assert.NotNil(t, enum01)
		assert.Equal(t, "AA", enum01.Identifier())
		assert.Nil(t, enum01.GetNextNode())
	})

	t.Run("enum with three value", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWenum, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: Identifier, data: "AA"},
				{token: ',', data: nil},
				{token: Identifier, data: "BB"},
				{token: ',', data: nil},
				{token: Identifier, data: "CC"},
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
		if !assert.Implements(t, (*IEnumDcl)(nil), lexer.GetSpecification()) {
			return
		}
		enumDcl, ok := lexer.GetSpecification().(IEnumDcl)
		assert.True(t, ok)
		assert.NotNil(t, enumDcl.Enumerator())

		enum01 := enumDcl.Enumerator()
		if !assert.NotNil(t, enum01) {
			return
		}
		assert.Equal(t, "AA", enum01.Identifier())
		assert.NotNil(t, enum01.GetNextNode())

		enum02 := enum01.GetNextNode()
		if !assert.NotNil(t, enum02) {
			return
		}
		assert.Equal(t, "BB", enum02.Identifier())
		assert.NotNil(t, enum02.GetNextNode())

		enum03 := enum02.GetNextNode()
		if !assert.NotNil(t, enum03) {
			return
		}
		assert.Equal(t, "CC", enum03.Identifier())
		assert.Nil(t, enum03.GetNextNode())
	})
}
