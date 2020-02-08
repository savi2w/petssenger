package server

import (
	"net"

	"google.golang.org/grpc"
)

// Listen serve an gRPC service
func Listen() (*grpc.Server, error) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return nil, err
	}

	defer lis.Close()
	ser := grpc.NewServer()
	if err := ser.Serve(lis); err != nil {
		return nil, err
	}

	defer ser.GracefulStop()
	return ser, nil
}
