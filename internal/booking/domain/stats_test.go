//+build unit

package domain_test

import (
	"booking-request-manager/internal/booking/domain"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const (
	bookataRequestId = "bookata_XY123"
	kayeteRequestId  = "Kayete_PP234"
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
	stats := domain.NewStatsFromBookingRequests(bookingRequests)

	statsExpected := domain.NewStats(8.29, 8, 8.58)
	s.Require().True(stats.Equal(statsExpected))
}

func (s *UnitSuite) TestShouldReturnStatsAsExpectedWhenThreeBookingRequestsAreProvided() {
	var bookingRequests []domain.BookingRequest

	bookataCheckIn := s.parseTimeFromDateString("2020-01-01")
	kayeteCheckIn := s.parseTimeFromDateString("2020-01-04")
	kayeteSecondCheckIn := s.parseTimeFromDateString("2020-01-07")

	bookingRequests = append(bookingRequests, domain.NewBookingRequest(bookataRequestId, bookataCheckIn, 1, 55, 22))
	bookingRequests = append(bookingRequests, domain.NewBookingRequest(kayeteRequestId, kayeteCheckIn, 1, 49, 21))
	bookingRequests = append(bookingRequests, domain.NewBookingRequest(kayeteRequestId, kayeteSecondCheckIn, 1, 50, 20))
	stats := domain.NewStatsFromBookingRequests(bookingRequests)

	statsExpected := domain.NewStats(10.80, 10, 12.1)
	s.Require().True(stats.Equal(statsExpected))
}

func (s *UnitSuite) parseTimeFromDateString(dateString string) time.Time {
	t, err := time.Parse("2006-01-02", dateString)
	s.Require().NoError(err)

	return t
}

