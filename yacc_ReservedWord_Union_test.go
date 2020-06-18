package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnion(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWunion, data: nil},
				{token: Identifier, data: "TargetAddress"},
				{token: RWswitch, data: nil},
				{token: '(', data: nil},
				{token: RWshort, data: nil},
				{token: ')', data: nil},
				{token: '{', data: nil},
				{token: RWcase, data: nil},
				{token: Identifier, data: "KeyAddr"},
				{token: ':', data: nil},
				{token: Identifier, data: "ABC"},
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
		unionDcl, ok := lexer.GetSpecification().(IUnionDcl)
		assert.True(t, ok)
		assert.Equal(t, "TargetAddress", unionDcl.Identifier())
		assert.False(t, unionDcl.Forward())
		assert.Equal(t, "IdlShort", unionDcl.SwitchType().Identifier())

		assert.NotNil(t, unionDcl.UnionBody())
		//unionBody := unionDcl.UnionBody()
		//unionBody.CaseLabels().
	})

	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWunion, data: nil},
				{token: Identifier, data: "TargetAddress"},
				{token: ';', data: nil},
			})
		parser := YaccIdlNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		unionDcl, ok := lexer.GetSpecification().(IUnionDcl)
		assert.True(t, ok)
		assert.Equal(t, "TargetAddress", unionDcl.Identifier())
		assert.True(t, unionDcl.Forward())
		assert.Nil(t, unionDcl.UnionBody())
	})
}
