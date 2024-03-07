package toolgrpc

import (
	"context"
	"fmt"
	"pb"
	"pkg/log"

	"google.golang.org/grpc"
)

type SlitherToolGrpc struct {
	conn   *grpc.ClientConn
	client pb.AnalyzeToolsClient
}

var mSlitherToolGrpc ToolGrpc = &SlitherToolGrpc{}

func (m *SlitherToolGrpc) Name() string {
	return "Slither"
}

func (m *SlitherToolGrpc) InitClient(port string) {
	gpcAdress := m.Name() + ":" + port
	conn, err := grpc.Dial(gpcAdress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	m.conn = conn
	m.client = pb.NewAnalyzeToolsClient(conn)
}

func (m *SlitherToolGrpc) CloseClient() {
	m.conn.Close()
}

func (m *SlitherToolGrpc) Dect(_sourceCode string) (string, error) {
	ctx := context.Background()
	req := &pb.SourceCodeRequest{SourceCode: _sourceCode}
	res, err := m.client.CheckSmartContractCode(ctx, req)
	if err != nil {
		log.LogError(fmt.Sprintf("call %s error ->: %v", m.Name(), err))
		//fix bug
		return "", err
	}
	log.LogInfo(fmt.Sprintf("%s Dect result: %v", m.Name(), res.Reslut))
	return res.Reslut, nil
}

func (m *SlitherToolGrpc) GetClient() *grpc.ClientConn {
	return m.conn
}
