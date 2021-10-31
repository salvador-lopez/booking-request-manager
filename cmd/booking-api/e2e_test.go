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

type E2eSuite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(E2eSuite))
}

func (s *E2eSuite) TestCalculateStatsAsExpected() {
	serverHealthCheckOk := make(chan bool, 1)

	go main()

	go func() {
		for {
			response, err := http.Get("http://127.0.0.1:8080/healthcheck")
			if err == nil && response.StatusCode == http.StatusOK {
				serverHealthCheckOk <- true
				break
			}
		}
	}()

	select {
	case <- serverHealthCheckOk:
		statsRequest, err := os.ReadFile("test-data/stats_request.json")
		s.Require().NoError(err)

		response, err := http.Post("http://127.0.0.1:8080/stats", "application/json", bytes.NewBuffer(statsRequest))
		s.Require().NoError(err)

		s.Require().Equal(http.StatusOK, response.StatusCode)
		respBody, err := ioutil.ReadAll(response.Body)
		s.Require().NoError(err)
		respBodyExpected, err := os.ReadFile("test-data/stats_response.json")
		s.Require().NoError(err)
		s.Require().Equal(string(respBodyExpected), string(respBody))
	case <-time.After(time.Second * 5):
		s.FailNow("deadline exceeded waiting to http server to initialize")
	}
}

