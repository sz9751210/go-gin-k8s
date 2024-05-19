package controller

import (
	"github.com/gin-gonic/gin"
	// "net/http"
)

var Router router

type router struct {
}

func (r *router) Init(router *gin.Engine) {
	router.
		// GET("/ping", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"message": "pong",
		// 		"status":  http.StatusOK,
		// 		"data":    nil,
		// 	})
		// })
		GET("/api/k8s/pods", Pod.GetPods).
		GET("/api/k8s/pod/detail", Pod.GetPodDetail).
		DELETE("/api/k8s/pod/delete", Pod.DeletePod).
		PUT("/api/k8s/pod/update", Pod.UpdatePod).
		GET("/api/k8s/pod/containers", Pod.GetPodContainers).
		GET("/api/k8s/pod/log", Pod.GetPodLog)

}
