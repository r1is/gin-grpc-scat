package mythril

import (
	"context"
	"fmt"
	"pb"
	"pkg/log"

	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn

// 新建连接
func InitMythrilClient() {

	gRpcMythrilAddr := fmt.Sprintf("%s:%d", "Mythril", 50051)
	conn, err := grpc.Dial(gRpcMythrilAddr, grpc.WithInsecure())
	if err != nil {
		log.LogError(fmt.Sprintf("did not connect: %v", err))
	}
	Conn = conn
}

// 关闭连接,记得defer关闭
func CloseMythrilClient(conn *grpc.ClientConn) {
	conn.Close()
}

func MythrilDect(_sourceCode string) (string, error) {
	ctx := context.Background()
	client := pb.NewAnalyzeToolsClient(Conn)
	req := &pb.SourceCodeRequest{SourceCode: _sourceCode}
	res, err := client.CheckSmartContractCode(ctx, req)
	if err != nil {
		log.LogError(fmt.Sprintf("error: %v", err))
	}
	log.LogInfo(fmt.Sprintf("Mythril Dect result: %v", res.Reslut))

	return res.Reslut, nil
}

func GetMythrilClient() *grpc.ClientConn {
	return Conn
}
