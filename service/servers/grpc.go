package servers

import (
	"context"
	"github.com/peake100/stalkforecaster-go/protogen/stalk_proto"
	"github.com/peake100/turnup-go"
	"github.com/peake100/turnup-go/models"
	"google.golang.org/grpc"
	"log"
	"net"
)

type StalkForecastServer struct{}

func (server *StalkForecastServer) ForecastPrices(
	ctx context.Context, ticker *stalkproto.Ticker,
) (*stalkproto.Forecast, error) {
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

func makeService() (service *grpc.Server, listener net.Listener, err error) {
	grpcAddress := grpcAddress()

	log.Printf("serving grpc on '%v'\n", grpcAddress)

	listener, err = net.Listen("tcp", grpcAddress)
	if err != nil {
		return nil, nil, err
	}
	service = grpc.NewServer()
	stalkproto.RegisterStalkForecasterServer(service, new(StalkForecastServer))
	return service, listener, nil
}

func RunGrpc(
	monitor *Monitor,
) {
	// Create the service and listeners
	service, listener, err := makeService()
	if err != nil {
		monitor.grpcErrs <- err
		monitor.grpcShutdownComplete <- nil
	}

	// Setup a routine to listen to the shutdown order channel and bring the
	// service to a stop if it triggers.
	go func() {
		<-monitor.shutdownGrpc
		log.Println("grpc shutdown request received")
		// Bring the server to a graceful stop
		service.GracefulStop()
		// Signal the monitor that the shutdown is complete
		defer func() {
			monitor.grpcShutdownComplete <- nil
		}()
	}()

	err = service.Serve(listener)
	if err != nil {
		monitor.grpcErrs <- err
	}

	log.Println("grpc process exiting")
}
