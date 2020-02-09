package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis/v7"
	pricingpb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/pricing/models"
	"github.com/weslenng/petssenger/services/pricing/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type pricingServer struct {
	pg    *pg.DB
	redis *redis.Client
}

const addr = "0.0.0.0:50051"

func (ps *pricingServer) GetPricingFeesByCity(
	ctx context.Context,
	req *pricingpb.GetFeesByCity,
) (*pricingpb.GetPricingFeesByCityResponse, error) {
	city := req.GetCity()

	fees, err := models.GetPricingFees(city, ps.pg, ps.redis)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("The city %v was not found", city),
			)
		}

		panic(err)
	}

	proto := models.ProtoPricingFees(fees)
	return proto, nil
}

func (ps *pricingServer) GetDynamicFeesByCity(
	ctx context.Context,
	req *pricingpb.GetFeesByCity,
) (*pricingpb.GetDynamicFeesByCityResponse, error) {
	city := req.GetCity()

	fees, err := models.GetDynamicFees(city, ps.pg)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("The city %v was not found", city),
			)
		}

		panic(err)
	}

	proto := models.ProtoDynamicFees(fees)
	return proto, nil
}

func (ps *pricingServer) IncreaseDynamicFeesByCity(
	ctx context.Context,
	req *pricingpb.GetFeesByCity,
) (*pricingpb.Empty, error) {
	city := req.GetCity()

	err := models.IncreaseDynamicFees(city, ps.pg)
	if err != nil {
		panic(err)
	}

	job := worker.DecreaseDynamicFees.WithArgs(context.Background(), city)
	job.Delay = 5 * time.Second

	err = worker.MainQueue.Add(job)
	if err != nil {
		panic(err)
	}

	return &pricingpb.Empty{}, nil
}

// PricingServerListen is a helper function to listen an pricing gRPC server
func PricingServerListen(pg *pg.DB, redis *redis.Client) (net.Listener, *grpc.Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	ser := grpc.NewServer()
	ps := &pricingServer{
		pg:    pg,
		redis: redis,
	}

	pricingpb.RegisterPricingServer(ser, ps)
	if err := ser.Serve(lis); err != nil {
		lis.Close()
		return nil, nil, err
	}

	return lis, ser, nil
}
