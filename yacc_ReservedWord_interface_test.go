package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterface(t *testing.T) {
	t.Run("forward interface decl", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}
		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, interfaceDcl.Body())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.True(t, interfaceDcl.Forward())
	})

	t.Run("interface decl", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}
		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, interfaceDcl.Body())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
	})

	t.Run("interface decl with declared struct", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}
		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}

		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
		assert.NotNil(t, interfaceDcl.Body())

		body := interfaceDcl.Body()
		body01 := body
		if !assert.Implements(t, (*IStructDcl)(nil), body01) {
			return
		}

		body02 := body01.GetNextNode()
		assert.Nil(t, body02)
	})

	t.Run("interface decl with one inheritance", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}
		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, interfaceDcl.Body())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
		assert.NotNil(t, interfaceDcl.InterfaceHeader().Inheritance())
		inheritance := interfaceDcl.InterfaceHeader().Inheritance()
		assert.Equal(t, "ABC", inheritance.Identifier())
		assert.Nil(t, inheritance.GetNextNode())
	})

	t.Run("interface decl with two inheritance", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}

		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.Nil(t, interfaceDcl.Body())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
		assert.NotNil(t, interfaceDcl.InterfaceHeader().Inheritance())

		inheritance := interfaceDcl.InterfaceHeader().Inheritance()
		assert.Equal(t, "ABC", inheritance.Identifier())

		inheritance02 := inheritance.GetNextNode()
		assert.Equal(t, "DEF", inheritance02.Identifier())
	})

	t.Run("interface decl with one operation", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}

		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
		assert.Nil(t, interfaceDcl.InterfaceHeader().Inheritance())
		assert.NotNil(t, interfaceDcl.Body())

		operation := interfaceDcl.Body()
		if !assert.Implements(t, (*IOperationDcl)(nil), operation) {
			return
		}
		operation01 := operation.(IOperationDcl)
		assert.Equal(t, "ABCDEF", operation01.Identifier())
		assert.Equal(t, "IdlVoidType", operation01.ReturnType().Identifier())
		assert.Nil(t, operation01.GetNextNode())
	})

	t.Run("interface decl with three operations, no params", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}

		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
		assert.Nil(t, interfaceDcl.InterfaceHeader().Inheritance())
		assert.NotNil(t, interfaceDcl.Body())

		operation := interfaceDcl.Body()
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
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}

		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
		assert.Nil(t, interfaceDcl.InterfaceHeader().Inheritance())
		assert.NotNil(t, interfaceDcl.Body())

		// first operation
		operation := interfaceDcl.Body()
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

	t.Run("", func(t *testing.T) {
		assert.Equal(t, ParamDirection(1), ParamDirectionIn)
		assert.Equal(t, ParamDirection(2), ParamDirectionOut)
		assert.Equal(t, ParamDirection(3), ParamDirectionOut|ParamDirectionIn)
	})

	t.Run("interface decl with one operation, and one param", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}

		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
		assert.Nil(t, interfaceDcl.InterfaceHeader().Inheritance())
		assert.NotNil(t, interfaceDcl.Body())

		operation := interfaceDcl.Body()
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
				{token: RWinterface, data: nil},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}

		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Abstract())
		assert.False(t, interfaceDcl.InterfaceHeader().InterfaceKind().Local())
		assert.False(t, interfaceDcl.Forward())
		assert.Nil(t, interfaceDcl.InterfaceHeader().Inheritance())
		assert.NotNil(t, interfaceDcl.Body())

		operation := interfaceDcl.Body()
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

	t.Run("interface with readonly attr", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWreadonly, data: nil},
				{token: RWattribute, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "a"},
				{token: ',', data: nil},
				{token: Identifier, data: "b"},
				{token: ',', data: nil},
				{token: Identifier, data: "c"},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}
		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}

		body := interfaceDcl.Body()
		if !assert.Implements(t, (*IAttributeDcl)(nil), body) {
			return
		}
		attrDcl := body.(IAttributeDcl)
		assert.Equal(t, "IdlLong", attrDcl.TypeSpec().Identifier())

		node := attrDcl.AttrDeclarator().SimpleDeclarator()
		assert.NotNil(t, node)
		assert.Equal(t, "a", node.Identifier())

		node = node.GetNextNode()
		assert.NotNil(t, node)
		assert.Equal(t, "b", node.Identifier())

		node = node.GetNextNode()
		assert.NotNil(t, node)
		assert.Equal(t, "c", node.Identifier())

		node = node.GetNextNode()
		assert.Nil(t, node)

	})

	t.Run("interface with attr", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RWinterface, data: nil},
				{token: Identifier, data: "ABCDEF"},
				{token: '{', data: nil},
				{token: RWattribute, data: nil},
				{token: RWlong, data: nil},
				{token: Identifier, data: "a"},
				{token: ',', data: nil},
				{token: Identifier, data: "b"},
				{token: ',', data: nil},
				{token: Identifier, data: "c"},
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
		if !assert.Implements(t, (*IInterfaceDcl)(nil), lexer.GetSpecification()) {
			return
		}
		interfaceDcl, ok1 := lexer.GetSpecification().(IInterfaceDcl)
		if !assert.True(t, ok1) {
			return
		}

		body := interfaceDcl.Body()
		if !assert.Implements(t, (*IAttributeDcl)(nil), body) {
			return
		}
		attrDcl := body.(IAttributeDcl)
		assert.Equal(t, "IdlLong", attrDcl.TypeSpec().Identifier())

		node := attrDcl.AttrDeclarator().SimpleDeclarator()
		assert.NotNil(t, node)
		assert.Equal(t, "a", node.Identifier())

		node = node.GetNextNode()
		assert.NotNil(t, node)
		assert.Equal(t, "b", node.Identifier())

		node = node.GetNextNode()
		assert.NotNil(t, node)
		assert.Equal(t, "c", node.Identifier())

		node = node.GetNextNode()
		assert.Nil(t, node)

	})

}
