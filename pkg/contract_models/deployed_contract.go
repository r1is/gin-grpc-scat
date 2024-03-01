package models

type DeployedContract struct {
	Name                    string                    `json:"name"`
	ChainName               string                    `json:"chain_name"`
	Address                 string                    `json:"address"`
	ByteCode                string                    `json:"byte_code"`
	SourceCode              string                    `json:"source_code"`
	CreateAtBlockNumber     int64                     `json:"create_at_block_number"`
	CreateTxHash            string                    `json:"create_tx_hash"`
	CreateTxInput           string                    `json:"create_tx_input"`
	AddAt                   int64                     `json:"add_at"`
	ByteCodeAnalyzeResult   []ByteCodeAnalyzeResult   `json:"byte_code_analyze_result"`
	SourceCodeAnalyzeResult []SourceCodeAnalyzeResult `json:"source_code_analyze_result"`
}
