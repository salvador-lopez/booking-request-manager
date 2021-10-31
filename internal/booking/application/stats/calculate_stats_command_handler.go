package stats

import (
	"booking-request-manager/internal/booking/domain"
	"errors"
	"fmt"
	"time"
)

type CalculateStatsCommandHandler struct {}

func NewCalculateStatsCommandHandler() *CalculateStatsCommandHandler {
	return &CalculateStatsCommandHandler{}
}

type BookingRequest struct {
	RequestId   string `json:"request_id" binding:"required"`
	CheckIn     string `json:"check_in" binding:"required"`
	Nights      int    `json:"nights" binding:"required"`
	SellingRate int    `json:"selling_rate" binding:"required"`
	Margin      int    `json:"margin" binding:"required"`
}

type CalculateStatsCommand struct {
	BookingRequests []BookingRequest
}

type CalculateStatsCommandResult struct {
	AvgNight float64 `json:"avg_night"`
	MinNight float64 `json:"min_night"`
	MaxNight float64 `json:"max_night"`
}

func (s CalculateStatsCommandHandler) Handle(command CalculateStatsCommand) (*CalculateStatsCommandResult, error) {
	bookingRequestVOs, err := s.buildBookingRequestVOs(command)
	if err != nil {
		return nil, err
	}

	stats := domain.NewStats(bookingRequestVOs)

	return &CalculateStatsCommandResult{
		AvgNight: stats.AvgNight(),
		MinNight: stats.MinNight(),
		MaxNight: stats.MaxNight(),
	}, nil
}

var InvalidCheckInFormatError = errors.New("invalid CheckIn format")
func (s CalculateStatsCommandHandler) buildBookingRequestVOs(command CalculateStatsCommand) ([]domain.BookingRequest, error) {
	var bookingRequestVOs []domain.BookingRequest
	for _, bookingRequest := range command.BookingRequests {
		checkIn, err := time.Parse("2006-01-02", bookingRequest.CheckIn)
		if err != nil {
			return []domain.BookingRequest{}, fmt.Errorf("%w:%v", InvalidCheckInFormatError, err)
		}
		bookingRequestVOs = append(
			bookingRequestVOs,
			domain.NewBookingRequest(bookingRequest.RequestId, checkIn, bookingRequest.Nights, bookingRequest.SellingRate, bookingRequest.Margin))
	}
	return bookingRequestVOs, nil
}


