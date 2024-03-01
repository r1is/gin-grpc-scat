package routes

import (
	"contract_analyzer/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.POST("/analyze_source_code", controllers.AnalyzToolController{}.AnalyzeSourceCode)
	// r.GET("/get_reslut_by_id",analyzeByteCode)

	r.GET("/query_res_by_req_id", controllers.QueryResController{}.QueryResByReqId)

	return r
}
