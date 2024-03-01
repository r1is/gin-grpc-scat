package controllers

import (
	"fmt"
	"net/http"
	"pkg/apis"

	"contract_analyzer/db"

	"github.com/gin-gonic/gin"
)

type QueryResController struct{}

func (query QueryResController) QueryResByReqId(ctx *gin.Context) {
	resqId, exists := ctx.GetQuery("requestId")
	if !exists {
		ctx.JSON(http.StatusOK, apis.Resp{StatusCode: -1, StatusMessage: "requestId is required"})
		return
	}

	result, err := db.QueryDataByReqId(resqId)
	if err != nil {
		ctx.JSON(http.StatusOK, apis.Resp{StatusCode: -2, StatusMessage: fmt.Sprintf("query error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
