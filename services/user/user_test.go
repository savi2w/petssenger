package main

import (
	"context"
	"testing"

	pb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/user/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAuthUser(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := pb.NewUserClient(gc)

	req := &pb.AuthUserRequest{
		User: "08842beb-a4fc-4cb2-9f87-d80f1a2d5045",
	}

	_, err = c.AuthUser(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}
func TestAuthUserWithInvalidUser(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := pb.NewUserClient(gc)

	req := &pb.AuthUserRequest{
		User: "08274e0b-1260-400d-b909-5bf49c14caef",
	}

	_, err = c.AuthUser(context.Background(), req)
	if err != nil {
		c, _ := status.FromError(err)
		if c.Code() != codes.PermissionDenied {
			t.Error(err)
		}
	}
}
