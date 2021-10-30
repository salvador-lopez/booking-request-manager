//+build unit

package domain_test

import (
	"booking-request-manager/internal/booking/domain"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type BookingRequestUnitSuite struct {
	suite.Suite
}

func TestBookingRequestSuite(t *testing.T) {
	suite.Run(t, new(BookingRequestUnitSuite))
}

func (s BookingRequestUnitSuite) TestOverlaps() {
	bookataCheckIn := s.parseTimeFromDateString("2020-01-01")
	kayeteCheckIn := s.parseTimeFromDateString("2020-01-02")
	bookingRequest := domain.NewBookingRequest(bookataRequestId, bookataCheckIn, 2, 10, 5)
	overlappingBookingRequest := domain.NewBookingRequest(kayeteRequestId, kayeteCheckIn, 1, 10, 5)

	s.Require().True(bookingRequest.Overlaps(overlappingBookingRequest))
}

func (s BookingRequestUnitSuite) TestNoOverlaps() {
	bookataCheckIn := s.parseTimeFromDateString("2020-01-01")
	kayeteCheckIn := s.parseTimeFromDateString("2020-01-03")
	bookingRequest := domain.NewBookingRequest(bookataRequestId, bookataCheckIn, 1, 10, 5)
	overlappingBookingRequest := domain.NewBookingRequest(kayeteRequestId, kayeteCheckIn, 1, 10, 5)

	s.Require().False(bookingRequest.Overlaps(overlappingBookingRequest))
}

func (s *BookingRequestUnitSuite) parseTimeFromDateString(dateString string) time.Time {
	t, err := time.Parse("2006-01-02", dateString)
	s.Require().NoError(err)

	return t
}