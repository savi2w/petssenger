package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/nglabo/petssenger/services/user/config"
	proto "github.com/nglabo/petssenger/services/user/protos"
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
	c := proto.NewUserClient(gc)

	req := &proto.AuthUserRequest{
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
	c := proto.NewUserClient(gc)

	req := &proto.AuthUserRequest{
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

func TestCreateUser(t *testing.T) {
	seeded := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	json := []byte(fmt.Sprintf(`{"email":"%v@petssenger.com"}`, seeded.Int()))
	res, err := http.Post(
		fmt.Sprintf(
			"http://localhost%v/user",
			config.Default.HTTPPort,
		),
		"application/json",
		bytes.NewBuffer(json),
	)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("createUser is returning a wrong status code")
	}
}

func TestCreateUserWithAnExistingUser(t *testing.T) {
	json := []byte(`{"email":"test@petssenger.com"}`)
	res, err := http.Post(
		fmt.Sprintf(
			"http://localhost%v/user",
			config.Default.HTTPPort,
		),
		"application/json",
		bytes.NewBuffer(json),
	)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("createUser is returning a wrong status code when email already exists")
	}
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	json := []byte(`{"email":"anything"}`)
	res, err := http.Post(
		fmt.Sprintf(
			"http://localhost%v/user",
			config.Default.HTTPPort,
		),
		"application/json",
		bytes.NewBuffer(json),
	)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("createUser is returning a wrong status code when a invalid body is provided")
	}
}
