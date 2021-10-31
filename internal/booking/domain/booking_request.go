package domain

import (
	"errors"
	"fmt"
	"time"
)

type BookingRequest struct {
	id string
	checkIn time.Time
	nights int
	sellingRate int
	margin int
}

func (r BookingRequest) Id() string {
	return r.id
}

func NewBookingRequest(id string, checkIn time.Time, nights, sellingRate, margin int) *BookingRequest {
	return &BookingRequest{id: id, checkIn: checkIn, nights: nights, sellingRate: sellingRate, margin: margin}
}

var InvalidCheckInFormatError = errors.New("invalid CheckIn format")
func NewWithCheckInDateString(id, checkInDateString string, nights, sellingRate, margin int) (*BookingRequest, error) {
	checkIn, err := time.Parse("2006-01-02", checkInDateString)
	if err != nil {
		return nil, fmt.Errorf("%w:%v", InvalidCheckInFormatError, err)
	}

	return NewBookingRequest(id, checkIn, nights, sellingRate, margin), nil
}

func (r BookingRequest) Overlaps(otherBookingRequest *BookingRequest) bool {
	checkOut := r.checkIn.Add(time.Hour * 24 * time.Duration(r.nights))
	otherCheckout := otherBookingRequest.checkIn.Add(time.Hour * 24 * time.Duration(otherBookingRequest.nights))

	if r.checkIn.Before(otherCheckout) && checkOut.After(otherBookingRequest.checkIn) {
		return true
	}

	return false
}
