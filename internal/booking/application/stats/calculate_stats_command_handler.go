package stats

import (
	"booking-request-manager/internal/booking/application"
	"booking-request-manager/internal/booking/domain"
)

type CalculateStatsCommandHandler struct {
	bookingRequestTransformer *application.BookingRequestTransformer
}

func NewCalculateStatsCommandHandler(bookingRequestTransformer *application.BookingRequestTransformer) *CalculateStatsCommandHandler {
	return &CalculateStatsCommandHandler{bookingRequestTransformer: bookingRequestTransformer}
}

type CalculateStatsCommand struct {
	BookingRequests []*application.BookingRequest
}

type CalculateStatsCommandResult struct {
	AvgNight float64 `json:"avg_night"`
	MinNight float64 `json:"min_night"`
	MaxNight float64 `json:"max_night"`
}

func (h *CalculateStatsCommandHandler) Handle(command *CalculateStatsCommand) (*CalculateStatsCommandResult, error) {
	bookingRequestVOs, err := h.bookingRequestTransformer.FromBookingRequestDTOs(command.BookingRequests)
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


