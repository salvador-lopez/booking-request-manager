package domain

import "time"

type BookingRequest struct {
	id string
	checkIn time.Time
	nights int
	sellingRate int
	margin int
}

func (r BookingRequest) Overlaps(otherBookingRequest BookingRequest) bool {
	checkOut := r.checkIn.Add(time.Hour * 24 * time.Duration(r.nights))
	otherCheckout := otherBookingRequest.checkIn.Add(time.Hour * 24 * time.Duration(otherBookingRequest.nights))

	if r.checkIn.Before(otherCheckout) && checkOut.After(otherBookingRequest.checkIn) {
		return true
	}

	return false
}

func NewBookingRequest(id string, checkIn time.Time, nights int, sellingRate int, margin int) BookingRequest {
	return BookingRequest{id: id, checkIn: checkIn, nights: nights, sellingRate: sellingRate, margin: margin}
}
