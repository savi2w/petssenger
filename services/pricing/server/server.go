package server

import (
	"context"
	"net"

	pricingpb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/pricing/models"
	"google.golang.org/grpc"
)

type pricingServer struct{}

func (*pricingServer) GetPricingFeesByCity(ctx context.Context, req *pricingpb.GetPricingFeesByCityRequest) (*pricingpb.GetPricingFeesByCityResponse, error) {
	city := req.GetCity()
	fees, err := models.GetPricingFees(city)
	if err != nil {
		panic(err)
	}

	proto := models.ProtoPricingFees(fees)
	return proto, nil
}

// PricingServerListen is a helper function to listen and gRPC server
func PricingServerListen() (net.Listener, *grpc.Server, error) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return nil, nil, err
	}

	ser := grpc.NewServer()
	pricingpb.RegisterPricingServer(ser, &pricingServer{})

	if err := ser.Serve(lis); err != nil {
		lis.Close()
		return nil, nil, err
	}

	return lis, ser, nil
}
