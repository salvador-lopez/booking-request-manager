//+build unit

package maximize_test

import (
	"booking-request-manager/internal/booking/application"
	"booking-request-manager/internal/booking/application/maximize"
	"booking-request-manager/internal/booking/domain"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	bookataRequestId  = "bookata_XY123"
	kayeteRequestId   = "Kayete_PP234"
	atropoteRequestId = "atropote_AA930"
	acmeRequestId     = "acme_AAAAA"
)

type UnitSuite struct {
	suite.Suite
	sut *maximize.CalculateMaximizedProfitCommandHandler
}

func(s *UnitSuite) SetupTest() {
	s.sut = maximize.NewCalculateMaximizedProfitCommandHandler(application.NewBookingRequestTransformer())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UnitSuite))
}

func (s *UnitSuite) TestMaximizeProfitAsExpectedWhenFourBookingRequestsAreProvided() {
	command := &maximize.CalculateMaximizedProfitCommand{
		BookingRequests: []*application.BookingRequest{
			{
				RequestId: bookataRequestId,
				CheckIn: "2020-01-01",
				Nights: 5,
				SellingRate: 200,
				Margin: 20,
			},
			{
				RequestId: kayeteRequestId,
				CheckIn: "2020-01-04",
				Nights: 4,
				SellingRate: 156,
				Margin: 5,
			},
			{
				RequestId: atropoteRequestId,
				CheckIn: "2020-01-04",
				Nights: 4,
				SellingRate: 150,
				Margin: 6,
			},
			{
				RequestId: acmeRequestId,
				CheckIn: "2020-01-10",
				Nights: 4,
				SellingRate: 160,
				Margin: 30,
			},
		},
	}

	commandResult, err := s.sut.Handle(command)
	s.Require().NoError(err)
	expectedCommandResult := &maximize.CalculateMaximizedProfitCommandResult{
		RequestIds: []string{bookataRequestId, acmeRequestId},
		TotalProfit: 88,
		AvgNight: 10,
		MinNight: 8,
		MaxNight: 12,
	}
	s.Require().Equal(expectedCommandResult, commandResult)
}

func (s *UnitSuite) TestReturnInvalidCheckInFormatErrorWhenCheckInHasInvalidDateString() {
	command := &maximize.CalculateMaximizedProfitCommand{
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
