package repository

import "taxi-fare/pkg/meter"

type MeterRepository interface {
	ReadRecords() ([]meter.Record, error)
}
