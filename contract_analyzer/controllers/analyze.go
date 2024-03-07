package controllers

import (
	"analyze_tool/tools"
	"contract_analyzer/utils/toolgrpc"
	"encoding/json"
	"fmt"
	"net/http"
	"pkg/apis"
	models "pkg/contract_models"
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

	// fmt.Println("req:", req, "req.ToolName: ", req.ToolName)
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
	if db.IsExistRequestId(req.RequestId) {
		ctx.JSON(http.StatusOK, apis.SuccessRespWithMsg("任务已存在!"))
	} else {
		ctx.JSON(http.StatusOK, apis.SuccessRespOK())
		go asyncCodeVulScan(req)
	}

}

func asyncCodeVulScan(req apis.AnalyzeSourceCodeReq) {
	//异步扫描代码漏洞
	resp := &apis.AnalyzeSourceCodeResp{RequestId: req.RequestId}
	tool := toolgrpc.GetToolByName(req.ToolName)
	// tool := tools.GetToolByName(req.ToolName)
	resp.Resp = apis.AuditingResp()
	//第一次操作数据库，写入正在审计中的状态
	db.InsertData(resp)

	_result, err := tool.Dect(req.SourceCode)

	if err != nil {
		// resp.StatusCode = -1
		// resp.StatusMessage = fmt.Sprintf("Tool %v run error: %v", tool.Name(), err)
		resp.Resp = apis.FinishAuditFailResp(fmt.Sprintf("Tool %v run error: %v", tool.Name(), err))
		_err := db.UpdateStatusByReqId(req.RequestId, resp.Resp)
		if _err != nil {
			log.LogInfo(fmt.Sprintf("InsertData error: %v", _err))
		}
		log.LogInfo(fmt.Sprintf("审计失败的Resp: %v", resp))
		return
	}
	log.LogInfo(fmt.Sprintf("Tool %v run result: %v", tool.Name(), _result))
	//以下是 grpc 条用正常，且审计正常的逻辑

	resp.Resp = apis.FinishAuditResp()
	_err := db.UpdateStatusByReqId(req.RequestId, resp.Resp)
	if _err != nil {
		log.LogInfo(fmt.Sprintf("InsertData error: %v", _err))
	}

	var aduitRes models.SourceCodeAnalyzeResult
	log.LogInfo("\ngrpc返回的扫描结果：" + _result + "\n")
	_err_ := json.Unmarshal([]byte(_result), &aduitRes)
	if _err_ != nil {
		fmt.Println("转换为JSON时出错:", err)
		return
	}

	resp.Result = aduitRes
	//将结果存入mongodb
	_err1 := db.UpdateVluResByReqId(req.RequestId, resp.Result)
	if _err1 != nil {
		log.LogInfo(fmt.Sprintf("InsertData error: %v", _err1))
	}
	log.LogInfo(fmt.Sprintf("Req: %v, Resp: %v", req, resp))
}
