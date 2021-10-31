package maximize

import (
	"booking-request-manager/internal/booking/application"
	"booking-request-manager/internal/booking/domain"
)

type CalculateMaximizedProfitCommandHandler struct {
	bookingRequestTransformer *application.BookingRequestTransformer
}

func NewCalculateMaximizedProfitCommandHandler(bookingRequestTransformer *application.BookingRequestTransformer) *CalculateMaximizedProfitCommandHandler {
	return &CalculateMaximizedProfitCommandHandler{bookingRequestTransformer: bookingRequestTransformer}
}

type CalculateMaximizedProfitCommand struct {
	BookingRequests []*application.BookingRequest
}

type CalculateMaximizedProfitCommandResult struct {
	RequestIds  []string `json:"request_ids"`
	TotalProfit float64  `json:"total_profit"`
	AvgNight    float64  `json:"avg_night"`
	MinNight    float64  `json:"min_night"`
	MaxNight    float64  `json:"max_night"`
}

func (h *CalculateMaximizedProfitCommandHandler) Handle(command *CalculateMaximizedProfitCommand) (*CalculateMaximizedProfitCommandResult, error) {
	bookingRequestVOs, err := h.bookingRequestTransformer.FromBookingRequestDTOs(command.BookingRequests)
	if err != nil {
		return nil, err
	}

	stats := domain.NewMaximizedProfitStats(bookingRequestVOs)
	return &CalculateMaximizedProfitCommandResult{
		RequestIds: fromBookingRequestsToRequestIds(stats.BookingRequests()),
		TotalProfit: stats.TotalProfit(),
		AvgNight:    stats.AvgNight(),
		MinNight:    stats.MinNight(),
		MaxNight:    stats.MaxNight(),
	}, nil
}

func fromBookingRequestsToRequestIds(bookingRequests []*domain.BookingRequest) []string {
	var requestIds []string
	for _, bookingRequest := range bookingRequests {
		requestIds = append(requestIds, bookingRequest.Id())
	}

	return requestIds
}