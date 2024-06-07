package usecase

import "taxi-fare/pkg/meter"

type FareUseCase interface {
	CalculateAndOutputFare(reader meter.Reader) error
}
