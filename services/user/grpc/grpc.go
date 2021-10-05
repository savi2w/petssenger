package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-pg/pg/v9"
	"github.com/nglabo/petssenger/services/user/config"
	"github.com/nglabo/petssenger/services/user/models"
	proto "github.com/nglabo/petssenger/services/user/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct{}

func (*userServer) AuthUser(
	ctx context.Context,
	req *proto.AuthUserRequest,
) (*proto.AuthUserResponse, error) {
	user := req.GetUser()
	_, err := models.AuthUserByID(user)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(
				codes.PermissionDenied,
				fmt.Sprintf("The user %v was not found", user),
			)
		}

		panic(err)
	}

	return &proto.AuthUserResponse{
		Authed: true,
	}, nil
}

func UserRPCListen() (net.Listener, error) {
	lis, err := net.Listen("tcp", config.Default.Addr)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &userServer{})
	if err := server.Serve(lis); err != nil {
		lis.Close()
		return nil, err
	}

	return lis, nil
}
