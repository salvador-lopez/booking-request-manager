package controller

import (
	"booking-request-manager/internal/booking/application"
	"booking-request-manager/internal/booking/application/maximize"
	"booking-request-manager/internal/booking/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostMaximizeController struct {
	calculateMaximizedProfitCommandHandler *maximize.CalculateMaximizedProfitCommandHandler
}

func NewPostMaximizeController(calculateMaximizedProfitCommandHandler *maximize.CalculateMaximizedProfitCommandHandler) *PostMaximizeController {
	return &PostMaximizeController{calculateMaximizedProfitCommandHandler: calculateMaximizedProfitCommandHandler}
}

func (c *PostMaximizeController) Run(context *gin.Context) {
	var bookingRequests []*application.BookingRequest
	err := context.ShouldBind(&bookingRequests)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	commandResult, err := c.calculateMaximizedProfitCommandHandler.Handle(&maximize.CalculateMaximizedProfitCommand{BookingRequests: bookingRequests})
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
