package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueType(t *testing.T) {
	t.Run("valuetype ABCDEF;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}
		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, valueDcl.Body())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.True(t, valueDcl.Forward())
	})

	t.Run("abstract valuetype CustomMarshal;", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWabstract, data: nil},
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "CustomMarshal"},
				{token: ';', data: nil},
			})

		parser := YaccIdlNewParser()

		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "CustomMarshal", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}
		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, valueDcl.Body())
		assert.True(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.True(t, valueDcl.Forward())
	})

	t.Run("valuetype ABCDEF{};", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
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
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}
		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, valueDcl.Body())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
	})

	t.Run("valuetype ABCDEF{struct GHI{};};", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABC"},
				{token: '{', data: nil},
				{token: RWstruct, data: nil},
				{token: Identifier, data: "GHI"},
				{token: '{', data: nil},
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
		assert.Equal(t, "ABC", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}
		def := lexer.GetSpecification()
		valueDcl, ok1 := def.(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
		assert.NotNil(t, valueDcl.Body())

		body := valueDcl.Body()
		body01 := body
		if !assert.Implements(t, (*IStructDcl)(nil), body01) {
			return
		}
		structDcl, ok := body01.(IStructDcl)
		if !assert.True(t, ok) {
			return
		}
		assert.Equal(t, "GHI", structDcl.Identifier())
		body02 := body01.GetNextNode()
		assert.Nil(t, body02)

		def = def.GetNextNode()
		assert.Nil(t, def)
	})

	t.Run("valuetype ABCDEF:ABC{};", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: ':', data: nil},
				{token: Identifier, data: "ABC"},
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
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}

		def := lexer.GetSpecification()
		valueDcl, ok1 := def.(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, valueDcl.Body())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
		assert.NotNil(t, valueDcl.ValueHeader().ValueInheritanceSpec())

		ValueInheritanceSpec := valueDcl.ValueHeader().ValueInheritanceSpec()
		valueName := ValueInheritanceSpec.ValueName()
		assert.Equal(t, "ABC", valueName.Identifier())
		assert.Nil(t, ValueInheritanceSpec.GetNextNode())

		def = def.GetNextNode()
		assert.Nil(t, def)
	})

	t.Run("valuetype ABCDEF:ABC,DEF{};", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: ':', data: nil},
				{token: Identifier, data: "ABC"},
				{token: ',', data: nil},
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
		assert.NotNil(t, lexer.GetSpecification())
		assert.Equal(t, "ABCDEF", lexer.GetSpecification().Identifier())
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}

		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, valueDcl.Body())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
		assert.NotNil(t, valueDcl.ValueHeader().ValueInheritanceSpec())

		ValueInheritanceSpec := valueDcl.ValueHeader().ValueInheritanceSpec()
		valueName := ValueInheritanceSpec.ValueName()
		assert.Equal(t, "ABC", valueName.Identifier())

		valueName = valueName.GetNextNode()
		assert.Equal(t, "DEF", valueName.Identifier())
	})

	t.Run("valuetype ABCDEF{void ABCDEF();};", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},

				{token: RWvoid, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '(', data: nil},
				{token: ')', data: nil},
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
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}

		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
		assert.Nil(t, valueDcl.ValueHeader().ValueInheritanceSpec())
		assert.NotNil(t, valueDcl.Body())

		operation := valueDcl.Body()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}
		operation01 := operation.(IOperationDcl)
		assert.Equal(t, "ABCDEF", operation01.Identifier())
		assert.Equal(t, "IdlVoidType", operation01.ReturnType().Identifier())
		assert.Nil(t, operation01.GetNextNode())
	})

	t.Run("valuetype ABCDEF{void A();void B();R C();};", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},

				{token: RWvoid, data: nil},
				{token: Identifier, data: "A"},
				{token: '(', data: nil},
				{token: ')', data: nil},
				{token: ';', data: nil},

				{token: RWint64, data: nil},
				{token: Identifier, data: "B"},
				{token: '(', data: nil},
				{token: ')', data: nil},
				{token: ';', data: nil},

				{token: Identifier, data: "R"},
				{token: Identifier, data: "C"},
				{token: '(', data: nil},
				{token: ')', data: nil},
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
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}

		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
		assert.Nil(t, valueDcl.ValueHeader().ValueInheritanceSpec())
		assert.NotNil(t, valueDcl.Body())

		operation := valueDcl.Body()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}
		operation01 := operation.(IOperationDcl)
		assert.Equal(t, "A", operation01.Identifier())
		assert.Equal(t, "IdlVoidType", operation01.ReturnType().Identifier())
		assert.NotNil(t, operation01.GetNextNode())

		operation = operation.GetNextNode()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}

		operation02 := operation.(IOperationDcl)
		assert.Equal(t, "B", operation02.Identifier())
		assert.Equal(t, "IdlInt64", operation02.ReturnType().Identifier())
		assert.NotNil(t, operation02.GetNextNode())

		operation = operation.GetNextNode()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}

		operation03 := operation.(IOperationDcl)
		assert.Equal(t, "C", operation03.Identifier())

		assert.Equal(t, "R", operation03.ReturnType().Identifier())
		assert.Nil(t, operation03.GetNextNode())
	})

	t.Run("interface decl with three operations, no params, with exceptions", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				// first operation with one excpetions
				{token: RWvoid, data: nil},
				{token: Identifier, data: "A"},
				{token: '(', data: nil},
				{token: ')', data: nil},
				{token: RWraises, data: nil},
				{token: '(', data: nil},
				{token: Identifier, data: "Ex01"},
				{token: ')', data: nil},
				{token: ';', data: nil},
				// second operation with three exceptions
				{token: RWint64, data: nil},
				{token: Identifier, data: "B"},
				{token: '(', data: nil},
				{token: ')', data: nil},
				{token: RWraises, data: nil},
				{token: '(', data: nil},
				{token: Identifier, data: "Ex01"},
				{token: ',', data: nil},
				{token: Identifier, data: "Ex02"},
				{token: ',', data: nil},
				{token: Identifier, data: "Ex03"},
				{token: ')', data: nil},
				{token: ';', data: nil},
				// third operation with no exceptions
				{token: Identifier, data: "R"},
				{token: Identifier, data: "C"},
				{token: '(', data: nil},
				{token: ')', data: nil},
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
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}

		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
		assert.Nil(t, valueDcl.ValueHeader().ValueInheritanceSpec())
		assert.NotNil(t, valueDcl.Body())

		// first operation
		operation := valueDcl.Body()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}

		operation01 := operation.(IOperationDcl)
		assert.Equal(t, "A", operation01.Identifier())
		assert.Equal(t, "IdlVoidType", operation01.ReturnType().Identifier())
		assert.NotNil(t, operation01.GetNextNode())
		assert.NotNil(t, operation01.ExceptionDcl())
		exception01 := operation01.ExceptionDcl()
		assert.Equal(t, "Ex01", exception01.Identifier())
		exception02 := exception01.GetNextNode()
		assert.Nil(t, exception02)

		// second operation
		operation = operation.GetNextNode()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}

		operation02 := operation.(IOperationDcl)
		assert.Equal(t, "B", operation02.Identifier())
		assert.Equal(t, "IdlInt64", operation02.ReturnType().Identifier())
		assert.NotNil(t, operation02.GetNextNode())

		operation = operation.GetNextNode()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}

		assert.NotNil(t, operation02.ExceptionDcl())
		exception01 = operation02.ExceptionDcl()
		assert.Equal(t, "Ex01", exception01.Identifier())
		exception02 = exception01.GetNextNode()
		assert.Equal(t, "Ex02", exception02.Identifier())
		exception03 := exception02.GetNextNode()
		assert.Equal(t, "Ex03", exception03.Identifier())
		assert.Nil(t, exception03.GetNextNode())

		// third operation
		operation03 := operation.(IOperationDcl)
		assert.Equal(t, "C", operation03.Identifier())

		assert.Equal(t, "R", operation03.ReturnType().Identifier())
		assert.Nil(t, operation03.GetNextNode())
	})

	t.Run("interface decl with one operation, and one param", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},

				{token: RWvoid, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '(', data: nil},
				{token: RWin, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "a"},

				{token: ')', data: nil},
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
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}

		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
		assert.Nil(t, valueDcl.ValueHeader().ValueInheritanceSpec())
		assert.NotNil(t, valueDcl.Body())

		operation := valueDcl.Body()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}
		operation01 := operation.(IOperationDcl)
		assert.Equal(t, "ABCDEF", operation01.Identifier())
		assert.Equal(t, "IdlVoidType", operation01.ReturnType().Identifier())
		assert.Nil(t, operation01.GetNextNode())
		assert.NotNil(t, operation01.ParamDcl())

		param := operation01.ParamDcl()
		if !assert.Implements(t, (*IOperationalParameter)(nil), param) {
			return
		}
		param01, ok := param.(IOperationalParameter)
		assert.True(t, ok)
		assert.Equal(t, "a", param01.Identifier())
		assert.Equal(t, "IdlLong", param01.ParamType().Identifier())
		assert.Nil(t, param01.GetNextNode())
	})

	t.Run("interface decl with one operation, and three param", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWvaluetype, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},

				{token: RWvoid, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '(', data: nil},
				{token: RWin, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "a"},
				{token: ',', data: nil},
				{token: RWout, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "b"},
				{token: ',', data: nil},
				{token: RWinout, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "c"},
				{token: ')', data: nil},
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
		if !assert.Implements(t, (*IValueDefDcl)(nil), lexer.GetSpecification()) {
			return
		}

		valueDcl, ok1 := lexer.GetSpecification().(IValueDefDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, valueDcl.ValueHeader().ValueKind().Abstract())
		assert.False(t, valueDcl.ValueHeader().ValueKind().Local())
		assert.False(t, valueDcl.Forward())
		assert.Nil(t, valueDcl.ValueHeader().ValueInheritanceSpec())
		assert.NotNil(t, valueDcl.Body())

		operation := valueDcl.Body()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}
		operation01 := operation.(IOperationDcl)
		assert.Equal(t, "ABCDEF", operation01.Identifier())
		assert.Equal(t, "IdlVoidType", operation01.ReturnType().Identifier())
		assert.Nil(t, operation01.GetNextNode())
		assert.NotNil(t, operation01.ParamDcl())

		param := operation01.ParamDcl()
		if !assert.Implements(t, (*IOperationalParameter)(nil), param) {
			return
		}
		param01, ok := param.(IOperationalParameter)
		assert.True(t, ok)
		assert.Equal(t, "a", param01.Identifier())
		assert.Equal(t, "IdlLong", param01.ParamType().Identifier())
		assert.Equal(t, ParamDirectionIn, param01.Direction())
		assert.NotNil(t, param01.GetNextNode())

		param = param.GetNextNode()
		param02, ok := param.(IOperationalParameter)
		assert.True(t, ok)
		assert.Equal(t, "b", param02.Identifier())
		assert.Equal(t, "IdlLong", param02.ParamType().Identifier())
		assert.Equal(t, ParamDirectionOut, param02.Direction())
		assert.NotNil(t, param02.GetNextNode())

		param = param.GetNextNode()
		param03, ok := param.(IOperationalParameter)
		assert.True(t, ok)
		assert.Equal(t, "c", param03.Identifier())
		assert.Equal(t, "IdlLong", param03.ParamType().Identifier())
		assert.Equal(t, ParamDirectionOut|ParamDirectionIn, param03.Direction())
		assert.Nil(t, param03.GetNextNode())
	})
}
