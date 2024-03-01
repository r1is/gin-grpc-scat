package controllers

import (
	"analyze_tool/tools"
	"fmt"
	"net/http"
	"pkg/apis"
	"pkg/log"

	"github.com/gin-gonic/gin"

	"contract_analyzer/db"
)

type AnalyzToolController struct{}

// AnalyzeSourceCode 分析源代码controllers
func (a AnalyzToolController) AnalyzeSourceCode(ctx *gin.Context) {
	var req apis.AnalyzeSourceCodeReq
	resp := apis.AnalyzeSourceCodeResp{}

	// 解析请求中的json数据
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.StatusCode = -1
		resp.StatusMessage = fmt.Sprintf("Illegal params: %v", err)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	fmt.Println("req:", req, "req.ToolName: ", req.ToolName)
	tool := tools.GetToolByName(req.ToolName)
	if tool == nil || tool.Name() == "" {
		resp.StatusCode = -1
		var toolsName []string
		for _, tool := range tools.Tools {
			toolsName = append(toolsName, tool.Name())
		}
		resp.StatusMessage = fmt.Sprintf("The tool name is invalid, and the available tools are: %v. You can also use 'all' to specify all tools", toolsName)
		ctx.JSON(http.StatusOK, resp)
		return
	}

	ctx.JSON(http.StatusOK, apis.Resp{StatusCode: 0, StatusMessage: "ok"})
	go asyncCodeVulScan(req)

}

func asyncCodeVulScan(req apis.AnalyzeSourceCodeReq) {
	//异步扫描代码漏洞
	resp := &apis.AnalyzeSourceCodeResp{RequestId: req.RequestId}
	tool := tools.GetToolByName(req.ToolName)

	result, err := tool.AnalyzeSourceCode(req.SourceCode)

	if err != nil {
		resp.StatusCode = -1
		resp.StatusMessage = fmt.Sprintf("Tool %v run error: %v", tool.Name(), err)
		_err := db.InsertData(resp)
		if _err != nil {
			log.LogInfo(fmt.Sprintf("InsertData error: %v", _err))
		}
		log.LogInfo(fmt.Sprintf("Resp: %v", resp))
		return
	}

	resp.Resp = apis.SuccessResp()
	resp.Result = result
	//将结果存入mongodb
	_err := db.InsertData(resp)
	if _err != nil {
		log.LogInfo(fmt.Sprintf("InsertData error: %v", _err))
	}
	log.LogInfo(fmt.Sprintf("Req: %v, Resp: %v", req, resp))
}
