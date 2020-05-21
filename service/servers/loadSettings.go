package servers

import "os"

func grpcAddress() string {
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")
	address := host + ":" + port
	return address
}
