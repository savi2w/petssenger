package models

import (
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack/v4"
	pricingpb "github.com/weslenng/petssenger/protos"
	"github.com/weslenng/petssenger/services/pricing/config"
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
func ProtoPricingFees(fees *Fees) *pricingpb.GetPricingFeesByCityResponse {
	return &pricingpb.GetPricingFeesByCityResponse{
		Id:       fees.ID,
		Base:     fees.Base,
		Distance: fees.Distance,
		Minute:   fees.Minute,
		Service:  fees.Service,
	}
}

// GetPricingFees retrieve the ride fees by a given city
func GetPricingFees(ID string, pg *pg.DB, redis *redis.Client) (*Fees, error) {
	fees := &Fees{}

	val, err := redis.Get(ID).Bytes()
	if err == nil {
		err = msgpack.Unmarshal(val, fees)
		if err == nil {
			return fees, nil
		}
	}

	err = pg.Model(fees).Column(
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

	err = redis.Set(ID, val, 5*time.Minute).Err()
	if err != nil {
		return nil, err
	}

	return fees, nil
}

// ProtoDynamicFees transforms the type Fees in a GetDynamicFeesByCityResponse
func ProtoDynamicFees(fees *Fees) *pricingpb.GetDynamicFeesByCityResponse {
	return &pricingpb.GetDynamicFeesByCityResponse{
		Dynamic: fees.Dynamic,
	}
}

// GetDynamicFees retrieve the dynamic fees by a given city
func GetDynamicFees(ID string, pg *pg.DB) (*Fees, error) {
	fees := &Fees{}

	err := pg.Model(fees).Column("dynamic").Where("id = ?", ID).Select()
	if err != nil {
		return nil, err
	}

	return fees, nil
}

const variation = 0.1

// IncreaseDynamicFees increase a dynamic fees by a given city
func IncreaseDynamicFees(ID string, pg *pg.DB) error {
	fees := &Fees{}

	_, err := pg.Model(fees).Set("dynamic = dynamic + ?", variation).Where("id = ?", ID).Update()
	if err != nil {
		return err
	}

	return nil
}

const minimal = 1

// DecreaseDynamicFees decrease a dynamic fees by a given city (used in worker)
func DecreaseDynamicFees(ID string) error {
	fees := &Fees{}

	pg := config.PricingPgConnect()
	defer pg.Close()

	_, err := pg.Model(fees).Set("dynamic = dynamic - ?", variation).Where("id = ? AND dynamic > ?", ID, minimal).Update()
	if err != nil {
		return err
	}

	return nil
}
