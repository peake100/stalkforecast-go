package servers

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/peake100/stalkforecaster-go/proto"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func RunRest(
	monitor *ServersMonitor,
) {
	grpcServerEndpoint := os.Getenv("SERVICE_HOST") +
		":" +
		os.Getenv("SERVICE_PORT")
	restPort := ":" + os.Getenv("REST_PORT")

	log.Printf("serving rest gateway on port '%v'\n", restPort)

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
	err := proto.RegisterStalkForecasterHandlerFromEndpoint(
		ctx, mux, grpcServerEndpoint, opts,
	)
	if err != nil {
		err = xerrors.Errorf("error registering REST endpoints: %v", err)
		monitor.restErrs <- err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = http.ListenAndServe(":"+os.Getenv("REST_PORT"), mux)
	if err != nil {
		err = xerrors.Errorf("error serving REST endpoints: %v", err)
		monitor.restErrs <- err
	}

	log.Println("rest process exiting")
}
