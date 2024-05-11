package main

import (
	"k8s-go-gin/config"
	"k8s-go-gin/controller"
	"k8s-go-gin/service"

	"github.com/gin-gonic/gin"
)

func main() {

	service.K8s.Init()

	r := gin.Default()
	controller.Router.Init(r)
	r.Run(config.ListerAddr)
}
