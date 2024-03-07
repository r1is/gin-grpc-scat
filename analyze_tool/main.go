package main

import (
	"analyze_tool/tools"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"pb"
	"pkg/log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedAnalyzeToolsServer
}

func (s *server) CheckSmartContractCode(ctx context.Context, in *pb.SourceCodeRequest) (*pb.CheckResult, error) {
	//获取智能合约代码
	SourceCode := in.SourceCode
	tool := tools.GetToolByName(toolName)
	res, err := tool.AnalyzeSourceCode(SourceCode)
	if err != nil {
		return &pb.CheckResult{Reslut: fmt.Sprintf("调用%v发生错误", toolName)}, err
	}
	// 将结构体转换为 JSON 格式的字符串
	jsonData, _ := json.Marshal(res)
	jsonString := string(jsonData)
	return &pb.CheckResult{Reslut: jsonString}, nil
}

var port string
var toolName string

func main() {
	flag.StringVar(&port, "p", "0", "web port")
	flag.StringVar(&toolName, "t", "", "tool name")
	flag.Parse()

	addr := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.LogError(fmt.Sprintf("failed to listen: %v", err))

	}
	log.LogInfo(fmt.Sprintf("server is running on port %s", addr))
	s := grpc.NewServer()
	pb.RegisterAnalyzeToolsServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.LogError(fmt.Sprintf("failed to serve: %v", err))
	}

}
