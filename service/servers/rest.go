package servers

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/peake100/stalkforecaster-go/protogen/stalk_proto"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func RunRest(
	monitor *ServersMonitor,
) {
	grpcServerEndpoint := grpcAddress()
	restAddress := restAddress()

	log.Printf("serving rest gateway on '%v'\n", restAddress)

	// Get a cancelable context for the server, and start a goroutine to listen for
	// a shutdown command and execute it.
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		<-monitor.shutdownRest
		log.Println("rest shutdown request received")
		cancelFunc()
		monitor.restShutdownComplete <- nil
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := stalkproto.RegisterStalkForecasterHandlerFromEndpoint(
		ctx, mux, grpcServerEndpoint, opts,
	)
	if err != nil {
		err = xerrors.Errorf("error registering REST endpoints: %v", err)
		monitor.restErrs <- err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = http.ListenAndServe(restAddress, mux)
	if err != nil {
		err = xerrors.Errorf("error serving REST endpoints: %v", err)
		monitor.restErrs <- err
	}

	log.Println("rest process exiting")
}
