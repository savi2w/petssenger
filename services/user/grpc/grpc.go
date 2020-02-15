package grpc

import (
	"context"
	"net"

	pb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/user/config"
	"google.golang.org/grpc"
)

type userServer struct{}

func (*userServer) AuthUser(
	ctx context.Context,
	req *pb.AuthUserRequest,
) (*pb.AuthUserResponse, error) {
	// user := req.GetUser()
	return &pb.AuthUserResponse{
		Authed: true,
	}, nil
}

// UserServerListen is a helper function to listen an user gRPC server
func UserServerListen() (net.Listener, error) {
	lis, err := net.Listen("tcp", config.Default.Addr)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	pb.RegisterUserServer(server, &userServer{})
	if err := server.Serve(lis); err != nil {
		lis.Close()
		return nil, err
	}

	return lis, nil
}
