package main

import (
	"booking-request-manager/internal/booking/application/stats"
	"booking-request-manager/internal/booking/infrastructure/http/rest/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", healthCheckHandler)
	router.POST("/stats", postStatsHandler)

	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func healthCheckHandler(context *gin.Context) {
	controller.NewHealthCheckController().Run(context)
}

func postStatsHandler(context *gin.Context) {
	controller.NewStatsController(stats.NewCalculateStatsCommandHandler()).Run(context)
}
