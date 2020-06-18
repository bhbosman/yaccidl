%{
package yaccidl
%}

%token None
%token SingleComment
%token MultiLineComment
%token WhiteSpace
%token HashLoadDefinition
%token HashPragma
%token HashDefine
%token HashUnDefine
%token Hashifdef
%token Hashifndef
%token Hashendif
%token Hashelse
%token HashInclude
%token ScopeOp // "::"
%token ShlOp // "<<"
%token ShrOp // ">>"
%token Annotation //"@annotation"
%token String_literal
%token Identifier
%token Integer_literal
%token Floating_pt_literal
%token Fixed_pt_literal
%token Character_literal
%token Wide_Character_literal
%token Wide_String_literal
%token Hex_literal
%token ErrorFileNotFound
%token RWabstract
%token RWany
%token RWalias
%token RWattribute
%token RWbitfield
%token RWbitmask
%token RWbitset
%token RWboolean
%token RWcase
%token RWchar
%token RWcomponent
%token RWconnector
%token RWconst
%token RWconsumes
%token RWcontext
%token RWcustom
%token RWdefault
%token RWdouble
%token RWexception
%token RWemits
%token RWenum
%token RWeventtype
%token RWfactory
%token RWFALSE
%token RWfinder
%token RWfixed
%token RWfloat
%token RWgetraises
%token RWhome
%token RWimport
%token RWin
%token RWinout
%token RWinterface
%token RWlocal
%token RWlong
%token RWmanages
%token RWmap
%token RWmirrorport
%token RWmodule
%token RWmultiple
%token RWnative
%token RWObject
%token RWoctet
%token RWoneway
%token RWout
%token RWprimarykey
%token RWprivate
%token RWport
%token RWporttype
%token RWprovides
%token RWpublic
%token RWpublishes
%token RWraises
%token RWreadonly
%token RWsetraises
%token RWsequence
%token RWshort
%token RWstring
%token RWstruct
%token RWsupports
%token RWswitch
%token RWTRUE
%token RWtruncatable
%token RWtypedef
%token RWtypeid
%token RWtypename
%token RWtypeprefix
%token RWunsigned
%token RWunion
%token RWuses
%token RWValueBase
%token RWvaluetype
%token RWvoid
%token RWwchar
%token RWwstring
%token RWint8
%token RWuint8
%token RWint16
%token RWint32
%token RWint64
%token RWuint16
%token RWuint32
%token RWuint64

%union{
	Node IYaccNode
	Value IValue
	CharValue rune
	BoolValue bool
	InterfaceKind IInterfaceKind
	InterfaceHeader IInterfaceHeader
	ParamDirection ParamDirection
	ValueInheritanceSpec IValueInheritanceSpec
	ValueHeader IValueHeader
	AttrDeclarator IAttrDeclarator
	ElementSpec IElementSpec
	ConstValue IBaseConstantValue
	UnionBody IUnionBody
}


%type <Value> fixed_array_size fixed_array_sizePlus
%type <AttrDeclarator> readonly_attr_declarator attr_declarator
%type <ValueHeader> value_header value_forward_dcl
%type <ParamDirection> param_attribute
%type <InterfaceHeader> interface_header interface_forward_dcl
%type <InterfaceKind> interface_kind value_kind

%type <BoolValue> boolean_literal
%type <Value> primary_expr literal
%type <CharValue> Character_literal Wide_Character_literal
%type <Value> Floating_pt_literal Fixed_pt_literal


%type <Value>  positive_int_const const_expr or_expr xor_expr and_expr shift_expr add_expr mult_expr unary_expr unary_operator Integer_literal

%type <Node> definitionPlus interface_namePlus interface_name exportStar String_literal Wide_String_literal component_forward_dcl component_def
%type <Node> specification Identifier simple_declarator
%type <Node> definition
%type <Node> module_dcl value_namePlus value_name
%type <Node> const_dcl simple_declaratorPlus
%type <Node> type_dcl switch_type_spec
%type <Node> except_dcl attr_raises_expr
%type <ConstValue> case_label case_labelPlus
%type <Node> interface_dcl
%type <Node> value_dcl enumerators

%type <Node> interface_inheritance_spec
%type <Node> type_id_dcl struct_dcl constr_type_dcl
%type <Node> type_prefix_dcl import_dcl component_dcl home_dcl event_dcl porttype_dcl connector_dcl
%type <Node> template_module_dcl template_module_inst annotation_dcl   scoped_name scoped_namePlus


%type <Node> const_type integer_type signed_int unsigned_int signed_short_int signed_long_int signed_longlong_int
%type <Node> signed_tiny_int unsigned_short_int unsigned_long_int unsigned_longlong_int unsigned_tiny_int char_type wide_char_type
%type <Node> boolean_type octet_type floating_pt_type fixed_pt_const_type string_type wide_string_type memberPlus
%type <Node> struct_def interface_def exportPlus interface_body op_type_spec raises_expr any_declaratorsPlus
%type <Node> memberStar union_dcl enum_dcl
%type <Node> member type_spec typedef_dcl base_type_spec readonly_attr_spec
%type <Node> declarators value_def value_elementPlus value_element state_member init_dcl
%type <Node> simple_type_spec native_dcl union_def union_forward_dcl
%type <Node> template_type_spec bitset_dcl bitmask_dcl
%type <Node> struct_forward_dcl sequence_type param_dclPlus event_header attr_spec
%type <Node>  any_declarator declaratorPlus declarator array_declarator any_type object_type value_base_type enumerator




%type <Node> RWmodule RWcustom RWsupports op_dcl export attr_dcl op_oneway_dcl op_with_context parameter_dcls
%type <Node>  RWfixed RWvoid init_param_dcls init_param_dclPlus init_param_dcl imported_scope in_parameter_dclsStar
%type <Node> RWstruct in_parameter_dcls in_param_dclPlus in_param_dcl String_literalPlus context_expr
%type <Node> RWfloat fixed_pt_type map_type
%type <Node> RWdouble value_box_def component_header component_inheritance_spec component_body
%type <Node> RWlong
%type <Node> RWshort
%type <Node> RWint16   factory_param_dcl supported_interface_spec component_exportStar component_export provides_dcl interface_type
%type <Node> RWint32 uses_dcl home_header home_inheritance_spec home_body home_exportStar home_export factory_dcl factory_param_dcls
%type <Node> RWint64 factory_param_dclPlus
%type <Node> RWunsigned value_elementStar
%type <Node> RWuint16
%type <Node> RWuint32 emits_dcl publishes_dcl consumes_dcl primary_key_spec finder_dcl event_forward_dcl event_abs_def
%type <Node> RWuint64 event_def
%type <Node> RWstring
%type <Node> RWwstring porttype_forward_dcl porttype_def port_ref port_exportStar port_export port_dcl
%type <Node> RWchar
%type <Node> RWwchar RWcase RWdefault

%type <Node> RWboolean
%type <Node> RWoctet RWnative
%type <Node> RWany connector_header connector_inherit_spec connector_export formal_parameters
%type <Node> RWObject formal_parameterPlus formal_parameter formal_parameter_type tpl_definitionPlus tpl_definition actual_parameters
%type <Node> RWValueBase actual_parameterPlus actual_parameter IdentifierStar bitfieldStar
%type <Node> RWint8
%type <Node> RWuint8 annotation_bodyEntryStar annotation_body annotation_member annotation_member_type any_const_type annotation_appl
%type <Node> RWinterface annotation_appl_params
%type <Node> RWlocal
%type <Node> RWabstract
%type <Node> RWenum bitfield bitfield_spec destination_type bit_valueStar bit_value annotation_header annotation_bodyEntry
%type <Node> RWin
%type <Node> RWinout
%type <Node> RWout
%type <Node> RWconst ScopeOp
%type <Node> RWexception
%type <Node> RWtypedef
%type <Node> RWvaluetype
%type <Node> RWtypeprefix RWimport
%type <Node> RWreadonly
%type <Node> RWattribute
%type <Node> RWtruncatable
%type <Node> RWsequence param_dcl type_declarator
%type <ValueInheritanceSpec> value_inheritance_spec
%type <UnionBody>   case casePlus switch_body

%type <ElementSpec> element_spec

%start specification
%%

//(1)
specification:
	definitionPlus
	{	
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("specification/definitionPlus", $1)
		}

		if assignSpecification, ok := YaccIdllex.(IAssignSpecification); ok {
			assignSpecification.AssignSpecification($1)
			goto exitBlock_specification_definitionPlus
		}
		YaccIdllex.Error("IAssignSpecification not implemented")
		return 1

		exitBlock_specification_definitionPlus:
	}

//(2)(71)//(98)//(111)//(133)//(144)//(153)//(171)//(184)//(218)
definitionPlus:
	{
//??
	}
	|definition
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definitionPlus/definition", $1)
                }
		$$ = $1
	}
	| definitionPlus definition {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                		yaccLogger.Log("definitionPlus/definitionPlus definition", $1, $2)
               	}
		if $1 != nil && $2 != nil {
			node := $1
			for node.GetNextNode() != nil {
				node = node.GetNextNode()
			}
			_ = node.SetNextNode($2)
			$$ = $1
		} else if $1 == nil && $2 != nil {
			$$ = $2
                } else if $1 != nil && $2 == nil {
                	$$ = $1
                }
        }

definition:
	module_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
	yaccLogger.Log("definition/module_dcl", $1)
                			}
		$$ = $1
	}
	| const_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/const_dcl", $1)
                }
		$$ = $1
	}
	| type_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/type_dcl", $1)
                }
		$$ = $1
	}
	| except_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/except_dcl", $1)
}
		$$ = $1
	}
	| interface_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/interface_dcl", $1)
}
		$$ = $1
	}
	| value_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/value_dcl", $1)
}
		$$ = $1
	}
	| type_id_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/type_id_dcl", $1)
		}
		$$ = $1
	}
	| type_prefix_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/type_prefix_dcl", $1)
}
		$$ = $1
	}
	| import_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/import_dcl", $1)
}
		$$ = $1
	}
	| component_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/component_dcl", $1)
}
		$$ = $1
	}
	| home_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/home_dcl", $1)
}
		$$ = $1
	}
	| event_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/event_dcl", $1)
}
		$$ = $1
	}
	| porttype_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/porttype_dcl", $1)
}
		$$ = $1
	}
	| connector_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/connector_dcl", $1)
}
		$$ = $1
	}
	| template_module_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/template_module_dcl", $1)
                }
		$$ = $1
	}
	| template_module_inst ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/template_module_inst", $1)
}
		$$ = $1

	}
	| annotation_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("definition/annotation_dcl", $1)
}
		$$ = $1
	}

//(3)
module_dcl:
	RWmodule Identifier '{' definitionPlus '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("RWmodule Identifier '{' definitionPlus '}'", $2, $4)
}
		var err error
		$$, err = newModuleDcl($1.LexemData(), $2.Identifier(), $4)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| RWmodule Identifier '{'  '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("RWmodule Identifier '{'  '}'", $2)
}
		var err error
		$$, err = newModuleDcl($1.LexemData(), $2.Identifier(), nil)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(4)
scoped_namePlus:
	scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("scoped_namePlus/scoped_name", $1)
}
		$$ = $1
	}
	| scoped_namePlus ',' scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("scoped_namePlus/scoped_namePlus ',' scoped_name", $1, $3)
}
		node := $1
		for node.GetNextNode() != nil{
			node = node.GetNextNode()
		}
		_ = node.SetNextNode($3)
	}
scoped_name:
	Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("scoped_name/Identifier", $1)
}
		$$, _ = newScopeName($1.LexemData(), $1.Identifier())
	}
	| ScopeOp Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("scoped_name/ScopeOp Identifier", $2)
}
		$$, _ = newScopeName($1.LexemData(), "::" + $2.Identifier())
	}
	| scoped_name ScopeOp Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("scoped_name/scoped_name ScopeOp Identifier", $1, $3)
}
		$$, _ = newScopeName($1.LexemData(), $1.Identifier() +  "::" + $3.Identifier())
	}
//(5)
const_dcl:
	RWconst const_type Identifier '=' const_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("RWconst const_type Identifier '=' const_expr", $1, $2, $3 , $5)
}

		var err error
		$$, err = NewConstDcl($1.LexemData(), $2, $3.Identifier(), $5)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(6)
const_type:
	integer_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/integer_type", $1)
}
		$$ = $1
	}
	| floating_pt_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/floating_pt_type", $1)
}
		$$ = $1
	}
	| fixed_pt_const_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/fixed_pt_const_type", $1)
}
		$$ = $1
	}
	| char_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/char_type", $1)
}
		$$ = $1
	}
	| wide_char_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/wide_char_type", $1)
}
		$$ = $1
	}
	| boolean_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/boolean_type", $1)
}
		$$ = $1
	}
	| octet_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/octet_type", $1)
}
		$$ = $1
	}
	| string_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/string_type", $1)
}
		$$ = $1
	}
	| wide_string_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/wide_string_type", $1)
}
		$$ = $1
	}
	| scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_type/wide_string_type", $1)
}
//		$$ = NewYaccNode($1.Identifier(),$1.LexemData())
		$$ = $1
	}
//(7)
const_expr:
	or_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("const_expr/or_expr")
}
		$$ = $1
	}
//(8)
or_expr:
	xor_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("or_expr/xor_expr", $1)
}
		$$ = $1
	}
	| or_expr '|' xor_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("or_expr/or_expr '|' xor_expr", $1, $3)
}
		v1, ok1 := $1.(IInt64Value)
		v3, ok3 := $3.(IInt64Value)
		if ok1 && ok3 {
			var err error
			$$, err = NewInt64Value(v1.Int64Value() | v3.Int64Value())
			if err != nil {
				YaccIdllex.Error(err.Error())
				return 1
			}
			goto exitBlock_or_expr_or_xor_expr
		}
		YaccIdllex.Error("type not supported (| operator)")
		return 1
		exitBlock_or_expr_or_xor_expr:
	}
//(9)
xor_expr:
	and_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("xor_expr/and_expr", $1)
}
		$$ = $1
	}
	| xor_expr '^' and_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("xor_expr/xor_expr '^' and_expr", $1, $3)
}
		YaccIdllex.Error("Impelemnt me!! '^'")
		return 1
//		$$ = $1 ^ $3
	}

//(10)
and_expr:
	shift_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("and_expr/shift_expr", $1)
}
		$$ = $1
	}
	| and_expr '&' shift_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("and_expr/and_expr '&' shift_expr", $1, $3)
}
		v1, ok1 := $1.(IInt64Value)
		v3, ok3 := $3.(IInt64Value)
		if ok1 && ok3 {
			var err error
			$$, err = NewInt64Value(v1.Int64Value() & v3.Int64Value())
			if err != nil {
				YaccIdllex.Error(err.Error())
				return 1
			}
		} else {
				YaccIdllex.Error("type not supported (& operator)")
				return 1
		}
	}

//(11)
shift_expr:
	add_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("shift_expr/add_expr", $1)
}
		var err error
		$$ = $1
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| shift_expr ShrOp add_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("shift_expr/shift_expr ShrOp add_expr", $1, $3)
}
		v1, ok1 := $1.(IInt64Value)
		v3, ok3 := $3.(IInt64Value)
		if ok1 && ok3 {
			if v3.Int64Value() < 0 {
				YaccIdllex.Error("shift count must be positive")
			}
			$$, _ = NewInt64Value(v1.Int64Value() >> v3.Int64Value())
			//
			goto exitBlock_shift_expr_ShrOp_add_expr
		}
		YaccIdllex.Error("type not supported(>> operator)")
		return 1
		//
		exitBlock_shift_expr_ShrOp_add_expr: // exit block
	}
	| shift_expr ShlOp add_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("shift_expr/shift_expr ShlOp add_expr", $1, $3)
}
		v1, ok1 := $1.(IInt64Value)
		v3, ok3 := $3.(IInt64Value)
		if ok1 && ok3 {
			if v3.Int64Value() < 0 {
				YaccIdllex.Error("shift count must be positive")
			}
			$$, _ = NewInt64Value(v1.Int64Value() << v3.Int64Value())
			//
			goto exitBlock_shift_expr_ShlOp_add_expr
		}
		YaccIdllex.Error("type not supported (<< operator)")
		return 1
		//
		exitBlock_shift_expr_ShlOp_add_expr: // exit block
	}
//(12)
add_expr:
	mult_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("add_expr/mult_expr", $1)
}
		var err error
		$$ = $1
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| add_expr '+' mult_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("add_expr/add_expr '+' mult_expr", $1, $3)
}
		if v1, ok1 := $1.(IInt64Value); ok1 {
			if v3, ok3 := $3.(IInt64Value); ok3{
				if ok1 && ok3 {
					$$, _ = NewInt64Value(v1.Int64Value() + v3.Int64Value())
					//
					goto exitBlock_add_expr_plus_mult_expr
				}
			}
		}
		if v1, ok1 := $1.(IStringValue); ok1 {
			if v3, ok3 := $3.(IStringValue); ok3{
				if ok1 && ok3 {
					$$, _ = NewStringValue(v1.StringValue() + v3.StringValue())
					//
					goto exitBlock_add_expr_plus_mult_expr
				}
			}
		}
		if v1, ok1 := $1.(IFloat64Value); ok1 {
			if v3, ok3 := $3.(IFloat64Value); ok3{
				if ok1 && ok3 {
					$$, _ = newFloat64Value(v1.FloatValue() + v3.FloatValue())
					//
					goto exitBlock_add_expr_plus_mult_expr
				}
			}
		}

		YaccIdllex.Error("type not supported (+ operator)")
		return 1
		//
		exitBlock_add_expr_plus_mult_expr:
	}
	| add_expr '-' mult_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                	yaccLogger.Log("add_expr/add_expr '-' mult_expr", $1, $3)
                }
		if v1, ok1 := $1.(IInt64Value); ok1 {
			if v3, ok3 := $3.(IInt64Value); ok3{
                		if ok1 && ok3 {
                			$$, _ = NewInt64Value(v1.Int64Value() - v3.Int64Value())
                			//
                			goto exitBlock_add_expr_minus_mult_expr
                		}
                	}
		}
		if v1, ok1 := $1.(IFloat64Value); ok1 {
			if v3, ok3 := $3.(IFloat64Value); ok3{
				if ok1 && ok3 {
					$$, _ = newFloat64Value(v1.FloatValue() - v3.FloatValue())
					//
					goto exitBlock_add_expr_minus_mult_expr
				}
			}
		}
		YaccIdllex.Error("type not supported (- Operator)")
		return 1
		//
		exitBlock_add_expr_minus_mult_expr:
	}
//(13)
mult_expr:
	unary_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("mult_expr/unary_expr", $1)
		}
		$$ = $1
	}
	| mult_expr '*' unary_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("mult_expr/mult_expr '*' unary_expr", $1, $3)
}
		if v1, ok1 := $1.(IInt64Value); ok1 {
			if v3, ok3 := $3.(IInt64Value); ok3 {
				$$, _ = NewInt64Value(v1.Int64Value() * v3.Int64Value())
				//
				goto exitBlock_mult_expr_star_unary_expr
			}
		}
		if v1, ok1 := $1.(IFloat64Value); ok1 {
			if v3, ok3 := $3.(IFloat64Value); ok3 {
				$$, _ = newFloat64Value(v1.FloatValue() * v3.FloatValue())
				//
				goto exitBlock_mult_expr_star_unary_expr
			}
		}

		YaccIdllex.Error("type not supported (* operator)")
		return 1
		//
		exitBlock_mult_expr_star_unary_expr:
	}
	| mult_expr '/' unary_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("mult_expr/mult_expr '/' unary_expr", $1, $3)
}
		if v1, ok1 := $1.(IInt64Value); ok1 {
			if v3, ok3 := $3.(IInt64Value); ok3{
				$$, _ = NewInt64Value(v1.Int64Value() / v3.Int64Value())
				//
				goto exitBlock_mult_expr_divide_unary_expr
			}
		}
		if v1, ok1 := $1.(IFloat64Value); ok1 {
			if v3, ok3 := $3.(IFloat64Value); ok3{
				$$, _ = newFloat64Value(v1.FloatValue() / v3.FloatValue())
				//
				goto exitBlock_mult_expr_divide_unary_expr
			}
		}

		YaccIdllex.Error("type not supported (/ operator)")
		return 1
		//
		exitBlock_mult_expr_divide_unary_expr:
	}
	| mult_expr '%' unary_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("mult_expr/mult_expr '%' unary_expr", $1, $3)
}
		if v1, ok1 := $1.(IInt64Value); ok1 {
			if v3, ok3 := $3.(IInt64Value); ok3 {
				$$, _ = NewInt64Value(v1.Int64Value() % v3.Int64Value())
				//
				goto exitBlock_mult_expr_perc_unary_expr
			}
		}
		YaccIdllex.Error("type not supported (% operator)")
		return 1
		//
		exitBlock_mult_expr_perc_unary_expr:
	}
//(14)
unary_expr:
	unary_operator primary_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("unary_expr/unary_operator primary_expr", $1, $2)
}
		if v1, ok1 := $1.(IInt64Value); ok1 {
			if v2, ok2 := $2.(IInt64Value); ok2 {
				$$, _ = NewInt64Value(v1.Int64Value() * v2.Int64Value())
				//
				goto exitBlock_unary_operator_primary_expr
			}
		}
		YaccIdllex.Error("type not supported (+- operator)")
		return 1
		//
		exitBlock_unary_operator_primary_expr:
	}
	| primary_expr
	{
		$$ = $1
	}
//(15)
unary_operator:
	'-'
	{
		$$, _ = NewInt64Value(-1)
	}
	| '+'
	{
		$$, _ = NewInt64Value(+1)
	}
	| '~'
	{
		YaccIdllex.Error("Impelemnt me!! '~'")
		return 1
	}
//(16)
primary_expr:
	scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("primary_expr/scoped_name", $1)
}
		$$ = newScopeNameValue($1.Identifier())
	}
	| literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("primary_expr/literal", $1)
}
		$$ = $1
	}
	| '(' const_expr ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("primary_expr/'(' const_expr ')'", $2)
}
		$$ = $2
	}
//(17)
literal:
	Integer_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("literal/Integer_literal", $1)
}
		$$ = $1
	}
	| Floating_pt_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("literal/Floating_pt_literal", $1)
}
		$$ = $1
	}
	| Fixed_pt_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("literal/Fixed_pt_literal", $1)
}
		$$ = $1
	}
	| Character_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("literal/Character_literal", $1)
}
		var err error
		$$, err = NewCharValue($1)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| Wide_Character_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("literal/Wide_Character_literal", $1)
}
		var err error
		$$, err = NewWideCharValue($1)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| boolean_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("literal/boolean_literal", $1)
		}
		var err error
		$$, err = NewBooleanValue($1)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| String_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("literal/String_literal", $1)
		}
		var err error
		$$, err = NewStringValue($1.Identifier())
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| Wide_String_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("literal/Wide_String_literal", $1)
}
		var err error
		$$, err = NewWideStringValue($1)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(18)
boolean_literal:
	RWTRUE
	{
		$$ = true
	}
	| RWFALSE
	{
		$$ = false
	}
//(19)
positive_int_const:
	const_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("positive_int_const/const_expr", $1)
		}
		v1, ok1 := $1.(IInt64Value)
		if ok1 {
			if v1.Int64Value() < 0 {
				YaccIdllex.Error("expression must be positive")
				return 1
			}
			$$ = $1
			//
			goto exitBlock_positive_int_const_const_expr
		}
		YaccIdllex.Error("type not supported (const_expr)")
		return 1
		//
		exitBlock_positive_int_const_const_expr:
	}
//(20)
type_dcl:
	constr_type_dcl
	{
		$$ = $1
	}
	| native_dcl
	{
		$$ = $1
	}
	| typedef_dcl
	{
		$$ = $1
	}
//(21)
type_spec:
	simple_type_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("type_spec/simple_type_spec", $1)
}
		$$ = $1
	}
	|template_type_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("type_spec/template_type_spec")
}
		$$ = $1
	}
//(22)
// ISimpleTypeSpec
simple_type_spec:
	base_type_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("simple_type_spec/base_type_spec", $1)
		}
		$$ =$1
	}
	| scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("simple_type_spec/scoped_name", $1)
		}
		$$ =$1
	}
//(23)(69)//(117)//(131)
base_type_spec:
	integer_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/integer_type", $1)
		}
        	$$ = $1
        }
	| floating_pt_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/floating_pt_type", $1)
		}
		$$ = $1
	}
	| char_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/char_type", $1)
		}
		$$ = $1
	}
	| wide_char_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/wide_char_type", $1)
		}
		$$ = $1
	}
	| boolean_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/boolean_type", $1)
		}
		$$ = $1
	}
	| octet_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/octet_type", $1)
		}
		$$ =$1
	}
	| any_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/any_type", $1)
		}
		$$ =$1
	}
	| object_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/object_type", $1)
		}
		$$ =$1
	}
	| value_base_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("base_type_spec/value_base_type", $1)
		}
		$$ =$1
	}

//(24)
floating_pt_type:
	RWfloat {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("floating_pt_type/RWfloat")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlFloatType); ok {
			$$ = defaultTypes.FloatType()
			if $$ == nil {
				YaccIdllex.Error("no float primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlFloatType interface")
			return 1
		}
	}
	| RWdouble {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("floating_pt_type/RWdouble")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlDoubleType); ok {
			$$ = defaultTypes.DoubleType()
			if $$ == nil {
				YaccIdllex.Error("no double primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlDoubleType interface")
			return 1
		}
	}
	| RWlong RWdouble {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("floating_pt_type/RWlong RWdouble")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlLongDoubleType); ok {
			$$ = defaultTypes.LongDoubleType()
			if $$ == nil {
				YaccIdllex.Error("no long double primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlLongDoubleType interface")
			return 1
		}
	}
//(25)
integer_type:
	signed_int
	{
		$$ = $1
	}
	| unsigned_int
	{
		$$ = $1
	}

//(26)//(206)
signed_int:
	signed_short_int {
		$$ =$1
	}
	| signed_long_int {
		$$ =$1
	}
	| signed_longlong_int {
		$$ =$1
	}
	| signed_tiny_int {
		$$ =$1
	}
//(27)//(210)
signed_short_int:
	RWshort
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("signed_short_int/RWshort")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlShortType); ok {
			$$ = defaultTypes.ShortType()
			if $$ == nil {
				YaccIdllex.Error("no short primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no ShortType interface")
			return 1
		}
	}
	| RWint16
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("signed_short_int/RWint16")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlInt16Type); ok {
			$$ = defaultTypes.Int16Type()
			if $$ == nil {
				YaccIdllex.Error("no Int16 primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlInt16Type interface")
			return 1
		}
	}
//(28)//(211)
signed_long_int:
	RWlong
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("signed_long_int/RWlong", $1)
		}
		if defaultTypes, ok := YaccIdllex.(IIdlLongType); ok {
			$$ = defaultTypes.LongType()
			if $$ == nil {
				YaccIdllex.Error("no long primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlLongType interface")
			return 1
		}
	}
	| RWint32
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("signed_long_int/RWint32")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlInt32Type); ok {
			$$ = defaultTypes.Int32Type()
			if $$ == nil {
				YaccIdllex.Error("no Int32 primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlInt32Type interface")
			return 1
		}
	}

//(29)//(212)
signed_longlong_int:
	RWlong RWlong
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("signed_longlong_int/RWlong RWlong")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlLongLongType); ok {
			$$ = defaultTypes.LongLongType()
			if $$ == nil {
				YaccIdllex.Error("no long long primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlLongLongType interface")
			return 1
		}
	}
	|RWint64
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("signed_longlong_int/RWint64")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlInt64Type); ok {
			$$ = defaultTypes.Int64Type()
			if $$ == nil {
				YaccIdllex.Error("no Int64 primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlInt64Type interface")
			return 1
		}


	}

//(30)//(207)
unsigned_int:
	unsigned_short_int
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("unsigned_int/unsigned_short_int", $1)
}
		$$ = $1
	}
	| unsigned_long_int {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("unsigned_int/unsigned_long_int", $1)
}
		$$ = $1
	}
	| unsigned_longlong_int {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("unsigned_int/unsigned_longlong_int", $1)
}
		$$ = $1
	}
	|unsigned_tiny_int {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("unsigned_int/unsigned_tiny_int", $1)
}
		$$ = $1
	}

//(31)//(213)
unsigned_short_int:
	RWunsigned RWshort
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("unsigned_short_int/RWunsigned RWshort")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlUnsignedShortType); ok {
			$$ = defaultTypes.UnsignedShortType()
			if $$ == nil {
				YaccIdllex.Error("no unsigned short primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlUnsignedShortType interface")
			return 1
		}
	}
	|RWuint16
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("unsigned_short_int/RWuint16")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlUint16Type); ok {
			$$ = defaultTypes.Uint16Type()
			if $$ == nil {
				YaccIdllex.Error("no Uint16 primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlUint16Type interface")
			return 1
		}
	}

//(32)//(214)
unsigned_long_int:
	RWunsigned RWlong
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("unsigned_long_int/RWunsigned RWlong")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlUnsignedLongType); ok {
			$$ = defaultTypes.UnsignedLongType()
			if $$ == nil {
				YaccIdllex.Error("no unsigned long primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlUnsignedLongType interface")
			return 1
		}
	}
	| RWuint32
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("unsigned_long_int/RWuint32")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlUint32Type); ok {
			$$ = defaultTypes.Uint32Type()
			if $$ == nil {
				YaccIdllex.Error("no Uint32 primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlUint32Type interface")
			return 1
		}
	}

//(33)//(215)
unsigned_longlong_int:
	RWunsigned RWlong RWlong
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("unsigned_longlong_int/RWunsigned RWlong RWlong")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlUnsignedLongLongType); ok {
			$$ = defaultTypes.UnsignedLongLongType()
			if $$ == nil {
				YaccIdllex.Error("no unsigned long long  primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlUnsignedLongLongType interface")
			return 1
		}
	}
	| RWuint64
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("unsigned_longlong_int/RWuint64")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlUint64Type); ok {
			$$ = defaultTypes.Uint64Type()
			if $$ == nil {
				YaccIdllex.Error("no Uint64 primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlUint64Type interface")
			return 1
		}
	}

//(34)
char_type:
	RWchar
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("char_type/RWchar")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlCharType); ok {
			$$ = defaultTypes.CharType()
			if $$ == nil {
				YaccIdllex.Error("no char primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlCharType interface")
			return 1
		}
	}
//(35)
wide_char_type:
	RWwchar {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("wide_char_type/RWwchar")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlWideCharType); ok {
			$$ = defaultTypes.WideCharType()
			if $$ == nil {
				YaccIdllex.Error("no wide char primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlWideCharType interface")
			return 1
		}
	}
//(36)
boolean_type:
	RWboolean
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("boolean_type/RWboolean")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlBooleanType); ok {
			$$ = defaultTypes.BooleanType()
			if $$ == nil {
				YaccIdllex.Error("no boolean primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlBooleanType interface")
			return 1
		}
	}
//(37)
octet_type:
	RWoctet
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("octet_type/RWoctet")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlOctetType); ok {
			$$ = defaultTypes.OctetType()
			if $$ == nil {
				YaccIdllex.Error("no octet primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlOctetType interface")
			return 1
		}
	}
//(38)//(197)
template_type_spec:
	sequence_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		yaccLogger.Log("template_type_spec/sequence_type")
}
		$$ =$1
	}
	| string_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		yaccLogger.Log("template_type_spec/string_type")
}
		$$ =$1
	}
	| wide_string_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		yaccLogger.Log("template_type_spec/wide_string_type")
}
		$$ =$1
	}
	| fixed_pt_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		yaccLogger.Log("template_type_spec/fixed_pt_type")
}
		$$ =$1
	}
	| map_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		yaccLogger.Log("template_type_spec/map_type")
}
		$$ =$1
	}
//(39)
sequence_type:
	RWsequence '<' type_spec ',' positive_int_const '>'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		yaccLogger.Log("sequence_type/RWsequence '<' type_spec ',' positive_int_const '>'", $3, $5)
}
		if v5, ok5 := $5.(IInt64Value); ok5 {
			if v5.Int64Value() <= 0 {
				YaccIdllex.Error("need a positive int")
				return 1
			}
			var err error
			$$, err = newSequenceType(__yyfmt__.Sprintf("%v", $3.Identifier()), $1.LexemData(), $3, v5.Int64Value())
			if err != nil {
				YaccIdllex.Error(err.Error())
				return 1
			}
		}
	}
	| RWsequence '<' type_spec '>'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		yaccLogger.Log("sequence_type/RWsequence '<' type_spec '>'", $3)
}
		var err error
		$$, err = newSequenceType(__yyfmt__.Sprintf("%v", $3.Identifier()), $1.LexemData(), $3, 0)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(40)
string_type:
	RWstring '<' positive_int_const '>'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		yaccLogger.Log("string_type/RWstring '<' positive_int_const '>'", $3)
}
		if v3, ok3 := $3.(IInt64Value); ok3 {
			if v3.Int64Value() <= 0 {
				YaccIdllex.Error("need a positive int")
				return 1
			}
			var err error
			$$, err = newStringType(v3.Int64Value())
			if err != nil {
				YaccIdllex.Error(err.Error())
				return 1
			}
		}
	}
	| RWstring
	{
		if defaultTypes, ok := YaccIdllex.(IIdlStringType); ok {
			$$ = defaultTypes.StringType()
			if $$ == nil {
				YaccIdllex.Error("no string primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlStringType interface")
			return 1
		}
	}
//(41)
wide_string_type:
	RWwstring '<' positive_int_const '>'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("wide_string_type/RWwstring '<' positive_int_const '>'", $3)
		}
		if v3, ok3 := $3.(IInt64Value); ok3 {
			if v3.Int64Value() <= 0 {
				YaccIdllex.Error("need a positive int")
				return 1
			}
			var err error
			$$, err = newWideStringType(v3.Int64Value())
			if err != nil {
				YaccIdllex.Error(err.Error())
				return 1
			}
		}
	}
	| RWwstring
	{
		if defaultTypes, ok := YaccIdllex.(IIdlWideStringType); ok {
			$$ = defaultTypes.WideStringType()
			if $$ == nil {
				YaccIdllex.Error("no wide string primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlWideStringType interface")
			return 1
		}
	}
//(42)
fixed_pt_type:
	RWfixed '<' positive_int_const ',' positive_int_const '>'
	{
//		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
//yaccLogger.Log("fixed_pt_type/RWfixed '<' positive_int_const ',' positive_int_const '>'", $3)
//}
//		if $3 <= 0 {
//			YaccIdllex.Error("need a positive int")
//			return 1
//		}
//		if $5 <= 0 {
//			YaccIdllex.Error("need a positive int")
//			return 1
//		}
//		var err error
//		$$, err = scopedObjects.NewFixedPointType($1, $3, $5)
//		if err != nil {
//			YaccIdllex.Error(err.Error())
//			return 1
//		}
	}
//(43)
fixed_pt_const_type:
	RWfixed
	{
//		YaccIdllex.Error("fix this")
//		return 1
	}
//(44)//(198)
constr_type_dcl:
	struct_dcl
	{
		$$ = $1
	}
	| union_dcl
	{
		$$ = $1
	}
	| enum_dcl
	{
		$$ = $1
	}
	| bitset_dcl
	{
		$$ = $1
	}
	| bitmask_dcl
	{
		$$ = $1
	}

//(45)
struct_dcl:
	struct_def
	{
		$$ = $1
	}
	| struct_forward_dcl
	{
	 	$$ = $1
	}
//(46)//(195)
struct_def:
	RWstruct Identifier '{' memberPlus '}'{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("RWstruct Identifier '{' memberPlus '}'", $1, $2, $4)
		}
		var err error
		$$, err = NewStructType($2, nil, $4, false)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| RWstruct Identifier ':' scoped_name '{' memberStar '}'
	{
		var err error
		$$, err =NewStructType($2, $4, $6, false)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| RWstruct Identifier '{' '}'
	{
		var err error
		$$, err = NewStructType($2, nil, nil, false)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}

//(47)
memberPlus:
	member
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("memberPlus/member")
		}
		$$ = $1
	}
	|memberPlus member
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("memberPlus/memberPlus member")
		}
		node := $1
		for node.GetNextNode() != nil{
			node = node.GetNextNode()
		}
		_ = node.SetNextNode($2)
		$$ = $1
	}
memberStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("memberStar/")
}
		$$ = nil
	}
	|memberPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("memberStar/memberPlus")
}
		$$ = $1
	}

member:
	type_spec declarators ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("member/type_spec declarators", $1, $2)
		}
		var err error
		$$, err = NewStructMember($1, $2)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(48)
struct_forward_dcl:
	RWstruct Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("struct_forward_dcl/RWstruct Identifier")
}
		var err error
		$$, err = NewStructType($2, nil, nil, true)
		if err != nil{
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(49)
union_dcl:
	union_def
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("union_dcl/union_def")
		}
		$$ = $1
	}
	| union_forward_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("union_dcl/union_forward_dcl")
		}
		$$ = $1
	}
//(50)
union_def:
	RWunion Identifier RWswitch '(' switch_type_spec ')' '{' switch_body '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("union_def/RWunion Identifier RWswitch '(' switch_type_spec ')' '{' switch_body '}'")
		}
		$$ = NewUnionDcl($2, $5, $8, false)
	}
//(51)//(196)
switch_type_spec:
	integer_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("switch_type_spec/integer_type")
		}
		$$ = $1
	}
	| char_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("switch_type_spec/char_type")
		}
		$$ = $1
	}
	| boolean_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("switch_type_spec/boolean_type")
		}
		$$ = $1
	}
	| scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("switch_type_spec/scoped_name")
		}
		$$ = $1
	}
	| wide_char_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("switch_type_spec/wide_char_type")
		}
		$$ = $1
	}
	| octet_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("switch_type_spec/octet_type")
		}
		$$ = $1
	}
//(52)
switch_body:
	casePlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("switch_body/casePlus")
		}
		$$ = $1
	}
casePlus:
	case
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("casePlus/case")
		}
		$$ = $1
	}
	|casePlus case
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("casePlus/casePlus case")
		}
		node := $1
		for node.GetNextNode() != nil{
			var err error
			node, err  = node.GetNextUnionBody()
			if err != nil {
				YaccIdllex.Error(err.Error())
				return 1
			}
		}
		node.SetNextUnionBody($2)
		$$ = $1
	}
//(53)
case:
	case_labelPlus element_spec ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("case/case_labelPlus element_spec ';'")
		}
		$$ = NewUnionBody($1.LexemData(), "UnionCaseLabel", $1, $2)
	}

case_labelPlus:
	case_label
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("case_labelPlus/case_label")
		}
		$$ = $1
	}
	|case_labelPlus case_label
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("case_labelPlus/ase_labelPlus case_label")
		}
		node := $1
		for node.GetNextNode() != nil{
			var err error
			node, err = node.GetNextConstValue()
			if err != nil {
				YaccIdllex.Error(err.Error())
				return 1
			}
		}
		node.SetNextConstValue($2)
		$$ = $1
	}
//(54)
case_label:
	RWcase const_expr ':'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("case_label/RWcase const_expr ':'")
		}
		$$ = NewConstValue("constvalue", $1.LexemData(), $2)

	}
	| RWdefault ':'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("case_label/RWdefault ':'")
		}
		$$ = NewDefaultValue("default", $1.LexemData())
	}
//(55)
element_spec:
	type_spec declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("element_spec/type_spec declarator")
		}
		$$ = NewElementSpec($1, $2)
	}
//(56)
union_forward_dcl:
	RWunion Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("union_forward_dcl/RWunion Identifier")
		}
		$$ = NewUnionDcl($2, nil, nil, true)
	}


//(57)
enum_dcl:
	RWenum Identifier '{' enumerators '}'
	{
		var err error
		$$, err = newEnumDcl($2, $4)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
enumerators:
	enumerator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("enumerators/enumerator| enumerators ',' enumerator", $1)
		}
		$$ = $1
	}
	| enumerators ',' enumerator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("enumerators/enumerator| enumerators ',' enumerator", $1, $3)
		}
		node := $1
		for node.GetNextNode() != nil{
			node = node.GetNextNode()
		}
		node.SetNextNode($3)
		$$ = $1
	}
//(58)
enumerator:
	Identifier {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("enumerator/Identifier", $1)
}
		$$ = $1
	}
//(59)
array_declarator:
	simple_declarator fixed_array_sizePlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("array_declarator/Identifier fixed_array_sizePlus", $1, $2)
		}
		var a []int64 = nil
		for node := $2; node != nil; node = node.GetNextNode(){
			if v1, ok1 := node.(IInt64Value); ok1{
				a = append(a, v1.Int64Value())
			}
		}
		$1.AddArray(a)
		$$ = $1
	}
//(60)
fixed_array_sizePlus:
	fixed_array_size
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("fixed_array_size", $1)
		}
		$$ = $1
	}
	|fixed_array_sizePlus fixed_array_size
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("fixed_array_sizePlus fixed_array_size", $1, $2)
		}
		node := $1
		for node.GetNextNode() != nil{
			node = node.GetNextNode()
		}
		node.SetNextNode($2)
		$$ = $1
	}
fixed_array_size:
	'[' positive_int_const ']'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("'[' positive_int_const ']'", $2)
		}
		$$ = $2
	}
//(61)
native_dcl:
	RWnative simple_declarator
	{
		$$, _ = newNativeDcl($1.LexemData(), $2.Identifier())
	}
//(62)
simple_declaratorPlus:
	simple_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("simple_declaratorPlus/simple_declarator")
		}
		$$ = $1
	}
	|simple_declaratorPlus ',' simple_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("simple_declaratorPlus/simple_declaratorPlus ',' simple_declarator")
		}
		node := $1
		for node.GetNextNode() != nil{
			node = node.GetNextNode()
		}
		node.SetNextNode($3)
		$$ = $1
	}

simple_declarator:
	Identifier
	{
		$$ = $1
	}
//(63)
typedef_dcl:
	RWtypedef type_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("typedef_dcl/RWtypedef type_declarator", $2)
		}
		$$ = $2
	}
//(64)
// ITypeDeclarator interface
type_declarator:
	simple_type_spec any_declaratorsPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("type_declarator/simple_type_spec any_declarators", $1, $2)
		}
		$$, _ = newTypeDeclarator($1.LexemData(), $1.Identifier(), $1, $2)
	}
	| template_type_spec any_declaratorsPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("type_declarator/template_type_spec any_declarators", $1, $2)
}
		$$, _ = newTypeDeclarator($1.LexemData(), $1.Identifier(),$1, $2)
	}
	| constr_type_dcl any_declaratorsPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("type_declarator/constr_type_dcl any_declarators", $1, $2)
}
		$$, _ = newTypeDeclarator($2.LexemData(), $1.Identifier(),$1, $2)
	}
//(65)
any_declaratorsPlus:
	any_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("any_declaratorsPlus/any_declarator", $1)
}
		$$ = $1
	}
	| any_declaratorsPlus ',' any_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("any_declaratorsPlus/any_declaratorsPlus ',' any_declarator", $1, $3)
}
		node := $1
		for node.GetNextNode() != nil{
			node = node.GetNextNode()
		}
		_ = node.SetNextNode($3)
	}

//(66)
any_declarator:
	simple_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("declarator/any_declarator:", $1)
		}
		$$ = $1
	}
	| array_declarator
	{

	}
//(67)
declarators:
	declaratorPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("declaratorPlus", $1)
		}
		$$ = $1
	}
//(68)//(217)
declaratorPlus:
	declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("declaratorPlus/declarator", $1)
		}
		$$ = $1
	}
	| declaratorPlus ',' declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("declaratorPlus/declaratorPlus ',' declarator", $1, $3)
		}
		// find last
		node := $1
		for node.GetNextNode() != nil{
			node = node.GetNextNode()
		}
		//assign
		node.SetNextNode($3)
		$$ = $1
	}
declarator:
	simple_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("declarator/simple_declarator", $1)
		}
		$$ = $1
	}
	|array_declarator
	{
         	if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("array_declarator", $1)
		}
		$$ = $1
        }
//(70)
any_type:
	RWany
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("any_type/RWany")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlAnyType); ok {
			$$ = defaultTypes.AnyType()
			if $$ == nil {
				YaccIdllex.Error("no any primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlAnyType interface")
			return 1
		}
	}
//(72)
except_dcl:
	RWexception Identifier '{' memberStar '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("RWexception Identifier '{' memberStar '}'", $1, $2)
}
		var err error
		$$, err = newExceptionDcl($2, $4)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(73)
interface_dcl:
	interface_def
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("interface_dcl/interface_def", $1)
}
		$$ = $1
	}
	|
	interface_forward_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("interface_dcl/interface_forward_dcl", $1)
}

		var err error
		$$, err = newInterfaceDcl($1.LexemData(), $1.Identifier(), $1, nil, true)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(74)
interface_def:
	interface_header '{' interface_body '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("interface_def/interface_header '{' interface_body '}'", $1, $3)
}
		var err error
		$$, err = newInterfaceDcl($1.LexemData(), $1.Identifier(), $1, $3, false)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(75)
interface_forward_dcl:
	interface_kind Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
yaccLogger.Log("interface_forward_dcl/interface_kind Identifier", $2)
}
		var err error
		$$, err = newInterfaceHeader($1.LexemData(), $2.Identifier(), $1, nil)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(76)
interface_header:
	interface_kind Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_header/interface_kind Identifier", $1, $2)
		}
		var err error
		$$, err = newInterfaceHeader($1.LexemData(), $2.Identifier(), $1, nil)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}

	|
	interface_kind Identifier interface_inheritance_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_header/interface_kind Identifier interface_inheritance_spec", $1, $2, $3)
		}
		var err error
		$$, err = newInterfaceHeader($2.LexemData(), $2.Identifier(), $1, $3)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}

	}
//(77)//(119)//(129)
interface_kind:
	RWinterface
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_kind/RWinterface", $1)
		}
		var err error
		$$, err = newInterfaceKind($1.LexemData(), false, false, false)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| RWlocal RWinterface
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_kind/RWlocal RWinterface", $1)
		}
		var err error
		$$, err = newInterfaceKind($2.LexemData(), true, false, false)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	|RWabstract RWinterface
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_kind/RWabstract RWinterface", $1)
		}
		var err error
		$$, err = newInterfaceKind($2.LexemData(), false, true, false)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}


//(78)
interface_inheritance_spec:
	':' interface_namePlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_inheritance_spec/':' interface_namePlus")
		}
		$$ = $2
	}
//(79)
interface_namePlus:
	interface_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_namePlus/interface_name")
		}
		$$ = $1
	}
	| interface_namePlus ',' interface_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_namePlus/interface_namePlus ',' interface_name")
		}
		node := $1
		for node.GetNextNode() != nil{
			node = node.GetNextNode()
		}
		_ = node.SetNextNode($3)
	}
interface_name:
	scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_name/scoped_name")
		}
		$$ = $1
	}
//(80)
interface_body:
	exportStar
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("interface_body/exportStar")
		}
		$$ = $1
	}
//(81)//(97)//(112)
exportStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("exportStar")
		}
		$$ = nil
	}
	| exportPlus
	{
		$$ = $1
	}

exportPlus:
	exportPlus export {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("exportPlus/exportPlus export", $1, $2)
		}
		if $1 != nil && $2 != nil {
			node := $1
			for node.GetNextNode() != nil {
				node = node.GetNextNode()
			}
			_ = node.SetNextNode($2)
		}else if $1 == nil && $2 != nil {
			$$ = $2
                }else if $1 != nil && $2 == nil {
                	$$ = $1
                }
	}
	| export {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("exportPlus/export", $1)
		}
		$$ = $1
	}
export:
	op_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/op_dcl ';'", $1)
		}
		$$ = $1
	}
	| attr_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/attr_dcl ';'", $1)
		}
		$$ = $1
	}
	| type_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/type_dcl ';'", $1)
		}
		$$ = $1
	}
	| const_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/const_dcl ';'", $1)
		}
		$$ = $1
	}
	| except_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/except_dcl ';'", $1)
		}
		$$ = $1
	}
	| type_id_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/type_id_dcl ';'", $1)
		}
		$$ = $1
	}
	| type_prefix_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/type_prefix_dcl ';'", $1)
		}
		$$ = $1
	}
	| import_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/import_dcl ';'", $1)
		}
		$$ = $1
	}
	| op_oneway_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/op_oneway_dcl ';'", $1)
		}
		$$ = $1
	}
	| op_with_context ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("export/op_with_context ';'", $1)
		}
		$$ = $1
	}

//(82)
op_dcl:
	op_type_spec Identifier '('  ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("op_dcl/op_type_spec Identifier '('  ')'", $1, $2)
		}
		var err error
		$$, err = newOperationDcl($2, nil, $1, nil)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	|op_type_spec Identifier '('  ')'  raises_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("op_dcl/op_type_spec Identifier '('  ')'  raises_expr", $1, $2, $5)
		}
		var err error
		$$, err = newOperationDcl($2, nil, $1, $5)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	|op_type_spec Identifier '('  parameter_dcls ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("op_dcl/op_type_spec Identifier '('  parameter_dcls ')'", $1, $2, $4)
		}
		var err error
		$$, err = newOperationDcl($2, $4, $1, nil)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	|op_type_spec Identifier '('  parameter_dcls ')'  raises_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("op_dcl/op_type_spec Identifier '('  parameter_dcls ')'  raises_expr", $1, $2, $4, $6)
		}
		var err error
		$$, err = newOperationDcl($2, $4, $1, $6)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(83)
op_type_spec:
	type_spec
	{
		$$ = $1
	}| RWvoid
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("op_type_spec/RWvoid")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlVoidType); ok {
			$$ = defaultTypes.VoidType()
			if $$ == nil {
				YaccIdllex.Error("no void primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlVoidType interface")
			return 1
		}
	}
//(84)
parameter_dcls:
	param_dclPlus
	{
		$$ = $1
	}
//(85)
param_dclPlus:
	param_dcl {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("param_dclPlus/param_dcl")
		}
		$$ = $1
	}
	| param_dclPlus ',' param_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("param_dclPlus/param_dclPlus ',' param_dcl")
		}
		node := $1
		for node.GetNextNode() != nil {
			node = node.GetNextNode()
		}
		_ = node.SetNextNode($3)
	}
param_dcl:
	param_attribute type_spec simple_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("param_dcl/param_attribute type_spec simple_declarator")
		}
		var err error
		$$, err = newOperationalParameter($3.Identifier(), $2.LexemData(), $1, $2)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(86)
param_attribute:
	RWin
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("param_attribute/RWin")
		}
		$$ = ParamDirectionIn
	}
	| RWout
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("param_attribute/RWout")
		}
		$$ = ParamDirectionOut
	}
	| RWinout
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("param_attribute/RWinout")
		}
		$$ = ParamDirectionOut | ParamDirectionIn
	}
//(87)
raises_expr:
	RWraises '(' scoped_namePlus ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("raises_expr/RWraises '(' scoped_namePlus ')'", $3)
		}
		$$ = $3
	}
//(88)
attr_dcl:
	readonly_attr_spec
	{
//		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
//yaccLogger.Log("attr_dcl/readonly_attr_spec")
//}
//		$$ = $1
	}
	| attr_spec
	{
//		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
//yaccLogger.Log("attr_dcl/attr_spec")
//}
//		$$ = $1
	}
//(89)
readonly_attr_spec:
	RWreadonly RWattribute type_spec readonly_attr_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("readonly_attr_spec/RWreadonly RWattribute type_spec readonly_attr_declarator")
		}
		var err error
		$$, err = newAttributeDcl(
		$2.LexemData(),
		$4.SimpleDeclarator().Identifier(),
		$3, $4, true)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(90)
readonly_attr_declarator:
	simple_declarator raises_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("readonly_attr_declarator/simple_declarator raises_expr")
		}
		var err error
		$$, err = newReadonlyAttrDeclarator($1, $2)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| simple_declaratorPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("readonly_attr_declarator/simple_declaratorPlus")
		}
		var err error
		$$, err = newReadonlyAttrDeclarator($1, nil)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(91)
attr_spec:
	RWattribute type_spec attr_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("attr_spec/RWattribute type_spec attr_declarator")
		}
		var err error
		$$, err = newAttributeDcl(
			$1.LexemData(),
			$3.SimpleDeclarator().Identifier(),
			$2,
			$3,
			false)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}

	}
//(92)
attr_declarator:
	simple_declarator attr_raises_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("readonly_attr_declarator/simple_declarator raises_expr")
		}
		var err error
		$$, err = newReadonlyAttrDeclarator($1, $2)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	| simple_declaratorPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("readonly_attr_declarator/simple_declarator raises_expr")
		}
		var err error
		$$, err = newReadonlyAttrDeclarator($1, nil)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
//(93)
attr_raises_expr:
	get_excep_expr  set_excep_expr
	{
//		
	}
	| get_excep_expr
	{
//		
	}
	| set_excep_expr
	{
//		
	}
//(94)
get_excep_expr:
	RWgetraises exception_list
	{
//		
	}
//(95)
set_excep_expr:
	RWsetraises exception_list
	{
//		
	}
//(96)
exception_list:
 	'(' scoped_namePlus ')'
	{
//		
	}
//(99)//(125)
value_dcl:
	value_def
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_dcl/value_def")
		}
		$$ = $1
	}
	| value_forward_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_dcl/value_forward_dcl")
		}
		$$, _ = newValueDef($1.LexemData(), $1.Identifier(), $1, nil, true)
	}
	| value_box_def
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_dcl/value_box_def")
		}
		$$ = $1
	}
//	| value_abs_def
//	{
//		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
		//yaccLogger.Log("value_dcl/value_abs_def")
//}
//		$$ = $1
//	}
//(100)
value_def:
	value_header '{' value_elementStar '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_header '{' value_elementPlus '}'")
		}
		$$, _ = newValueDef($1.LexemData(), $1.Identifier(), $1,$3, false)
	}
//(101)
value_header:
	value_kind Identifier value_inheritance_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_header/value_kind Identifier value_inheritance_spec")
		}
		$$, _ = newValueHeader($1, $2, $3)
	}
	| value_kind Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_header/value_kind Identifier")
		}
		$$, _ = newValueHeader($1, $2, nil)
	}
//(102) //(128) //(127)
value_kind:
	RWvaluetype
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_kind/RWvaluetype")
		}
		var err error
		$$, err = newInterfaceKind($1.LexemData(), false, false, false)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	|RWcustom RWvaluetype
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_kind/RWcustom RWvaluetype")
		}
		var err error
		$$, err = newInterfaceKind($2.LexemData(), false, false, true)
		if err != nil {
			YaccIdllex.Error(err.Error())
			return 1
		}
	}
	|RWabstract RWvaluetype
        	{
        		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_kind/RWcustom RWvaluetype")
		}
        		var err error
        		$$, err = newInterfaceKind($2.LexemData(), false, true, false)
        		if err != nil {
        			YaccIdllex.Error(err.Error())
        			return 1
        		}
        	}

//(103)//(130)
value_inheritance_spec:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_inheritance_spec")
		}
		$$, _ = newValueInheritanceSpec(nil, nil)
	}
	|   RWsupports interface_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_inheritance_spec/RWsupports interface_name")
		}
		$$, _ = newValueInheritanceSpec($2, nil)
	}
	|  ':' value_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_inheritance_spec/':' value_name")
		}
		$$, _ = newValueInheritanceSpec(nil, $2)
	}
	|  ':' value_name  RWsupports interface_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_inheritance_spec/':' value_name  RWsupports interface_name")
		}
		$$, _ = newValueInheritanceSpec($4, $2)
	}
	|  ':' value_namePlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_inheritance_spec/':' value_namePlus")
		}
		$$, _ = newValueInheritanceSpec(nil, $2)
	}
	|  ':' value_namePlus RWsupports interface_namePlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_inheritance_spec/':' value_namePlus RWsupports interface_namePlus")
		}
		$$, _ = newValueInheritanceSpec($4, $2)
	}
	|  ':' RWtruncatable value_namePlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_inheritance_spec/':' RWtruncatable value_namePlus")
		}
		$$, _ = newValueInheritanceSpec(nil, $3)
	}
	|  ':' RWtruncatable value_namePlus  RWsupports interface_namePlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_inheritance_spec/':' RWtruncatable value_namePlus  RWsupports interface_namePlus")
		}
		$$, _ = newValueInheritanceSpec($5, $3)
	}

//(104)
value_namePlus:
	value_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_namePlus/value_name")
		}
		$$ = $1
	}
	| value_namePlus ',' value_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_namePlus/value_namePlus ',' value_name")
		}
		node := $1
		for node.GetNextNode() != nil {
			node = node.GetNextNode()
		}
		_ = node.SetNextNode($3)
	}
value_name:
	scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_name/scoped_name")
		}
		$$ = $1
	}
//(105)
value_elementPlus:
	value_element
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_elementPlus/value_element")
		}
		$$ = $1
	}
	|value_elementPlus value_element
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_elementPlus/value_elementPlus value_element")
		}
		node := $1
		for node.GetNextNode() != nil {
			node = node.GetNextNode()
		}
		_ = node.SetNextNode($2)
	}

value_elementStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_elementStar")
		}
		$$ = nil
	}
	|value_elementPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_elementStar/value_elementPlus")
		}
		$$ = $1
	}

value_element:
	export
	{

		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_element/export")
		}
		$$ = $1
	}
	| state_member
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
                          yaccLogger.Log("value_element/state_member")
		}
		$$ = $1
	}
	| init_dcl
	{
		
		$$ = $1
	}
//(106)
state_member:
	RWpublic type_spec declarators ';'
	{
		
	}
	| RWprivate  type_spec declarators ';'
	{
		
	}
//(107)
init_dcl:
	RWfactory Identifier '('  ')'  ';'
	{
		
		$$ = nil
	}
	| RWfactory Identifier '('  ')'  raises_expr  ';'
	{
		
		$$ = nil
	}
	| RWfactory Identifier '('  init_param_dcls  ')'  ';'
	{
		
		$$ = nil
	}
	| RWfactory Identifier '('  init_param_dcls  ')'  raises_expr  ';'
	{
		
		$$ = nil
	}
//(108)
init_param_dcls:
	init_param_dclPlus
	{
		
		$$ = nil
	}
//(109)
init_param_dclPlus:
	init_param_dcl
	{
		
		$$ = nil
	}
	|init_param_dclPlus ',' init_param_dcl
	{
		
		$$ = nil
	}
init_param_dcl:
	RWin type_spec simple_declarator
	{
		
		$$ = nil
	}
//(110)
value_forward_dcl:
	value_kind Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_forward_dcl/value_kind Identifier")
		}
		$$, _ = newValueHeader($1, $2, nil)

	}
//(113)
type_id_dcl: RWtypeid scoped_name String_literal {}
//(114)
type_prefix_dcl:
	RWtypeprefix scoped_name String_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("type_prefix_dcl:/RWtypeprefix scoped_name String_literal")
		}
		$$ = NewTypePrefix($1.LexemData(), $2.Identifier(), $3.Identifier())
	}
//(115)
import_dcl:
	RWimport imported_scope
	{
		$$ = newImportDcl($1.LexemData(), $2)
	}
//(116)
imported_scope:
	scoped_name
	{
		$$ = $1
	} | String_literal
	{
		$$ = $1
	}
//(118)
object_type:
	RWObject
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("object_type/RWObject")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlObjectType); ok {
			$$ = defaultTypes.ObjectType()
			if $$ == nil {
				YaccIdllex.Error("no object primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlObjectType interface")
			return 1
		}
	}
//(120)
op_oneway_dcl: RWoneway RWvoid Identifier '('  in_parameter_dclsStar  ')' {}
//(121)
in_parameter_dclsStar:
	 {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	 }
	| in_parameter_dcls
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| in_parameter_dclsStar in_parameter_dcls
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
in_parameter_dcls:
	in_param_dclPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(122)
in_param_dclPlus:
	in_param_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| in_param_dclPlus ',' in_param_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
in_param_dcl:
	RWin type_spec simple_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(123)
op_with_context:
 	 op_oneway_dcl context_expr
 	 {
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
 	 }
 	| op_dcl  context_expr
 	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
 	}
//(124)
String_literalPlus:
	String_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| String_literalPlus ',' String_literal
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
context_expr:
	RWcontext '(' String_literalPlus  ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(126)
value_box_def:
	value_kind Identifier type_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(132)
value_base_type:
	RWValueBase
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("value_base_type/RWValueBase")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlValueBaseType); ok {
			$$ = defaultTypes.ValueBaseType()
			if $$ == nil {
				YaccIdllex.Error("no valuebase primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlValueBaseType interface")
			return 1
		}
	}
//(134)
component_dcl:
	component_def
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| component_forward_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(135)
component_forward_dcl:
	RWcomponent Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(136)
component_def:
	component_header '{' component_body '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(137) //(154)
component_header:
	RWcomponent Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWcomponent Identifier  component_inheritance_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWcomponent Identifier  supported_interface_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWcomponent Identifier  component_inheritance_spec  supported_interface_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(138)
component_inheritance_spec:
	':' scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(139)
component_body:
	component_exportStar
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(140)//(156)//(179)
component_exportStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| component_export
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| component_exportStar component_export
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
component_export:
	provides_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| uses_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| attr_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| emits_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| publishes_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| consumes_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| port_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(141)
provides_dcl:
	RWprovides interface_type Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(142)//(157)
interface_type:
	scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWObject
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(143)//(158)
uses_dcl:
	RWuses interface_type Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|RWuses RWmultiple interface_type Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(145)
home_dcl:
	home_header '{' home_body '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(146)//(162)
home_header:
	RWhome Identifier  RWmanages scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier  home_inheritance_spec  RWmanages scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier  RWmanages scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier RWmanages scoped_name  primary_key_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier   supported_interface_spec  RWmanages scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier   supported_interface_spec  RWmanages scoped_name  primary_key_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier  home_inheritance_spec   RWmanages scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier  home_inheritance_spec   RWmanages scoped_name  primary_key_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier  home_inheritance_spec   supported_interface_spec  RWmanages scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWhome Identifier  home_inheritance_spec   supported_interface_spec  RWmanages scoped_name  primary_key_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}


//(147)
home_inheritance_spec:
	':' scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(148)
home_body:
	home_exportStar
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(149)//(164)
home_exportStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| home_export
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| home_exportStar home_export
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
home_export:
	export
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| factory_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| finder_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(150)
factory_dcl:
	RWfactory Identifier '('  ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|RWfactory Identifier '('  ')'  raises_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|RWfactory Identifier '('  factory_param_dcls  ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|RWfactory Identifier '('  factory_param_dcls  ')'  raises_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(151)
factory_param_dcls:
	factory_param_dclPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(152)
factory_param_dclPlus:
	factory_param_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| factory_param_dclPlus ',' factory_param_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
factory_param_dcl:
	RWin type_spec simple_declarator
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(155)
supported_interface_spec:
	RWsupports scoped_namePlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(159)
emits_dcl:
	RWemits scoped_name Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(160)
publishes_dcl:
	RWpublishes scoped_name Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(161)
consumes_dcl:
	RWconsumes scoped_name Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(163)
primary_key_spec:
	RWprimarykey scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(165)
finder_dcl:
	RWfinder Identifier '('  ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|RWfinder Identifier '('  ')'  raises_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|RWfinder Identifier '('  init_param_dcls  ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|RWfinder Identifier '('  init_param_dcls  ')'  raises_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(166)
event_dcl:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| event_def
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| event_abs_def
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| event_forward_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(167)
event_forward_dcl:
	RWeventtype Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|  RWabstract  RWeventtype Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(168)
event_abs_def:
	RWabstract RWeventtype Identifier  '{' exportStar '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWabstract RWeventtype Identifier  value_inheritance_spec  '{' exportStar '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(169)
event_def:
	event_header '{' value_elementStar '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(170)
event_header:
	RWeventtype Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|  RWeventtype Identifier  value_inheritance_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|  RWcustom  RWeventtype Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|  RWcustom  RWeventtype Identifier  value_inheritance_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(172)
porttype_dcl:
	porttype_def
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| porttype_forward_dcl
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(173)
porttype_forward_dcl:
	RWporttype Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(174)
porttype_def:
	RWporttype Identifier '{' port_body '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(175)
port_body: port_ref port_exportStar{}
//(176)
port_ref:
	provides_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| uses_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| port_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(177)
port_exportStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
port_export:
	port_ref
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| attr_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(178)
port_dcl:
 	RWmirrorport scoped_name Identifier
 	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
 	}
 	| RWport  scoped_name Identifier
 	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
 	}
//(180)
connector_dcl:
	connector_header '{' connector_exportPlus '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(181)
connector_header:
	RWconnector Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	|RWconnector Identifier  connector_inherit_spec
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(182)
connector_inherit_spec: ':'
	scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(183)
connector_exportPlus:{}
connector_export:
	port_ref
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| attr_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(185)
template_module_dcl:
	RWmodule Identifier '<' formal_parameters '>' '{' tpl_definitionPlus'}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(186)
formal_parameters:
	formal_parameterPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(187)
formal_parameterPlus:
	formal_parameter
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| formal_parameterPlus ',' formal_parameter
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
formal_parameter:
	formal_parameter_type Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(188)
formal_parameter_type:
	RWtypename
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWinterface
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWvaluetype
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWeventtype
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWstruct
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWunion
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWexception
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWenum
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWsequence
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWconst const_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| sequence_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(189)
tpl_definitionPlus:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
tpl_definition:
	definition
		{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
        		$$ = nil
        	}
	| template_module_ref ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(190)
template_module_inst:
	RWmodule scoped_name '<' actual_parameters '>' Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(191)
actual_parameters:
	actual_parameterPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(192)
actual_parameterPlus:
	actual_parameter
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| actual_parameterPlus ',' actual_parameter
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
actual_parameter:
 	type_spec
 	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
 		$$ = nil
 	}
 	| const_expr
 	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
 		$$ = nil
 	}
//(193)
template_module_ref: RWalias scoped_name '<' formal_parameter_names '>' Identifier{}
//(194)
IdentifierPlus: Identifier{}| IdentifierPlus ',' Identifier{}
formal_parameter_names: IdentifierPlus{}
//(199)
map_type:
 	RWmap '<' type_spec ',' type_spec ',' positive_int_const '>'
 	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
 	}
 	| RWmap '<' type_spec ',' type_spec '>'
 	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
 	}
//(200)
bitset_dcl:
	RWbitset Identifier  '{' bitfieldStar '}'
	{
		YaccIdllex.Error("implement bitset_dcl1")
		return 1
	}
	RWbitset Identifier ':' scoped_name '{' bitfieldStar '}'
	{
		YaccIdllex.Error("implement bitset_dcl2")
		return 1
	}
//(201)
IdentifierStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

bitfieldStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
bitfield:
	bitfield_spec IdentifierStar ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(202)
bitfield_spec:
	RWbitfield '<' positive_int_const '>'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| RWbitfield '<' positive_int_const ',' destination_type '>'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(203)
destination_type:
	boolean_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| octet_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| integer_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(204)
bitmask_dcl:
	RWbitmask Identifier '{' bit_valueStar '}'
	{
		YaccIdllex.Error("implement bitmask_dcl")
		return 1
	}
//(205)
bit_valueStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

bit_value: Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

//(208)
signed_tiny_int:
	RWint8
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("signed_tiny_int/RWint8")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlInt8Type); ok {
			$$ = defaultTypes.Int8Type()
			if $$ == nil {
				YaccIdllex.Error("no int8 primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlInt8Type interface")
			return 1
		}
	}
//(209)
unsigned_tiny_int:
	RWuint8
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("unsigned_tiny_int/RWuint8")
		}
		if defaultTypes, ok := YaccIdllex.(IIdlUInt8Type); ok {
			$$ = defaultTypes.UInt8Type()
			if $$ == nil {
				YaccIdllex.Error("no uint8 primitive")
				return 1
			}
		} else {
			YaccIdllex.Error("no IIdlUInt8Type interface")
			return 1
		}
	}
//(216)

//(219)
annotation_dcl:
	annotation_header '{' annotation_body '}'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(220)
annotation_header:
	Annotation Identifier
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(221)
annotation_bodyEntry:
	annotation_member
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| enum_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| const_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| typedef_dcl ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}

annotation_bodyEntryStar:
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
annotation_body:
	annotation_bodyEntryStar
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(222)
annotation_member:
	annotation_member_type simple_declarator  ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| annotation_member_type simple_declarator  RWdefault const_expr  ';'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(223)
annotation_member_type:
	const_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| any_const_type
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(224)
any_const_type:
	RWany
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(225)
annotation_appl:
	'@' scoped_name
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| '@' scoped_name  '(' annotation_appl_params ')'
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(226)
annotation_appl_params:
	const_expr
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
	| annotation_appl_paramPlus
	{
		if yaccLogger, ok := YaccIdllex.(IYaccLogger); ok {
			yaccLogger.Log("")
		}
		$$ = nil
	}
//(227)
annotation_appl_paramPlus:
	annotation_appl_param{}
	| annotation_appl_paramPlus ',' annotation_appl_param{}
annotation_appl_param: Identifier '=' const_expr{}
%%



