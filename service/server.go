package service

import (
	"context"
	"github.com/peake100/stalkforecast-go/proto"
	"github.com/peake100/turnup-go"
	"github.com/peake100/turnup-go/models"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type StalkForecastServer struct{}

func (server *StalkForecastServer) ForecastPrices(
	ctx context.Context, ticker *proto.Ticker,
) (*proto.Forecast, error) {
	// Deserialize the proto object into the turnup ticker object
	turnupTicker := turnup.NewPriceTicker(
		int(ticker.PurchasePrice), models.PricePattern(ticker.PreviousPattern),
	)

	// Load it with our values from the protobuff
	for i, price := range ticker.Prices {
		turnupTicker.Prices[i] = int(price)
	}

	prediction, err := turnup.Predict(turnupTicker)
	if err != nil {
		return nil, err
	}

	response := serializePrediction(prediction)

	return response, nil
}

func makeService() (*grpc.Server, net.Listener) {
	servicePort := ":" + os.Getenv("SERVICE_PORT")

	listener, err := net.Listen("tcp", servicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	service := grpc.NewServer()
	proto.RegisterStalkForecasterServer(service, new(StalkForecastServer))
	return service, listener
}

func RunService(errChan chan error) {
	defer close(errChan)
	service, listener := makeService()
	defer service.GracefulStop()

	err := service.Serve(listener)
	if err != nil {
		errChan <- err
	}
}
