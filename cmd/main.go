package main

import (
	"fmt"
	"os"

	"github.com/unklejo/xyz.taxi-fares/internal/repository"
	"github.com/unklejo/xyz.taxi-fares/internal/service"
	"github.com/unklejo/xyz.taxi-fares/internal/usecase"
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

func main() {
	reader := meter.NewReader(os.Stdin)

	repo := repository.NewMeterRepository()
	fareService := service.NewFareService(repo)
	fareUseCase := usecase.NewCalculateAndOutputFareUseCase(*fareService)

	if err := fareUseCase.Execute(*reader, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
