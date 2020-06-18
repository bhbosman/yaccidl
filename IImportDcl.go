package yaccidl

type IImportDcl interface {
	IYaccNode
	IsImportDcl() bool
	ImportedScope() IYaccNode
}

type importDcl struct {
	YaccNode
	importedScope IYaccNode
}

func (i importDcl) ImportedScope() IYaccNode {
	return i.importedScope
}

func (i importDcl) IsImportDcl() bool {
	return true
}

func newImportDcl(lexemData *LexemValue, importedScope IYaccNode) IImportDcl {
	return &importDcl{
		YaccNode:      *NewYaccNode(importedScope.Identifier(), lexemData),
		importedScope: importedScope,
	}
}
