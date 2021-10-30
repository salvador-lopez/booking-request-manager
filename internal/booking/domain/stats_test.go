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
)

type UnitSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UnitSuite))
}

func (s *UnitSuite) TestShouldReturnStatsAsExpectedWhenTwoBookingRequestsAreProvided() {
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

func (s *UnitSuite) TestShouldReturnStatsAsExpectedWhenThreeBookingRequestsAreProvided() {
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

func (s *UnitSuite) TestShouldReturnMaximizedProfitStatsWhenFourBookingRequestsAreProvided() {
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

func (s *UnitSuite) parseTimeFromDateString(dateString string) time.Time {
	t, err := time.Parse("2006-01-02", dateString)
	s.Require().NoError(err)

	return t
}

