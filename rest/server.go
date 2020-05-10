package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/peake100/stalkforecast-go/proto"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"net/http"
	"os"
)

func RunRest(errChan chan error) {
	defer close(errChan)
	grpcServerEndpoint := os.Getenv("SERVICE_HOST") +
		":" +
		os.Getenv("SERVICE_PORT")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterStalkForecasterHandlerFromEndpoint(
		ctx, mux, grpcServerEndpoint, opts,
	)
	if err != nil {
		err = xerrors.Errorf("error registering REST endpoints: %v", err)
		errChan <- err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = http.ListenAndServe(":"+os.Getenv("REST_PORT"), mux)
	if err != nil {
		err = xerrors.Errorf("error serving REST endpoints: %v", err)
		errChan <- err
	}
}
