package usecase

import "xyz.taxi-fares/pkg/meter"

type FareUseCase interface {
	CalculateAndOutputFare(reader meter.Reader) error
}
