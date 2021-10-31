package application

import "booking-request-manager/internal/booking/domain"

type BookingRequestTransformer struct {}

func NewBookingRequestTransformer() *BookingRequestTransformer {
	return &BookingRequestTransformer{}
}

func (t BookingRequestTransformer) FromBookingRequestDTOs(bookingRequestDTOs []*BookingRequest) ([]*domain.BookingRequest, error) {
	var bookingRequestVOs []*domain.BookingRequest
	for _, bookingRequestDTO := range bookingRequestDTOs {
		bookingRequestVO, err := t.fromBookingRequestDTO(bookingRequestDTO)
		if err != nil {
			return nil, err
		}

		bookingRequestVOs = append(bookingRequestVOs, bookingRequestVO)
	}
	return bookingRequestVOs, nil
}

func (t BookingRequestTransformer) fromBookingRequestDTO(bookingRequestDTO *BookingRequest) (*domain.BookingRequest, error) {
	bookingRequestVO, err := domain.NewWithCheckInDateString(
		bookingRequestDTO.RequestId,
		bookingRequestDTO.CheckIn,
		bookingRequestDTO.Nights,
		bookingRequestDTO.SellingRate,
		bookingRequestDTO.Margin)
	if err != nil {
		return nil, err
	}
	
	return bookingRequestVO, nil
}