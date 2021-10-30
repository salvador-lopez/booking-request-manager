package domain

type Stats struct {
	bookingRequests []BookingRequest
	avgNight    float64
	minNight    float64
	maxNight    float64
	totalProfit float64
}

func NewStats(bookingRequests []BookingRequest) Stats {
	var totalProfit float64
	var totalProfitPerNight float64
	var minNight float64
	var maxNight float64

	for i, bookingRequest := range bookingRequests {
		profitPerNight := Stats{}.calculateProfitPerNight(bookingRequest)
		totalProfitPerNight += profitPerNight

		profit := Stats{}.calculateProfit(bookingRequest)
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

	return Stats{
		bookingRequests: bookingRequests,
		avgNight:        totalProfitPerNight / float64(len(bookingRequests)),
		minNight:        minNight,
		maxNight:        maxNight,
		totalProfit:     totalProfit,
	}
}

func (s Stats) calculateProfit(bookingRequest BookingRequest) float64 {
	return float64(bookingRequest.sellingRate) * (float64(bookingRequest.margin) / 100)
}

func (s Stats) calculateProfitPerNight(bookingRequest BookingRequest) float64 {
	return s.calculateProfit(bookingRequest) / float64(bookingRequest.nights)
}

func NewMaximizedProfitStats(bookingRequests []BookingRequest) Stats {
	return NewStats(Stats{}.calculateBestRequestsCombination(bookingRequests))
}

func (s Stats) calculateBestRequestsCombination(bookingRequests []BookingRequest) []BookingRequest {
	noOverlappingCombinations := make([][]BookingRequest, len(bookingRequests))
	for i, bookingRequest := range bookingRequests {
		noOverlappingCombinations[i] = append(noOverlappingCombinations[i], bookingRequest)
		for j := i+1; j < len(bookingRequests); j++ {
			overlaps := false
			for _, bookingRequestToCheckOverlap := range noOverlappingCombinations[i] {
				if bookingRequests[j].Overlaps(bookingRequestToCheckOverlap) {
					overlaps = true
				}
			}
			if !overlaps {
				noOverlappingCombinations[i] = append(noOverlappingCombinations[i], bookingRequests[j])
			}
		}
	}

	var maxTotalProfit float64
	var bestRequestCombination []BookingRequest

	for _, noOverlappingCombination := range noOverlappingCombinations {
		totalProfit := s.calculateTotalProfit(noOverlappingCombination)
		if maxTotalProfit < totalProfit {
			maxTotalProfit = totalProfit
			bestRequestCombination = noOverlappingCombination
		}
	}
	return bestRequestCombination
}

func (s Stats) calculateTotalProfit(bookingRequests []BookingRequest) float64 {
	var totalProfit float64

	for _, bookingRequest := range bookingRequests {
		totalProfit += s.calculateProfit(bookingRequest)
	}

	return totalProfit
}

func (s Stats) BookingRequests() []BookingRequest {
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
