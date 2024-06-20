package main

import (
	"k8s-go-gin/config"
	"k8s-go-gin/controller"
	"k8s-go-gin/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	service.K8s.Init()

	r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:7070"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	controller.Router.Init(r)
	r.Run(config.ListerAddr)
}
