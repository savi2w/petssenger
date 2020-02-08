package models

import (
	pricingpb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/pricing/config"
)

// Fees payload structure
type Fees struct {
	ID       string
	Base     int32
	Distance int32
	Dynamic  int32
	Minute   int32
	Service  int32
}

// GetPricingFees retrieve the ride FEES by a given city
func GetPricingFees(ID string) (*Fees, error) {
	conn := config.PricingPostgresConnect()
	defer conn.Close()

	fees := &Fees{ID: ID}
	err := conn.Select(fees)

	return fees, err
}

// PricingFeesToProto transforms the type Fees in a protobuf message
func PricingFeesToProto(fees *Fees) *pricingpb.GetPricingFeesByCityResponse {
	return &pricingpb.GetPricingFeesByCityResponse{
		Id:       fees.ID,
		Base:     fees.Base,
		Distance: fees.Distance,
		Dynamic:  fees.Dynamic,
		Minute:   fees.Minute,
		Service:  fees.Service,
	}
}
