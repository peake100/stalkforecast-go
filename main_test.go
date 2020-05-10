package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func runService(service *grpc.Server, listener net.Listener) {
	err := service.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestForecaster(t *testing.T) {

}
