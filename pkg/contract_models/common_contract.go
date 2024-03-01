package models

type CommonContract struct {
	Name                    string                    `json:"name"`
	ByteCode                string                    `json:"byte_code"`
	SourceCode              string                    `json:"source_code"`
	AddAt                   int64                     `json:"add_at"`
	ByteCodeAnalyzeResult   []ByteCodeAnalyzeResult   `json:"byte_code_analyze_result"`
	SourceCodeAnalyzeResult []SourceCodeAnalyzeResult `json:"source_code_analyze_result"`
}
