package tools

import (
	models "pkg/contract_models"
)

type Securify2 struct{}

var ToolSecurify2 Tool = Securify2{}

func (s Securify2) Name() string {
	return models.ToolSecurify2.Name
}

func (s Securify2) ShortDescription() string {
	return models.ToolSecurify2.ShortDescription
}

func (s Securify2) LongDescription() string {
	return models.ToolSecurify2.LongDescription
}

func (s Securify2) Version() string {
	return models.ToolSecurify2.Version
}

func (s Securify2) Reference() string {
	return models.ToolSecurify2.Reference
}

func (s Securify2) Install() error {
	return nil
}

func (s Securify2) AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error) {
	return models.ByteCodeAnalyzeResult{}, nil
}

func (s Securify2) AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error) {
	return models.SourceCodeAnalyzeResult{}, nil
}
