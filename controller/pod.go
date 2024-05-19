package controller

import (
	"k8s-go-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Pod pod

type pod struct{}

// Controller中的方法入參是gin.Context，用於從上下文中獲取請求參數及定義響應內容
//流程：綁定參數 -> 調用 service 代碼 -> 根據調用結果響應具體內容

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
		return
	}
	data, err := service.Pod.GetPods(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Error("get pods error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get pods error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

// get pod detail
func (p *pod) GetPodDetail(ctx *gin.Context) {
	params := new(struct {
		PodName   string `form:"pod_name"`
		Namespace string `form:"namespace"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Pod.GetPodDetail(params.PodName, params.Namespace)
	if err != nil {
		logger.Error("get pod detail error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get pod detail error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
