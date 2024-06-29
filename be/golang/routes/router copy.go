// package controller

// import (
// 	"github.com/gin-gonic/gin"
// 	// "net/http"
// )

// var Router router

// type router struct {
// }

// func (r *router) Init(router *gin.Engine) {

// 	router.
// 		// GET("/ping", func(c *gin.Context) {
// 		// 	c.JSON(http.StatusOK, gin.H{
// 		// 		"message": "pong",
// 		// 		"status":  http.StatusOK,
// 		// 		"data":    nil,
// 		// 	})
// 		// })
// 		// pod操作
// 		GET("/api/k8s/pods", Pod.GetPods).
// 		GET("/api/k8s/pod/detail", Pod.GetPodDetail).
// 		DELETE("/api/k8s/pod/delete", Pod.DeletePod).
// 		PUT("/api/k8s/pod/update", Pod.UpdatePod).
// 		GET("/api/k8s/pod/containers", Pod.GetPodContainers).
// 		GET("/api/k8s/pod/log", Pod.GetPodLog).
// 		GET("/api/k8s/pod/numpernp", Pod.GetPodNumPerNp).
// 		// deployment操作
// 		GET("/api/k8s/deployments", Deployment.GetDeployments).
// 		GET("/api/k8s/deployment/detail", Deployment.GetDeploymentDetail).
// 		PUT("/api/k8s/deployment/scale", Deployment.ScaleDeployment).
// 		POST("/api/k8s/deployment/create", Deployment.CreateDeployment).
// 		DELETE("/api/k8s/deployment/delete", Deployment.DeleteDeployment).
// 		PUT("/api/k8s/deployment/restart", Deployment.RestartDeployment).
// 		// namespace操作
// 		GET("/api/k8s/namespaces", Namespace.GetNamespaces).
// 		GET("/api/k8s/namespace/detail", Namespace.GetNamespaceDetail).
// 		DELETE("/api/k8s/namespace/delete", Namespace.DeleteNamespace).
// 		// devops操作
// 		GET("/api/devops/linklist")
// }
