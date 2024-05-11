package main

import (
	"k8s-go-gin/config"
	"k8s-go-gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.Router.Init(r)
	r.Run(config.ListerAddr)
}
