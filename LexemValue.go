package yaccidl

import (
	"fmt"
	"github.com/bhbosman/yaccpragma"
)

type LexemValue struct {
	Row              int     `json:"row"`
	Col              int     `json:"col"`
	TokenName        string  `json:"token_name"`
	TypeKind         int     `json:"type_kind"`
	ValidToken       bool    `json:"valid_token"`
	FloatValue       float64 `json:"float_value,omitempty"`
	IntValue         int64   `json:"int_value,omitempty"`
	StringValue      string  `json:"string_value,omitempty"`
	Eof              bool    `json:"eof,omitempty"`
	SourceFolderName string  `json:"folder_name,omitempty"`
	SourceFileName   string  `json:"file_name"`
	SourceFolderId   string  `json:"folder_id,omitempty"`
	//TargetFileName   string                 `json:"target_file_name"`
	PragmaNodeValue yaccpragma.IPragmaNode `json:"pragma_node_value"`
	//CheckTarget      bool                   `json:"check_target"`
}

//func (l LexemValue) GetCheckTarget() bool {
//	return l.CheckTarget
//}

//func (l LexemValue) GetTargetFolderName() string {
//	return l.TargetFileName
//}

func (l LexemValue) GetRow() int {
	return l.Row
}

func (l LexemValue) GetCol() int {
	return l.Col
}

func (l LexemValue) GetSourceFolderId() string {
	return l.SourceFolderId
}

func (l LexemValue) String() string {
	return fmt.Sprintf("LexemValue: {SourceFileName: %v, Row: %v, Col: %v}", l.SourceFileName, l.Row, l.Col)
}

func (l LexemValue) GetSourceFileName() string {
	return l.SourceFileName
}

func (l LexemValue) GetSourceFolderName() string {
	return l.SourceFolderName
}

type LexemCollection struct {
	LexemValues []LexemValue
}

func NewLexemEofValue(currentContext CurrentContext) (*LexemValue, error) {
	return &LexemValue{
		Row:              currentContext.Row,
		Col:              currentContext.Col,
		TokenName:        "",
		TypeKind:         0,
		ValidToken:       false,
		FloatValue:       0,
		IntValue:         0,
		StringValue:      "",
		Eof:              true,
		SourceFolderName: currentContext.FileInformationEntry.Path,
		SourceFileName:   currentContext.FileName,
		SourceFolderId:   currentContext.FileInformationEntry.FolderId,
		//TargetFileName:   "",
		PragmaNodeValue: nil,
		//CheckTarget:      false,
	}, nil
}

func NewLexemNoValue(typeKind int, tokenName func(int) string, currentContext CurrentContext, validToken bool) (*LexemValue, error) {
	var v string = ""
	if tokenName != nil {
		v = tokenName(typeKind)
	}
	return &LexemValue{
		Row:              currentContext.Row,
		Col:              currentContext.Col,
		TokenName:        v,
		TypeKind:         typeKind,
		ValidToken:       validToken,
		FloatValue:       0,
		IntValue:         0,
		StringValue:      "",
		Eof:              false,
		SourceFolderName: currentContext.FileInformationEntry.Path,
		SourceFileName:   currentContext.FileName,
		SourceFolderId:   currentContext.FileInformationEntry.FolderId,
		//TargetFileName:   "",
		PragmaNodeValue: nil,
		//CheckTarget:      false,
	}, nil
}

func NewLexemStringValue(typeKind int, tokenName func(int) string, stringValue string, currentContext CurrentContext, validToken bool) (*LexemValue, error) {
	var v string = ""
	if tokenName != nil {
		v = tokenName(typeKind)
	}
	return &LexemValue{
		Row:              currentContext.Row,
		Col:              currentContext.Col,
		TokenName:        v,
		TypeKind:         typeKind,
		ValidToken:       validToken,
		FloatValue:       0,
		IntValue:         0,
		StringValue:      stringValue,
		Eof:              false,
		SourceFolderName: currentContext.FileInformationEntry.Path,
		SourceFileName:   currentContext.FileName,
		SourceFolderId:   currentContext.FileInformationEntry.FolderId,
		//TargetFileName:   "",
		PragmaNodeValue: nil,
		//CheckTarget:      false,
	}, nil
}

func NewLexemIntValue(typeKind int, tokenName func(int) string, intValue int64, currentContext CurrentContext, validToken bool) (*LexemValue, error) {
	var v string = ""
	if tokenName != nil {
		v = tokenName(typeKind)
	}
	return &LexemValue{
		Row:              currentContext.Row,
		Col:              currentContext.Col,
		TokenName:        v,
		TypeKind:         typeKind,
		ValidToken:       validToken,
		FloatValue:       0,
		IntValue:         intValue,
		StringValue:      "",
		Eof:              false,
		SourceFolderName: currentContext.FileInformationEntry.Path,
		SourceFileName:   currentContext.FileName,
		SourceFolderId:   currentContext.FileInformationEntry.FolderId,
		//TargetFileName:   "",
		PragmaNodeValue: nil,
		//CheckTarget:      false,
	}, nil
}

func NewLexemFloatValue(typeKind int, tokenName func(int) string, floatValue float64, currentContext CurrentContext, validToken bool) (*LexemValue, error) {
	var v string = ""
	if tokenName != nil {
		v = tokenName(typeKind)
	}
	return &LexemValue{
		Row:              currentContext.Row,
		Col:              currentContext.Col,
		TokenName:        v,
		TypeKind:         typeKind,
		ValidToken:       validToken,
		FloatValue:       floatValue,
		IntValue:         0,
		StringValue:      "",
		Eof:              false,
		SourceFolderName: currentContext.FileInformationEntry.Path,
		SourceFileName:   currentContext.FileName,
		SourceFolderId:   currentContext.FileInformationEntry.FolderId,
		//TargetFileName:   "",
		PragmaNodeValue: nil,
		//CheckTarget:      false,
	}, nil
}

func NewLexemPragmaNodeValue(typeKind int, tokenName func(int) string, value yaccpragma.IPragmaNode, currentContext CurrentContext, validToken bool) (*LexemValue, error) {
	var v string = ""
	if tokenName != nil {
		v = tokenName(typeKind)
	}
	return &LexemValue{
		Row:              currentContext.Row,
		Col:              currentContext.Col,
		TokenName:        v,
		TypeKind:         typeKind,
		ValidToken:       validToken,
		FloatValue:       0,
		IntValue:         0,
		StringValue:      "",
		Eof:              false,
		SourceFolderName: currentContext.FileInformationEntry.Path,
		SourceFileName:   currentContext.FileName,
		SourceFolderId:   currentContext.FileInformationEntry.FolderId,
		//TargetFileName:   "",
		PragmaNodeValue: value,
		//CheckTarget:      false,
	}, nil
}
