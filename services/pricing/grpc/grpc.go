package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-pg/pg/v9"
	"github.com/savi2w/petssenger/services/pricing/config"
	"github.com/savi2w/petssenger/services/pricing/models"
	proto "github.com/savi2w/petssenger/services/pricing/protos"
	"github.com/savi2w/petssenger/services/pricing/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type pricingServer struct{}

func (*pricingServer) GetPricingFeesByCity(
	ctx context.Context,
	req *proto.GetFeesByCity,
) (*proto.GetPricingFeesByCityResponse, error) {
	city := req.GetCity()

	fees, err := models.GetPricingFees(city)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf(`The city "%v" was not found`, city),
			)
		}

		panic(err)
	}

	proto := models.ProtoPricingFees(fees)
	return proto, nil
}

func (*pricingServer) GetDynamicFeesByCity(
	ctx context.Context,
	req *proto.GetFeesByCity,
) (*proto.GetDynamicFeesByCityResponse, error) {
	city := req.GetCity()

	fees, err := models.GetDynamicFees(city)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf(`The city "%v" was not found`, city),
			)
		}

		panic(err)
	}

	proto := models.ProtoDynamicFees(fees)
	return proto, nil
}

func (*pricingServer) IncreaseDynamicFeesByCity(
	ctx context.Context,
	req *proto.GetFeesByCity,
) (*proto.Empty, error) {
	city := req.GetCity()

	if err := models.IncreaseDynamicFees(city); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf(`The city "%v" is invalid`, city),
		)
	}

	job := worker.DecreaseDynamicFees.WithArgs(context.Background(), city)
	job.Delay = config.Default.DynamicFeesDecreaseTime

	if err := worker.MainQueue.Add(job); err != nil {
		panic(err)
	}

	return &proto.Empty{}, nil
}

func PricingRPCListen() (net.Listener, error) {
	lis, err := net.Listen("tcp", config.Default.Addr)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	proto.RegisterPricingServer(server, &pricingServer{})
	if err := server.Serve(lis); err != nil {
		lis.Close()
		return nil, err
	}

	return lis, nil
}
