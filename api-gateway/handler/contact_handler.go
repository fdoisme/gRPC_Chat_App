package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/api-gateway/service"
)

type ContactController struct {
	service service.Service
	logger  *zap.Logger
}

func NewContactController(service service.Service, logger *zap.Logger) *ContactController {
	return &ContactController{service, logger}
}

func (ctrl *ContactController) Add(c *gin.Context) {
	username := c.MustGet("username")

	GoodResponseWithData(c, "contact added", http.StatusOK, gin.H{"username": username})
}

func (ctrl *ContactController) Remove(c *gin.Context) {
	username := c.MustGet("username")

	GoodResponseWithData(c, "contact removed", http.StatusOK, gin.H{"username": username})
}
