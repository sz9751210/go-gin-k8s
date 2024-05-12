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
		GET("/api/k8s/pods", Pod.GetPods)
}
