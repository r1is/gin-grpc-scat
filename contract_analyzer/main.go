package main

import (
	"contract_analyzer/db"
	"contract_analyzer/routes"
	"contract_analyzer/utils/toolgrpc"
)

var port, mythrilPort, manticorePort, slitherPort int

var InsToolGrpc = map[string]toolgrpc.ToolGrpc{}

func main() {
	initConfig()
	r := routes.Router()
	r.Run(":8080")
}

func initConfig() {
	// 初始化mongodb连接
	db.InitMongoDB()
	//程序退出前关闭连接
	defer db.CloseDB()

	// 初始化grpc连接
	for _, tool := range toolgrpc.ToolGrpcs {
		tool.InitClient("50051")
		InsToolGrpc[tool.Name()] = tool
	}
	//程序退出前关闭连接
	defer func() {
		for _, tool := range InsToolGrpc {
			tool.CloseClient()
		}
	}()

}
