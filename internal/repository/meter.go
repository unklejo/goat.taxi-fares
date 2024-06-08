package repository

import (
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

type MeterRepository interface {
	ReadRecords(reader *meter.Reader) ([]meter.Record, error)
}

type meterRepository struct{}

func NewMeterRepository() MeterRepository {
	return &meterRepository{}
}

func (mr *meterRepository) ReadRecords(reader *meter.Reader) ([]meter.Record, error) {
	return reader.ReadRecords()
}
