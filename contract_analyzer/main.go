package main

import (
	"contract_analyzer/db"
	"contract_analyzer/routes"
	"contract_analyzer/utils/toolgrpc"
	"flag"
	models "pkg/contract_models"
	"strconv"
)

func main() {
	var port, mythrilPort, manticorePort, slitherPort int
	flag.IntVar(&port, "p", 0, "web port")
	flag.IntVar(&mythrilPort, "mythril_port", 0, "mythril port")
	flag.IntVar(&manticorePort, "manticore_port", 0, "manticore port")
	flag.IntVar(&slitherPort, "slither_port", 0, "slither port")
	flag.Parse()
	if port == 0 || mythrilPort == 0 || manticorePort == 0 || slitherPort == 0 {
		panic("web port, mythril port, manticore port and slither port are required")
	}
	ToolPorts := make(map[string]int)
	ToolPorts[models.ToolMythril.Name] = mythrilPort
	ToolPorts[models.ToolManticore.Name] = manticorePort
	ToolPorts[models.ToolSlither.Name] = slitherPort

	// 初始化mongodb连接
	db.InitMongoDB()
	//程序退出前关闭连接
	defer db.CloseDB()

	// 初始化grpc连接
	for _, tool := range toolgrpc.ToolGrpcs {
		port := ToolPorts[tool.Name()] //int
		tool.InitClient(strconv.Itoa(port))
	}
	//程序退出前关闭连接
	defer func() {
		for _, tool := range toolgrpc.ToolGrpcs {
			tool.CloseClient()
		}
	}()

	r := routes.Router()
	//监听地址
	r.Run(":" + strconv.Itoa(port))
}
