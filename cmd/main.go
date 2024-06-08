package main

import (
	"fmt"
	"os"

	"github.com/unklejo/xyz.taxi-fares/internal/repository"
	"github.com/unklejo/xyz.taxi-fares/internal/service"
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

func main() {
	repo := repository.NewMeterRepository()
	fareService := service.NewFareService(repo)

	reader := meter.NewReader(os.Stdin)

	err := fareService.CalculateAndOutputFare(reader, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
