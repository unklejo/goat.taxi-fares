package repository

import (
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

type MeterRepository struct{}

func NewMeterRepository() *MeterRepository {
	return &MeterRepository{}
}

func (r *MeterRepository) ReadRecords(reader *meter.Reader) ([]meter.Record, error) {
	records, err := reader.ReadRecords()
	if err != nil {
		return nil, err
	}
	return records, nil
}
