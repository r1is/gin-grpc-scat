package tools

import (
	models "pkg/contract_models"
)

type Conkas struct{}

var ToolConkas Tool = Conkas{}

func (c Conkas) Name() string {
	return models.ToolConkas.Name
}

func (c Conkas) ShortDescription() string {
	return models.ToolConkas.ShortDescription
}

func (c Conkas) LongDescription() string {
	return models.ToolConkas.LongDescription
}

func (c Conkas) Version() string {
	return models.ToolConkas.Version
}

func (c Conkas) Reference() string {
	return models.ToolConkas.Reference
}

func (c Conkas) AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error) {
	return models.ByteCodeAnalyzeResult{}, nil
}

func (c Conkas) AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error) {
	return models.SourceCodeAnalyzeResult{}, nil
}
