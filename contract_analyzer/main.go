package main

import (
	"contract_analyzer/db"
	"contract_analyzer/utils/mythril"
)

var port, mythrilPort, manticorePort, slitherPort int

func main() {
	// 初始化mongodb连接
	db.InitMongoDB()
	//程序退出前关闭连接
	defer db.CloseDB()

	mythril.InitMythrilClient()
	defer mythril.CloseMythrilClient(mythril.Conn)
}
