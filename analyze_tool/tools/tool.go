package tools

import (
	"fmt"
	models "pkg/contract_models"
)

type Tool interface {
	Name() string
	ShortDescription() string
	LongDescription() string
	Version() string
	Reference() string
	// AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error)
	AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error)
}

var Tools []Tool = []Tool{ToolMythril, ToolManticore, ToolOyente, ToolHoneyBadger, ToolSecurify1, ToolSecurify2, ToolSlither, ToolConkas}
var ToolsMap map[string]Tool

func init() {
	ToolsMap = make(map[string]Tool)
	for _, tool := range Tools {
		ToolsMap[tool.Name()] = tool
	}
}

func GetToolByName(name string) Tool {
	return ToolsMap[name]
}

func ToolDoc(tool Tool) string {
	return fmt.Sprintf("Tool %v\nName: %v\nShortDescription: %v\nLongDescription: %v\n",
		tool.Name(), tool.Name(), tool.ShortDescription(), tool.LongDescription())
}
