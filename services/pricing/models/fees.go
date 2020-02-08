package models

import (
	"time"

	"github.com/vmihailenco/msgpack/v4"
	pricingpb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/pricing/config"
)

// Fees payload structure
type Fees struct {
	ID       string
	Base     int32
	Distance int32
	Minute   int32
	Service  int32
}

// ProtoPricingFees transforms the type Fees in a protobuf message
func ProtoPricingFees(fees *Fees) *pricingpb.GetPricingFeesByCityResponse {
	return &pricingpb.GetPricingFeesByCityResponse{
		Id:       fees.ID,
		Base:     fees.Base,
		Distance: fees.Distance,
		Minute:   fees.Minute,
		Service:  fees.Service,
	}
}

// GetPricingFees retrieve the ride FEES by a given city
func GetPricingFees(ID string) (*Fees, error) {
	fees := &Fees{ID: ID}

	redis := config.PricingRedisClient()
	defer redis.Close()

	val, err := redis.Get(ID).Bytes()
	if err == nil {
		err = msgpack.Unmarshal(val, fees)
		if err == nil {
			return fees, nil
		}
	}

	postgres := config.PricingPostgresConnect()
	defer postgres.Close()

	err = postgres.Select(fees)
	if err != nil {
		panic(err)
	}

	val, err = msgpack.Marshal(fees)
	if err != nil {
		panic(err)
	}

	err = redis.Set(ID, val, 1*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	return fees, err
}
