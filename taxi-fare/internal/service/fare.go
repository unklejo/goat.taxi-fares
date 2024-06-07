package service

import (
	"fmt"
	"os"
	"sort"
	"taxi-fare/internal/domain"
	"taxi-fare/internal/repository"
	"taxi-fare/pkg/meter"
)

type FareService struct {
	repo repository.MeterRepository
}

func NewFareService(repo repository.MeterRepository) *FareService {
	return &FareService{repo: repo}
}

func (fs *FareService) CalculateAndOutputFare(reader meter.Reader) error {
	records, err := fs.repo.ReadRecords(reader)
	if err != nil {
		return err
	}

	if len(records) < 2 || records[len(records)-1].Distance == 0 {
		return fmt.Errorf("insufficient or invalid data")
	}

	for i := 1; i < len(records); i++ {
		records[i].DistanceDiff = records[i].Distance - records[i-1].Distance
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].DistanceDiff > records[j].DistanceDiff
	})

	finalFare := domain.CalculateFare(records[len(records)-1].Distance)
	fmt.Println(finalFare)

	for _, record := range records {
		fmt.Printf("%s %.1f %.1f\n", record.Time, record.Distance, record.DistanceDiff)
	}

	return nil
}
