package domain

import (
	"math"
	"time"
)

type BookingRequest struct {
	id string
	checkIn time.Time
	nights int
	sellingRate int
	margin int
}

func NewBookingRequest(id string, checkIn time.Time, nights int, sellingRate int, margin int) BookingRequest {
	return BookingRequest{id: id, checkIn: checkIn, nights: nights, sellingRate: sellingRate, margin: margin}
}

type Stats struct {
	avgNight float64
	minNight float64
	maxNight float64
}

func NewStats(avgNight float64, minNight float64, maxNight float64) Stats {
	return Stats{avgNight: avgNight, minNight: minNight, maxNight: maxNight}
}

func NewStatsFromBookingRequests(bookingRequests []BookingRequest) Stats {
	var totalProfit float64
	var avgNight float64
	var minNight float64
	var maxNight float64

	for i, bookingRequest := range bookingRequests {
		profitPerNight := float64(bookingRequest.sellingRate) * (float64(bookingRequest.margin) / 100) / float64(bookingRequest.nights)
		totalProfit += profitPerNight

		if i == 0 {
			minNight = profitPerNight
			maxNight = profitPerNight
			continue
		}

		if maxNight < profitPerNight {
			maxNight = profitPerNight
			continue
		}

		if minNight > profitPerNight {
			minNight = profitPerNight
		}
	}

	avgNight = totalProfit / float64(len(bookingRequests))

	return NewStats(avgNight, minNight, maxNight)
}

func (s Stats) Equal(anotherStat Stats) bool {
	return math.Ceil(s.avgNight) == math.Ceil(anotherStat.avgNight) &&
		math.Ceil(s.minNight) == math.Ceil(anotherStat.minNight) &&
		math.Ceil(s.maxNight) == math.Ceil(anotherStat.maxNight)
}
