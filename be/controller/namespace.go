package controller

import (
	"k8s-go-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Namespace namespace

type namespace struct{}

func (ns *namespace) GetNamespaces(ctx *gin.Context) {
	params := new(struct {
		FilterName string `form:"filter_name"`
		Limit      int    `form:"limit"`
		Page       int    `form:"page"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}

	data, err := service.Namespace.GetNamespaces(params.FilterName, params.Limit, params.Page)
	if err != nil {
		logger.Error("get namespaces error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get namespaces error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (ns *namespace) GetNamespaceDetail(ctx *gin.Context) {
	params := new(struct {
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
	data, err := service.Namespace.GetNamespaceDetail(params.Namespace)
	if err != nil {
		logger.Error("get namespace detail error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get namespace detail error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (ns *namespace) DeleteNamespace(ctx *gin.Context) {
	params := new(struct {
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
	err := service.Namespace.DeleteNamespace(params.Namespace)
	if err != nil {
		logger.Error("delete namespace error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete namespace error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete namespace success",
		"data":    nil,
	})
}
