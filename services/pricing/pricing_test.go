package main

import (
	"context"
	"testing"

	pb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/pricing/config"
	"google.golang.org/grpc"
)

func TestGetPricingFeesByCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := pb.NewPricingClient(gc)
	req := &pb.GetFeesByCity{
		City: "SAO_PAULO",
	}

	_, err = c.GetPricingFeesByCity(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestGetDynamicFeesByCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := pb.NewPricingClient(gc)
	req := &pb.GetFeesByCity{
		City: "SAO_PAULO",
	}

	_, err = c.GetDynamicFeesByCity(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestIncreaseDynamicFeesByCity(t *testing.T) {
	gc, err := grpc.Dial(config.Default.Addr, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}

	defer gc.Close()
	c := pb.NewPricingClient(gc)
	req := &pb.GetFeesByCity{
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

	next := fees.GetDynamic()
	if old+config.Default.DynamicFeesIncreaseRate != next {
		t.Errorf("Dynamic is not being incremented properly")
	}
}
