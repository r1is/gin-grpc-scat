package tools

import (
	models "pkg/contract_models"
)

type Oyente struct{}

var ToolOyente Tool = Oyente{}

func (o Oyente) Name() string {
	return models.ToolOyente.Name
}

func (o Oyente) ShortDescription() string {
	return models.ToolOyente.ShortDescription
}

func (o Oyente) LongDescription() string {
	return models.ToolOyente.LongDescription
}

func (o Oyente) Version() string {
	return models.ToolOyente.Version
}

func (o Oyente) Reference() string {
	return models.ToolOyente.Reference
}

func (o Oyente) Install() error {
	return nil
}

func (o Oyente) AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error) {
	return models.ByteCodeAnalyzeResult{}, nil
}

func (o Oyente) AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error) {
	return models.SourceCodeAnalyzeResult{}, nil
}
