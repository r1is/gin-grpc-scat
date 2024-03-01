package toolgrpc

import "google.golang.org/grpc"

type ToolGrpc interface {
	Name() string
	InitClient(port string)
	CloseClient()
	Dect(_sourceCode string) (string, error)
	GetClient() *grpc.ClientConn
}

var ToolGrpcs []ToolGrpc = []ToolGrpc{mMythrilToolGrpc, mSlitherToolGrpc, mManticoreToolGrpc}
var ToolGrpcsMap map[string]ToolGrpc

func init() {
	ToolGrpcsMap = make(map[string]ToolGrpc)
	for _, tool := range ToolGrpcs {
		ToolGrpcsMap[tool.Name()] = tool
	}
}
func GetToolByName(name string) ToolGrpc {
	return ToolGrpcsMap[name]
}
