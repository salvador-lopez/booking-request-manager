package controller

import (
	"booking-request-manager/internal/booking/application"
	"booking-request-manager/internal/booking/application/stats"
	"booking-request-manager/internal/booking/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostStatsController struct {
	calculateStatsCommandHandler *stats.CalculateStatsCommandHandler
}

func NewPostStatsController(calculateStatsCommandHandler *stats.CalculateStatsCommandHandler) *PostStatsController {
	return &PostStatsController{calculateStatsCommandHandler: calculateStatsCommandHandler}
}

func (c *PostStatsController) Run(context *gin.Context) {
	var bookingRequests []*application.BookingRequest
	err := context.ShouldBind(&bookingRequests)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	commandResult, err := c.calculateStatsCommandHandler.Handle(&stats.CalculateStatsCommand{BookingRequests: bookingRequests})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, domain.InvalidCheckInFormatError) {
			statusCode = http.StatusBadRequest
		}
		context.JSON(statusCode, err.Error())
		return
	}

	context.JSON(http.StatusOK, commandResult)
}
