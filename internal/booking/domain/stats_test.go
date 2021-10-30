//+build unit

package domain_test

import (
	"booking-request-manager/internal/booking/domain"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const (
	bookataRequestId  = "bookata_XY123"
	kayeteRequestId   = "Kayete_PP234"
	atropoteRequestId = "atropote_AA930"
	acmeRequestId     = "acme_AAAAA"

	timeLayout = "2006-01-02"
)

type StatsUnitSuite struct {
	suite.Suite
}

func TestStatsSuite(t *testing.T) {
	suite.Run(t, new(StatsUnitSuite))
}

func (s *StatsUnitSuite) TestShouldReturnStatsAsExpectedWhenTwoBookingRequestsAreProvided() {
	var bookingRequests []domain.BookingRequest

	bookataCheckIn := s.parseTimeFromDateString("2020-01-01")
	kayeteCheckIn := s.parseTimeFromDateString("2020-01-04")

	bookingRequests = append(bookingRequests, domain.NewBookingRequest(bookataRequestId, bookataCheckIn, 5, 200, 20))
	bookingRequests = append(bookingRequests, domain.NewBookingRequest(kayeteRequestId, kayeteCheckIn, 4, 156, 22))
	stats := domain.NewStats(bookingRequests)

	s.Require().Equal(8.29, stats.AvgNight())
	s.Require().Equal(8.0, stats.MinNight())
	s.Require().Equal(8.58, stats.MaxNight())
	s.Require().Equal(74.32, stats.TotalProfit())
	s.Require().Equal(bookingRequests, stats.BookingRequests())
}

func (s *StatsUnitSuite) TestShouldReturnStatsAsExpectedWhenThreeBookingRequestsAreProvided() {
	var bookingRequests []domain.BookingRequest

	bookataCheckIn := s.parseTimeFromDateString("2020-01-01")
	kayeteCheckIn := s.parseTimeFromDateString("2020-01-04")
	kayeteSecondCheckIn := s.parseTimeFromDateString("2020-01-07")

	bookingRequests = append(bookingRequests, domain.NewBookingRequest(bookataRequestId, bookataCheckIn, 1, 55, 22))
	bookingRequests = append(bookingRequests, domain.NewBookingRequest(kayeteRequestId, kayeteCheckIn, 1, 49, 21))
	bookingRequests = append(bookingRequests, domain.NewBookingRequest(kayeteRequestId, kayeteSecondCheckIn, 1, 50, 20))
	stats := domain.NewStats(bookingRequests)

	s.Require().Equal(10.796666666666667, stats.AvgNight())
	s.Require().Equal(10.0, stats.MinNight())
	s.Require().Equal(12.1, stats.MaxNight())
	s.Require().Equal(32.39, stats.TotalProfit())
	s.Require().Equal(bookingRequests, stats.BookingRequests())
}

func (s *StatsUnitSuite) TestShouldReturnMaximizedProfitStatsWhenFourBookingRequestsAreProvided() {
	var bookingRequests []domain.BookingRequest

	bookataCheckIn := s.parseTimeFromDateString("2020-01-01")
	kayeteCheckIn := s.parseTimeFromDateString("2020-01-04")
	atropoteCheckIn := s.parseTimeFromDateString("2020-01-04")
	acmeCheckIn := s.parseTimeFromDateString("2020-01-10")

	bookataRequest := domain.NewBookingRequest(bookataRequestId, bookataCheckIn, 5, 200, 20)
	acmeRequest := domain.NewBookingRequest(acmeRequestId, acmeCheckIn, 4, 160, 30)
	bookingRequests = append(bookingRequests, bookataRequest)
	bookingRequests = append(bookingRequests, domain.NewBookingRequest(kayeteRequestId, kayeteCheckIn, 4, 156, 5))
	bookingRequests = append(bookingRequests, domain.NewBookingRequest(atropoteRequestId, atropoteCheckIn, 4, 150, 6))
	bookingRequests = append(bookingRequests, acmeRequest)

	var expectedBookingRequests []domain.BookingRequest
	expectedBookingRequests = append(expectedBookingRequests, bookataRequest)
	expectedBookingRequests = append(expectedBookingRequests, acmeRequest)

	maximizedProfitStats := domain.NewMaximizedProfitStats(bookingRequests)
	s.Require().Equal(expectedBookingRequests, maximizedProfitStats.BookingRequests())
	s.Require().Equal(88.0, maximizedProfitStats.TotalProfit())
	s.Require().Equal(10.0, maximizedProfitStats.AvgNight())
	s.Require().Equal(8.0, maximizedProfitStats.MinNight())
	s.Require().Equal(12.0, maximizedProfitStats.MaxNight())
}

func (s StatsUnitSuite) parseTimeFromDateString(dateString string) time.Time {
	t, err := time.Parse(timeLayout, dateString)
	s.Require().NoError(err)

	return t
}

func BenchmarkStats(b *testing.B) {
	bookataCheckIn, _ := time.Parse(timeLayout, "2020-01-01")
	bookingRequests := []domain.BookingRequest {
		domain.NewBookingRequest(bookataRequestId, bookataCheckIn, 5, 200, 20),
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		domain.NewStats(bookingRequests)
	}
}

func BenchmarkMaximize(b *testing.B) {
	bookataCheckIn, _ := time.Parse(timeLayout, "2020-01-01")
	kayeteCheckIn, _ := time.Parse(timeLayout, "2020-01-04")
	atropoteCheckIn, _ := time.Parse(timeLayout, "2020-01-04")
	acmeCheckIn, _ := time.Parse(timeLayout, "2020-01-10")

	bookingRequests := []domain.BookingRequest {
		domain.NewBookingRequest(bookataRequestId, bookataCheckIn, 5, 200, 20),
		domain.NewBookingRequest(kayeteRequestId, kayeteCheckIn, 4, 156, 5),
		domain.NewBookingRequest(atropoteRequestId, atropoteCheckIn, 4, 150, 6),
		domain.NewBookingRequest(acmeRequestId, acmeCheckIn, 4, 160, 30),
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		domain.NewMaximizedProfitStats(bookingRequests)
	}
}

