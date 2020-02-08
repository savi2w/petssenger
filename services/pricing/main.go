package main

import (
	"github.com/weslenng/petssenger/services/pricing/server"
)

func main() {
	lis, ser, err := server.PricingServerListen()
	if err != nil {
		panic(err)
	}

	defer lis.Close()
	defer ser.GracefulStop()
}
