package routes

import (
	"contract_analyzer/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.POST("/pushCode", controllers.AnalyzToolController{}.AnalyzeSourceCode)
	// r.GET("/get_reslut_by_id",analyzeByteCode)

	r.GET("/getReport", controllers.QueryResController{}.QueryResByReqId)

	return r
}
