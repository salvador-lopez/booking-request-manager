package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckController struct {}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (c *HealthCheckController) Run(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}
