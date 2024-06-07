package repository

import "taxi-fare-calculator/pkg/meter"

type MeterRepository interface {
	ReadRecords() ([]meter.Record, error)
}
