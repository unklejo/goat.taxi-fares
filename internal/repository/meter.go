package repository

import "xyz.taxi-fares/pkg/meter"

type MeterRepository interface {
	ReadRecords() ([]meter.Record, error)
}
