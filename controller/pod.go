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
// 表單參數統一用 form，ctx方法要用bind方法，如果是json格式請用ShouldBindJSON方法
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

// delete
func (p *pod) DeletePod(ctx *gin.Context) {
	params := new(struct {
		PodName   string `json:"pod_name"`
		Namespace string `json:"namespace"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Pod.DeletePod(params.PodName, params.Namespace)
	if err != nil {
		logger.Error("delete pod error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete pod error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete pod success",
		"data":    nil,
	})
}

// update
func (p *pod) UpdatePod(ctx *gin.Context) {
	params := new(struct {
		PodName   string `json:"pod_name"`
		Namespace string `json:"namespace"`
		Content   string `json:"content"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Pod.UpdatePod(params.PodName, params.Namespace, params.Content)
	if err != nil {
		logger.Error("update pod error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "update pod error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update pod success",
		"data":    nil,
	})
}

// get pod container
func (p *pod) GetPodContainers(ctx *gin.Context) {
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
	data, err := service.Pod.GetPodContainers(params.PodName, params.Namespace)
	if err != nil {
		logger.Error("get pod containers error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get pod containers error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get pod containers success",
		"data":    data,
	})
}

// get pod log
func (p *pod) GetPodLog(ctx *gin.Context) {
	params := new(struct {
		PodName   string `form:"pod_name"`
		Container string `form:"container"`
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
	data, err := service.Pod.GetPodLog(params.PodName, params.Container, params.Namespace)
	if err != nil {
		logger.Error("get pod log error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get pod log error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get pod log success",
		"data":    data,
	})
}

func (p *pod) GetPodNumPerNp(ctx *gin.Context) {
	data, err := service.Pod.GetPodNumPerNp()
	if err != nil {
		logger.Error("get pod number per namespace error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get pod number per namespace error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get pod number per namespace success",
		"data":    data,
	})
}
