//+build unit

package stats

import (
	"booking-request-manager/internal/booking/application"
	"booking-request-manager/internal/booking/domain"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	bookataRequestId  = "bookata_XY123"
	bookataCheckIn    = "2020-01-01"
	kayeteRequestId   = "Kayete_PP234"
)

type UnitSuite struct {
	suite.Suite
	sut *CalculateStatsCommandHandler
}

func(s *UnitSuite) SetupTest() {
	s.sut = NewCalculateStatsCommandHandler(application.NewBookingRequestTransformer())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UnitSuite))
}

func (s *UnitSuite) TestCalculateStatsAsExpectedWhenCommandHaveTwoBookingRequests() {
	command := &CalculateStatsCommand{
		BookingRequests: []*application.BookingRequest{
			{
				RequestId: bookataRequestId,
				CheckIn: bookataCheckIn,
				Nights: 5,
				SellingRate: 200,
				Margin: 20,
			},
			{
				RequestId: kayeteRequestId,
				CheckIn: "2020-01-04",
				Nights: 4,
				SellingRate: 156,
				Margin: 22,
			},
		},
	}

	commandResult, err := s.sut.Handle(command)
	s.Require().NoError(err)
	expectedCommandResult := &CalculateStatsCommandResult{
		AvgNight: 8.29,
		MinNight: 8,
		MaxNight: 8.58,
	}
	s.Require().Equal(expectedCommandResult, commandResult)
}

func (s *UnitSuite) TestReturnInvalidCheckInFormatErrorWhenCheckInHasInvalidDateString() {
	command := &CalculateStatsCommand{
		BookingRequests: []*application.BookingRequest{
			{
				RequestId: bookataRequestId,
				CheckIn: "invalid date string",
				Nights: 5,
				SellingRate: 200,
				Margin: 20,
			},
		},
	}

	_, err := s.sut.Handle(command)
	s.Require().Error(err)
	s.Require().True(errors.Is(err, domain.InvalidCheckInFormatError))
}