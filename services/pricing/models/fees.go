package models

import (
	"github.com/vmihailenco/msgpack/v4"
	"github.com/omppye-tech/petssenger/services/pricing/config"
	proto "github.com/omppye-tech/petssenger/services/pricing/protos"
	"github.com/omppye-tech/petssenger/services/pricing/redis"
)

type Fees struct {
	ID       string
	Base     float32
	Distance float32
	Dynamic  float32
	Minute   float32
	Service  float32
}

func ProtoPricingFees(fees *Fees) *proto.GetPricingFeesByCityResponse {
	return &proto.GetPricingFeesByCityResponse{
		Id:       fees.ID,
		Base:     fees.Base,
		Distance: fees.Distance,
		Minute:   fees.Minute,
		Service:  fees.Service,
	}
}

func GetPricingFees(ID string) (*Fees, error) {
	fees := &Fees{}

	val, err := redis.Client.Get(ID).Bytes()
	if err == nil {
		if err := msgpack.Unmarshal(val, fees); err == nil {
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

func ProtoDynamicFees(fees *Fees) *proto.GetDynamicFeesByCityResponse {
	return &proto.GetDynamicFeesByCityResponse{
		Dynamic: fees.Dynamic,
	}
}

func GetDynamicFees(ID string) (*Fees, error) {
	fees := &Fees{}

	err := db.Model(fees).Column("dynamic").Where("id = ?", ID).Select()
	if err != nil {
		return nil, err
	}

	return fees, nil
}

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
