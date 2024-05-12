package controller

import (
	"k8s-go-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Pod pod

type pod struct{}

func (p *pod) GetPods(ctx *gin.Context) {
	// handler parameter
	// 匿名結構體，用於定義入參，get請求為 form 格式，其他請求為 json 格式
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
		Limit      int    `form:"limit"`
		Page       int    `form:"page"`
	})

	// form 格式使用 Bind 方法，json 格式使用 ShouldBindJSON 方法
	if err := ctx.Bind(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
	}
	data, err := service.Pod.GetPods(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Error("get pods error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get pods error" + err.Error(),
			"data":    nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
