package main

import (
	"booking-request-manager/internal/booking/application"
	calculate_maximized_profit "booking-request-manager/internal/booking/application/maximize"
	"booking-request-manager/internal/booking/application/stats"
	"booking-request-manager/internal/booking/infrastructure/http/rest/controller"
	"github.com/gin-gonic/gin"
)

const (
	addr            = ":8080"
	healthCheckPath = "/healthcheck"
	statsPath       = "/stats"
	maximizePath    = "/maximize"
)

func main() {
	router := gin.Default()
	router.GET(healthCheckPath, healthCheckHandler)
	router.POST(statsPath, postStatsHandler)
	router.POST(maximizePath, postMaximizeHandler)

	err := router.Run(addr)
	if err != nil {
		panic(err)
	}
}

func healthCheckHandler(context *gin.Context) {
	controller.NewHealthCheckController().Run(context)
}

func postStatsHandler(context *gin.Context) {
	controller.NewPostStatsController(stats.NewCalculateStatsCommandHandler(application.NewBookingRequestTransformer())).Run(context)
}

func postMaximizeHandler(context *gin.Context) {
	controller.NewPostMaximizeController(calculate_maximized_profit.NewCalculateMaximizedProfitCommandHandler(application.NewBookingRequestTransformer())).Run(context)
}
