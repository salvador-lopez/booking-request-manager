//+build e2e

package main

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

const addrWithProtocol = "http://"+addr

type E2eSuite struct {
	suite.Suite
	serverHealthCheckOk chan bool
}

func (s *E2eSuite) SetupSuite() {
	go main()
}

func (s *E2eSuite) SetupTest() {
	s.serverHealthCheckOk = make(chan bool, 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(E2eSuite))
}

func (s *E2eSuite) TestCalculateStatsAsExpected() {
	select {
	case <-s.runHealthCheckUntilOk():
		s.assertStatsEndpoint()
	case <-time.After(time.Second * 5):
		s.FailNow("deadline exceeded waiting to http server to initialize")
	}
}

func (s *E2eSuite) runHealthCheckUntilOk() chan bool {
	go func() {
		for {
			response, err := http.Get(addrWithProtocol+healthCheckPath)
			if err == nil && response.StatusCode == http.StatusOK {
				s.serverHealthCheckOk <- true
				break
			}
		}
	}()
	return s.serverHealthCheckOk
}

func (s *E2eSuite) assertStatsEndpoint() {
	statsRequest, err := os.ReadFile("fixtures/stats_request.json")
	s.Require().NoError(err)

	response, err := http.Post(addrWithProtocol+statsPath, "application/json", bytes.NewBuffer(statsRequest))
	s.Require().NoError(err)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	respBody, err := ioutil.ReadAll(response.Body)
	s.Require().NoError(err)
	respBodyExpected, err := os.ReadFile("fixtures/stats_response.json")
	s.Require().NoError(err)
	s.Require().Equal(string(respBodyExpected), string(respBody))
}

func (s *E2eSuite) TestCalculateMaximizedProfitAsExpected() {
	select {
	case <-s.runHealthCheckUntilOk():
		s.assertMaximizeEndpoint()
	case <-time.After(time.Second * 5):
		s.FailNow("deadline exceeded waiting to http server to initialize")
	}
}

func (s *E2eSuite) assertMaximizeEndpoint() {
	maximizeRequest, err := os.ReadFile("fixtures/maximize_request.json")
	s.Require().NoError(err)

	response, err := http.Post(addrWithProtocol+maximizePath, "application/json", bytes.NewBuffer(maximizeRequest))
	s.Require().NoError(err)

	s.Require().Equal(http.StatusOK, response.StatusCode)
	respBody, err := ioutil.ReadAll(response.Body)
	s.Require().NoError(err)
	respBodyExpected, err := os.ReadFile("fixtures/maximize_response.json")
	s.Require().NoError(err)
	s.Require().Equal(string(respBodyExpected), string(respBody))
}
