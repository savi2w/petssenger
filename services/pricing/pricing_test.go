package main

import (
	"context"
	"testing"

	"github.com/omppye-tech/petssenger/services/pricing/config"
	proto "github.com/omppye-tech/petssenger/services/pricing/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetPricingFeesByCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := proto.NewPricingClient(gc)
	req := &proto.GetFeesByCity{
		City: "SAO_PAULO",
	}

	_, err = c.GetPricingFeesByCity(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPricingFeesByCityWithInvalidCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := proto.NewPricingClient(gc)
	req := &proto.GetFeesByCity{
		City: "INVALID_CITY",
	}

	_, err = c.GetPricingFeesByCity(context.Background(), req)
	if err != nil {
		c, _ := status.FromError(err)
		if c.Code() != codes.NotFound {
			t.Error(err)
		}
	}
}

func TestGetDynamicFeesByCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := proto.NewPricingClient(gc)
	req := &proto.GetFeesByCity{
		City: "SAO_PAULO",
	}

	_, err = c.GetDynamicFeesByCity(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestGetDynamicFeesByCityWithInvalidCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := proto.NewPricingClient(gc)
	req := &proto.GetFeesByCity{
		City: "INVALID_CITY",
	}

	_, err = c.GetDynamicFeesByCity(context.Background(), req)
	if err != nil {
		c, _ := status.FromError(err)
		if c.Code() != codes.NotFound {
			t.Error(err)
		}
	}
}

func TestIncreaseDynamicFeesByCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := proto.NewPricingClient(gc)
	req := &proto.GetFeesByCity{
		City: "SAO_PAULO",
	}

	fees, err := c.GetDynamicFeesByCity(context.Background(), req)
	if err != nil {
		t.Error(err)
	}

	old := fees.GetDynamic()
	_, err = c.IncreaseDynamicFeesByCity(context.Background(), req)
	if err != nil {
		t.Error(err)
	}

	fees, err = c.GetDynamicFeesByCity(context.Background(), req)
	if err != nil {
		t.Error(err)
	}

	expected := old + config.Default.DynamicFeesIncreaseRate
	newer := fees.GetDynamic()
	if expected < newer {
		t.Errorf("Dynamic fees is not being incremented properly")
	}
}

func TestIncreaseDynamicFeesByCityWithInvalidCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := proto.NewPricingClient(gc)
	req := &proto.GetFeesByCity{
		City: "INVALID_CITY",
	}

	_, err = c.IncreaseDynamicFeesByCity(context.Background(), req)
	if err != nil {
		c, _ := status.FromError(err)
		if c.Code() != codes.InvalidArgument {
			t.Error(err)
		}
	}
}
