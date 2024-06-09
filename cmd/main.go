package main

import (
	"fmt"
	"log"
	"os"

	"github.com/unklejo/xyz.taxi-fares/internal/repository"
	"github.com/unklejo/xyz.taxi-fares/internal/service"
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

func main() {
	log.Println("Starting taxi fare calculator application...")

	repo := repository.NewMeterRepository()
	fareService := service.NewFareService(repo)

	reader := meter.NewReader(os.Stdin)

	err := fareService.CalculateAndOutputFare(reader, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		log.Fatalf("Error calculating fare: %v", err)
		os.Exit(1)
	}

	log.Println("Calculation and output completed successfully.")
}
