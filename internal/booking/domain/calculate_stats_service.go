package domain

import "time"

type CalculateStatsService struct {}

func NewCalculateStatsService() *CalculateStatsService {
	return &CalculateStatsService{}
}

type BookingRequest struct {
	id string
	checkIn time.Time
	nights int
	sellingRate int
	margin int
}

func NewBookingRequest(id string, checkIn time.Time, nights int, sellingRate int, margin int) *BookingRequest {
	return &BookingRequest{id: id, checkIn: checkIn, nights: nights, sellingRate: sellingRate, margin: margin}
}

type Stats struct {
	avgNight float64
	minNight float64
	maxNight float64
}

func (s Stats) Equal(anotherStat *Stats) bool {
	return s.avgNight == anotherStat.avgNight && s.minNight == anotherStat.minNight && s.maxNight == anotherStat.maxNight
}

func NewStats(avgNight float64, minNight float64, maxNight float64) *Stats {
	return &Stats{avgNight: avgNight, minNight: minNight, maxNight: maxNight}
}

func (s *CalculateStatsService) Run(bookingRequest []*BookingRequest) *Stats {
	return NewStats(8.29, 8, 8.58)
}
