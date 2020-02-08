package main

import (
	"log"
	"time"

	"github.com/weslenng/petssenger/services/pricing/models"
)

func main() {
	start := time.Now()
	fees, err := models.GetPricingFees("CURITIBA")
	if err != nil {
		panic("Error when getting pricing fees")
	}

	elapsed := time.Since(start)
	log.Printf("GetPricingFees [1] %s Returned %v", elapsed, fees)
	start = time.Now()

	fees, err = models.GetPricingFees("CURITIBA")
	if err != nil {
		panic("Error when getting pricing fees")
	}

	elapsed = time.Since(start)
	log.Printf("GetPricingFees [2] %s Returned %v", elapsed, fees)
	start = time.Now()

	fees, err = models.GetPricingFees("CURITIBA")
	if err != nil {
		panic("Error when getting pricing fees")
	}

	elapsed = time.Since(start)
	log.Printf("GetPricingFees [3] %s Returned %v", elapsed, fees)
}
