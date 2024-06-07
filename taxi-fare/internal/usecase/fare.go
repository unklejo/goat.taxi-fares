package usecase

import "taxi-fare-calculator/pkg/meter"

type FareUseCase interface {
	CalculateAndOutputFare(reader meter.Reader) error
}
