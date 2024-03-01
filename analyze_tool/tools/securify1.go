package tools

import (
	models "pkg/contract_models"
)

type Securify1 struct{}

var ToolSecurify1 Tool = Securify1{}

func (s Securify1) Name() string {
	return models.ToolSecurify1.Name
}

func (s Securify1) ShortDescription() string {
	return models.ToolSecurify1.ShortDescription
}

func (s Securify1) LongDescription() string {
	return models.ToolSecurify1.LongDescription
}

func (s Securify1) Version() string {
	return models.ToolSecurify1.Version
}

func (s Securify1) Reference() string {
	return models.ToolSecurify1.Reference
}

func (s Securify1) Install() error {
	return nil
}

func (s Securify1) AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error) {
	return models.ByteCodeAnalyzeResult{}, nil
}

func (s Securify1) AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error) {
	return models.SourceCodeAnalyzeResult{}, nil
}
