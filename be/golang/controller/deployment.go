package controller

import (
	"fmt"
	"k8s-go-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

var Deployment deployment

type deployment struct {
}

func (d *deployment) GetDeployments(ctx *gin.Context) {
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

func (d *deployment) GetDeploymentDetail(ctx *gin.Context) {
	params := new(struct {
		DeploymentName string `form:"deployment_name"`
		Namespace      string `form:"namespace"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Deployment.GetDeploymentDetail(params.DeploymentName, params.Namespace)
	if err != nil {
		logger.Error("get deployment detail error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get deployment detail error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (d *deployment) ScaleDeployment(ctx *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
		ScaleNum       int    `json:"scale_num"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}
	data, err := service.Deployment.ScaleDeployment(params.DeploymentName, params.Namespace, params.ScaleNum)
	if err != nil {
		logger.Error("scale deployment error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "scale deployment error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    fmt.Sprintf("scale deployment %d success", data),
	})
}

func (d *deployment) CreateDeployment(ctx *gin.Context) {
	var (
		deployCreate = new(service.DeployCreate)
		err          = ctx.ShouldBindJSON(deployCreate)
	)
	if err = ctx.ShouldBindJSON(deployCreate); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}
	if err = service.Deployment.CreateDeployment(deployCreate); err != nil {
		logger.Error("create deployment error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "create deployment error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "create deployment success",
	})
}

func (d *deployment) DeleteDeployment(ctx *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Deployment.DeleteDeployment(params.DeploymentName, params.Namespace)
	if err != nil {
		logger.Error("delete deployment error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "delete deployment error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete deployment success",
		"data":    nil,
	})
}

func (d *deployment) RestartDeployment(ctx *gin.Context) {
	params := new(struct {
		DeploymentName string `json:"deployment_name"`
		Namespace      string `json:"namespace"`
	})
	if err := ctx.Bind(params); err != nil {
		logger.Error("bind params error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "bind params error" + err.Error(),
			"data":    nil,
		})
		return
	}
	err := service.Deployment.RestartDeployment(params.DeploymentName, params.Namespace)
	if err != nil {
		logger.Error("restart deployment error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "restart deployment error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "restart deployment success",
		"data":    nil,
	})
}
