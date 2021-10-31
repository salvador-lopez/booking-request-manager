package controller

import (
	"booking-request-manager/internal/booking/application/stats"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostStatsController struct {
	calculateStatsCommandHandler *stats.CalculateStatsCommandHandler
}

func NewStatsController(calculateStatsCommandHandler *stats.CalculateStatsCommandHandler) *PostStatsController {
	return &PostStatsController{calculateStatsCommandHandler: calculateStatsCommandHandler}
}

func (c *PostStatsController) Run(context *gin.Context) {
	var bookingRequests []stats.BookingRequest
	err := context.ShouldBind(&bookingRequests)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	commandResult, err := c.calculateStatsCommandHandler.Handle(stats.CalculateStatsCommand{BookingRequests: bookingRequests})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, stats.InvalidCheckInFormatError) {
			statusCode = http.StatusBadRequest
		}
		context.JSON(statusCode, err.Error())
		return
	}

	context.JSON(http.StatusOK, commandResult)
}
