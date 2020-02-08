package config

import (
	"net"

	"google.golang.org/grpc"
)

// type pricingServer struct{}

// func (*pricingServer) GetPricingFeesByCity(ctx context.Context, req *pricingpb.GetPricingFeesByCityRequest) (*pricingpb.GetPricingFeesByCityResponse, error) {
// 	city := req.GetCity()
// 	fees, err := models.GetFees(city)
// 	if err != nil {
// 		log.Fatalf("City not exist: %v", err)
// 	}

// 	res := pricingpb.GetPricingFeesByCityResponse{

// 	}
// }

// PricingServerListen is a helper function to listen and gRPC server
func PricingServerListen() (*grpc.Server, error) {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return nil, err
	}

	ser := grpc.NewServer()
	// pricingpb.RegisterPricingServer(ser, &pricingServer{})

	if err := ser.Serve(lis); err != nil {
		lis.Close()
		return nil, err
	}

	return ser, nil
}
