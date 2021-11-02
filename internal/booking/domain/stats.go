package domain

import "math"

type Stats struct {
	bookingRequests []*BookingRequest
	avgNight    float64
	minNight    float64
	maxNight    float64
	totalProfit float64
}

func NewStats(bookingRequests []*BookingRequest) Stats {
	s := Stats{bookingRequests: bookingRequests}

	var totalProfit float64
	var totalProfitPerNight float64
	var minNight float64
	var maxNight float64

	for i, bookingRequest := range bookingRequests {
		profitPerNight := s.calculateProfitPerNight(bookingRequest)
		totalProfitPerNight += profitPerNight

		profit := s.calculateProfit(bookingRequest)
		totalProfit += profit

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

	s.avgNight = s.roundFloatValue(totalProfitPerNight / float64(len(bookingRequests)))
	s.minNight = s.roundFloatValue(minNight)
	s.maxNight = s.roundFloatValue(maxNight)
	s.totalProfit = s.roundFloatValue(totalProfit)

	return s
}

func (s Stats) calculateProfit(bookingRequest *BookingRequest) float64 {
	return float64(bookingRequest.sellingRate) * (float64(bookingRequest.margin) / 100)
}

func (s Stats) calculateProfitPerNight(bookingRequest *BookingRequest) float64 {
	return s.calculateProfit(bookingRequest) / float64(bookingRequest.nights)
}

func NewMaximizedProfitStats(bookingRequests []*BookingRequest) Stats {
	noOverlappingCombinations := make([][]*BookingRequest, len(bookingRequests))
	for i, bookingRequest := range bookingRequests {
		noOverlappingCombinations[i] = append(noOverlappingCombinations[i], bookingRequest)
		for j := i+1; j < len(bookingRequests); j++ {
			var overlaps bool
			for _, bookingRequestToCheckOverlap := range noOverlappingCombinations[i] {
				overlaps = bookingRequests[j].Overlaps(bookingRequestToCheckOverlap)
			}
			if !overlaps {
				noOverlappingCombinations[i] = append(noOverlappingCombinations[i], bookingRequests[j])
			}
		}
	}

	var maxProfitStats Stats

	for _, noOverlappingCombination := range noOverlappingCombinations {
		stats := NewStats(noOverlappingCombination)
		if maxProfitStats.TotalProfit() < stats.TotalProfit() {
			maxProfitStats = stats
		}
	}

	return maxProfitStats
}

func (s Stats) BookingRequests() []*BookingRequest {
	return s.bookingRequests
}

func (s Stats) AvgNight() float64 {
	return s.avgNight
}

func (s Stats) MinNight() float64 {
	return s.minNight
}

func (s Stats) MaxNight() float64 {
	return s.maxNight
}

func (s Stats) TotalProfit() float64 {
	return s.totalProfit
}

func (s Stats) roundFloatValue(value float64) float64 {
	return math.Round(value*100)/100
}
