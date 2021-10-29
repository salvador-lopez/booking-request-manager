//+build unit

package domain_test

import (
	"booking-request-manager/internal/booking/domain"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const (
	id               = "bookata_XY123"
	nights           = 5
	sellingRate      = 200
	margin           = 20
	otherId          = "Kayete_PP234"
	otherNights      = 4
	otherSellingRate = 156
	otherMargin      = 22
)

type UnitSuite struct {
	suite.Suite
	sut *domain.CalculateStatsService
}

func (s *UnitSuite) SetupTest() {
	s.sut = domain.NewCalculateStatsService()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UnitSuite))
}

func (s *UnitSuite) TestShouldReturnStatsAsExpectedWhenTwoBookingRequestsAreProvided() {
	var bookingRequests []*domain.BookingRequest

	checkIn := s.parseTimeFromDateString("2020-01-01")
	otherCheckIn := s.parseTimeFromDateString("2020-01-04")

	bookingRequests = append(bookingRequests, domain.NewBookingRequest(id, checkIn, nights, sellingRate, margin))
	bookingRequests = append(bookingRequests, domain.NewBookingRequest(otherId, otherCheckIn, otherNights, otherSellingRate, otherMargin))
	stats := s.sut.Run(bookingRequests)

	statsExpected := domain.NewStats(8.29, 8, 8.58)
	s.Require().True(stats.Equal(statsExpected))
}

func (s *UnitSuite) parseTimeFromDateString(dateString string) time.Time {
	t, err := time.Parse("2006-01-02", dateString)
	s.Require().NoError(err)

	return t
}

