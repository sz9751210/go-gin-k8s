package controller

import (
	"k8s-go-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

type LinkListController struct {
	linkListService *service.LinkListService
}

func NewLinkListController(linkListService *service.LinkListService) *LinkListController {
	return &LinkListController{linkListService: linkListService}
}

func (c *LinkListController) GetLinkList(ctx *gin.Context) {
	data, err := c.linkListService.GetLinkList()
	if err != nil {
		logger.Error("get link list error: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "get link list error" + err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}
