package service

import (
	"fmt"
	"io"
	"log"
	"sort"

	"github.com/unklejo/xyz.taxi-fares/internal/domain"
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

const (
	timeLayout = "15:04:05.000"
)

type MeterRepository interface {
	ReadRecords(reader *meter.Reader) ([]meter.Record, error)
}

type FareService struct {
	meterRepo MeterRepository
}

func NewFareService(meterRepo MeterRepository) *FareService {
	return &FareService{meterRepo: meterRepo}
}

func (s *FareService) CalculateAndOutputFare(reader *meter.Reader, output io.Writer) error {
	records, err := s.meterRepo.ReadRecords(reader)
	if err != nil {
		log.Println("Error reading records:", err)
		return err
	}

	if len(records) < 2 {
		log.Println("Error reading records: insufficient or invalid data")
		return fmt.Errorf("insufficient or invalid data")
	}

	totalDistance := records[len(records)-1].Distance
	if totalDistance == 0 {
		log.Println("Error reading records: insufficient or invalid data")
		return fmt.Errorf("insufficient or invalid data")
	}

	fare := domain.CalculateFare(totalDistance)
	if fare == -1 {
		log.Println("Error calculating fare")
		return fmt.Errorf("error calculating fare")
	}

	_, err = fmt.Fprintf(output, "%d\n", fare)
	if err != nil {
		return err
	}

	// Calculate the mileage difference and sort in descending order
	type outputRecord struct {
		time     string
		distance float64
		diff     float64
	}

	var outputRecords []outputRecord
	for i := 0; i < len(records); i++ {
		var diff float64
		if i == 0 {
			diff = 0
		} else {
			diff = records[i].Distance - records[i-1].Distance
		}
		outputRecords = append(outputRecords, outputRecord{
			time:     records[i].Time.Format(timeLayout),
			distance: records[i].Distance,
			diff:     diff,
		})
	}

	sort.Slice(outputRecords, func(i, j int) bool {
		return outputRecords[i].diff > outputRecords[j].diff
	})

	for _, rec := range outputRecords {
		_, err := fmt.Fprintf(output, "%s %.1f %.1f\n", rec.time, rec.distance, rec.diff)
		if err != nil {
			return err
		}
	}

	return nil
}
