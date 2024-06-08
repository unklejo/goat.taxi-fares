package usecase

import (
	"github.com/unklejo/xyz.taxi-fares/internal/service"
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
	"io"
)

type CalculateAndOutputFareUseCase struct {
	fareService service.FareService
}

func NewCalculateAndOutputFareUseCase(fareService service.FareService) *CalculateAndOutputFareUseCase {
	return &CalculateAndOutputFareUseCase{
		fareService: fareService,
	}
}

func (uc *CalculateAndOutputFareUseCase) Execute(reader meter.Reader, w io.Writer) error {
	return uc.fareService.CalculateAndOutputFare(reader, w)
}
