package controller

import (
	"k8s-go-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Deployment deployment

type deployment struct {
}

func (d *deployment) GetDeployments(ctx *gin.Context){
	params := new(struct {
		FilterName string `form:"filter_name"`
		Namespace  string `form:"namespace"`
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

	data, err := service.Deployment.GetDeployments(params.FilterName, params.Namespace, params.Limit, params.Page)
	if err != nil {
		logger.Error("get deployments error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get deployments error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}