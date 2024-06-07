package main

import (
	"fmt"
	"os"

	"taxi-fare-calculator/internal/usecase"
	"taxi-fare-calculator/pkg/meter"
)

func main() {
	reader := meter.NewReader(os.Stdin)
	fareUseCase := usecase.NewFareUseCase()

	if err := fareUseCase.CalculateAndOutputFare(reader); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
