package toolgrpc

import (
	"context"
	"fmt"
	"pb"
	"pkg/log"

	"google.golang.org/grpc"
)

type ManticoreToolGrpc struct {
	conn   *grpc.ClientConn
	client pb.AnalyzeToolsClient
}

var mManticoreToolGrpc ToolGrpc = &ManticoreToolGrpc{}

func (m *ManticoreToolGrpc) Name() string {
	return "Manticore"
}

func (m *ManticoreToolGrpc) InitClient(port string) {
	gpcAdress := m.Name() + ":" + port
	conn, err := grpc.Dial(gpcAdress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	m.conn = conn
	m.client = pb.NewAnalyzeToolsClient(conn)
}

func (m *ManticoreToolGrpc) CloseClient() {
	m.conn.Close()
}

func (m *ManticoreToolGrpc) Dect(_sourceCode string) (string, error) {
	ctx := context.Background()
	req := &pb.SourceCodeRequest{SourceCode: _sourceCode}
	res, err := m.client.CheckSmartContractCode(ctx, req)
	if err != nil {
		log.LogError(fmt.Sprintf("error: %v", err))
	}
	log.LogInfo(fmt.Sprintf("Mythril Dect result: %v", res.Reslut))
	return res.Reslut, nil
}

func (m *ManticoreToolGrpc) GetClient() *grpc.ClientConn {
	return m.conn
}
