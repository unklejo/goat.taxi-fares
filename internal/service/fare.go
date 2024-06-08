package service

import (
	"fmt"
	"io"
	"log"
	"sort"

	"github.com/unklejo/xyz.taxi-fares/internal/domain"
	"github.com/unklejo/xyz.taxi-fares/internal/repository"
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

// FareService represents the fare calculation service.
type FareService struct {
	repo repository.MeterRepository
}

// NewFareService creates a new FareService instance.
func NewFareService(repo repository.MeterRepository) *FareService {
	return &FareService{repo: repo}
}

// CalculateAndOutputFare calculates and outputs the fare based on meter records.
func (fareService *FareService) CalculateAndOutputFare(reader *meter.Reader, w io.Writer) error {
	// Use the repository's ReadRecords to get the meter records
	records, err := fareService.repo.ReadRecords(reader)
	if err != nil {
		log.Printf("Error reading records: %v", err)
		return err
	}

	// Handle case where there are no records (i.e., empty input)
	if len(records) == 0 {
		return nil // No input, so no fare to calculate
	}

	if records[len(records)-1].Distance == 0 {
		return fmt.Errorf("invalid data: total distance is zero")
	}

	// Calculate distance differences
	for i := 1; i < len(records); i++ {
		records[i].DistanceDiff = records[i].Distance - records[i-1].Distance
	}

	// Sort by distance difference
	sort.Slice(records, func(i, j int) bool {
		return records[i].DistanceDiff > records[j].DistanceDiff
	})

	finalFare := domain.CalculateFare(records[len(records)-1].Distance)

	// Output to the provided io.Writer
	fmt.Fprintln(w, finalFare)
	for _, record := range records {
		fmt.Fprintf(w, "%s %.1f %.1f\n", record.Time, record.Distance, record.DistanceDiff)
	}

	return nil
}
