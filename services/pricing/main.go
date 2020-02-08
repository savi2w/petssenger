package main

import (
	"fmt"

	"github.com/weslenng/petssenger/services/pricing/models"
)

func main() {
	fees, err := models.GetPricingFees("SAO_PAULO")
	if err != nil {
		panic("Error when getting pricing fees")
	}

	fmt.Printf("This is the pricing fees: %v", fees)
}
