package service

//revive:disable:import-shadowing reason: Disabled for assert := assert.New(), which is
// the preferred method of using multiple asserts in a test.

import (
	"context"
	"github.com/peake100/stalkforecaster-go/protogen/stalk_proto"
	"github.com/peake100/stalkforecaster-go/service/servers"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
	"time"
)

func requestTimeout() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	return context.WithTimeout(ctx, time.Second*2)
}

type TestBasicPredictionSuite struct {
	suite.Suite
	monitor    *servers.Monitor
	client     stalkproto.StalkForecasterClient
	clientConn *grpc.ClientConn

	grpcAddress string

	forecast *stalkproto.Forecast
}

func (suite *TestBasicPredictionSuite) SetupSuite() {
	envVars := map[string]string{
		"GRPC_HOST": "localhost",
		"GRPC_PORT": "50051",
		"REST_HOST": "localhost",
		"PORT":      "8080",
	}
	for key, value := range envVars {
		err := os.Setenv(key, value)
		if err != nil {
			suite.FailNow("error setting env vars", err)
		}
	}

	log.Println("starting servers")
	suite.monitor = StartServers()

	log.Println("servers started")
	suite.grpcAddress = os.Getenv("GRPC_HOST") +
		":" +
		os.Getenv("GRPC_PORT")

	// Set up a connection to the server.
	log.Println("dialing grpc")
	conn, err := grpc.Dial(suite.grpcAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		suite.FailNow("grpc client did not connect: %v", err)
	}
	log.Println("grpc client connected")

	suite.client = stalkproto.NewStalkForecasterClient(conn)
	suite.clientConn = conn
}

func (suite *TestBasicPredictionSuite) TearDownSuite() {
	defer suite.clientConn.Close()
	defer suite.monitor.WaitOnShutdown()
	defer suite.monitor.ShutdownServers()
}

func (suite *TestBasicPredictionSuite) TestForecaster() {
	ticker := &stalkproto.Ticker{
		PurchasePrice:   100,
		PreviousPattern: stalkproto.PricePatterns_UNKNOWN,
		Prices:          make([]int32, 12),
	}

	ctx, cancel := requestTimeout()
	defer cancel()

	forecast, err := suite.client.ForecastPrices(ctx, ticker)
	suite.NoError(err, "rpc client err")
	if err != nil {
		suite.FailNow(err.Error())
	}
	suite.forecast = forecast
}

func (suite *TestBasicPredictionSuite) TestPrices() {
	forecast := suite.forecast
	suite.Equal(int32(10), forecast.PricesSummary.Min, "min")
	suite.Equal(
		int32(85), forecast.PricesFuture.Guaranteed, "guaranteed",
	)
	suite.Equal(int32(600), forecast.PricesSummary.Max, "max potential")
}

func (suite *TestBasicPredictionSuite) TestSpikeBreakdownLength() {
	forecast := suite.forecast
	suite.Len(forecast.Spikes.Any.Breakdown, 12, "Any")
	suite.Len(forecast.Spikes.Small.Breakdown, 12, "Small")
	suite.Len(forecast.Spikes.Big.Breakdown, 12, "Big")
}

func TestBasicPrediction(t *testing.T) {
	suite.Run(t, new(TestBasicPredictionSuite))
}
