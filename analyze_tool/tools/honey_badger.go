package tools

import (
	models "pkg/contract_models"
)

type HoneyBadger struct{}

var ToolHoneyBadger Tool = HoneyBadger{}

func (h HoneyBadger) Name() string {
	return models.ToolHoneyBadger.Name
}

func (h HoneyBadger) ShortDescription() string {
	return models.ToolHoneyBadger.ShortDescription
}

func (h HoneyBadger) LongDescription() string {
	return models.ToolHoneyBadger.LongDescription
}

func (h HoneyBadger) Version() string {
	return models.ToolHoneyBadger.Version
}

func (h HoneyBadger) Reference() string {
	return models.ToolHoneyBadger.Reference
}

func (h HoneyBadger) AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error) {
	return models.ByteCodeAnalyzeResult{}, nil
}

func (h HoneyBadger) AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error) {
	return models.SourceCodeAnalyzeResult{}, nil
}
