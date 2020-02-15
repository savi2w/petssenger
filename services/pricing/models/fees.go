package models

import (
	"github.com/vmihailenco/msgpack/v4"
	pb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/pricing/config"
	"github.com/weslenng/petssenger/services/pricing/redis"
)

// Fees represents a city-fees payload structure
type Fees struct {
	ID       string
	Base     float32
	Distance float32
	Dynamic  float32
	Minute   float32
	Service  float32
}

// ProtoPricingFees transforms the type Fees in a GetPricingFeesByCityResponse
func ProtoPricingFees(fees *Fees) *pb.GetPricingFeesByCityResponse {
	return &pb.GetPricingFeesByCityResponse{
		Id:       fees.ID,
		Base:     fees.Base,
		Distance: fees.Distance,
		Minute:   fees.Minute,
		Service:  fees.Service,
	}
}

// GetPricingFees retrieve the ride fees by a given city
func GetPricingFees(ID string) (*Fees, error) {
	fees := &Fees{}

	val, err := redis.Client.Get(ID).Bytes()
	if err == nil {
		err = msgpack.Unmarshal(val, fees)
		if err == nil {
			return fees, nil
		}
	}

	err = db.Model(fees).Column(
		"id",
		"base",
		"distance",
		"minute",
		"service",
	).Where("id = ?", ID).Select()

	if err != nil {
		return nil, err
	}

	val, err = msgpack.Marshal(fees)
	if err != nil {
		return nil, err
	}

	err = redis.Client.Set(ID, val, config.Default.RedisExpTime).Err()
	if err != nil {
		return nil, err
	}

	return fees, nil
}

// ProtoDynamicFees transforms the type Fees in a GetDynamicFeesByCityResponse
func ProtoDynamicFees(fees *Fees) *pb.GetDynamicFeesByCityResponse {
	return &pb.GetDynamicFeesByCityResponse{
		Dynamic: fees.Dynamic,
	}
}

// GetDynamicFees retrieve the dynamic fees by a given city
func GetDynamicFees(ID string) (*Fees, error) {
	fees := &Fees{}

	err := db.Model(fees).Column("dynamic").Where("id = ?", ID).Select()
	if err != nil {
		return nil, err
	}

	return fees, nil
}

// IncreaseDynamicFees increase a dynamic fees by a given city
func IncreaseDynamicFees(ID string) error {
	fees := &Fees{}

	_, err := db.Model(fees).Set(
		"dynamic = dynamic + ?",
		config.Default.DynamicFeesIncreaseRate,
	).Where("id = ?", ID).Update()

	if err != nil {
		return err
	}

	return nil
}

// DecreaseDynamicFees decrease a dynamic fees by a given city (used in worker)
func DecreaseDynamicFees(ID string) error {
	fees := &Fees{}

	_, err := db.Model(fees).Set(
		"dynamic = dynamic - ?",
		config.Default.DynamicFeesDecreaseRate,
	).Where(
		"id = ? AND dynamic > ?",
		ID,
		config.Default.DynamicFeesMinimumValue,
	).Update()

	if err != nil {
		return err
	}

	return nil
}
