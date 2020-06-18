package yaccidl

func NewLexemDataForPrimitives(
	sourceFilename string,
	sourceFolderName string,
	sourceUniqueFolderName string,
	row, col int) *LexemValue {
	return &LexemValue{
		Row:              row,
		Col:              col,
		TokenName:        "",
		TypeKind:         0,
		ValidToken:       false,
		FloatValue:       0,
		IntValue:         0,
		StringValue:      "",
		Eof:              false,
		SourceFolderName: sourceFolderName,
		SourceFileName:   sourceFilename,
		SourceFolderId:   sourceUniqueFolderName,
		//TargetFileName:   targetFolderName,
		PragmaNodeValue: nil,
		//CheckTarget:      true,
	}
}
