package main

import (
	"fmt"
	"os"

	"internal/usecase"
	"pkg/meter"
)

func main() {
	reader := meter.NewReader(os.Stdin)
	fareUseCase := usecase.NewFareUseCase()

	if err := fareUseCase.CalculateAndOutputFare(reader); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
