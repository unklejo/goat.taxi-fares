Go
package repository

import (
	"bytes"
	"strings"
	"testing"
	"taxi-fare-calculator/pkg/meter"
)

type mockMeterReader struct {
	input string
}

func (m *mockMeterReader) ReadRecords() ([]meter.Record, error) {
	// ... (implement to return records based on m.input)
}

func TestMeterRepository_ReadRecords(t *testing.T) {
    // Test case for valid input
    validInput := "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n"
    expectedRecords := []meter.Record{
        // ... (your expected Record values here)
    }
    mockReader := &mockMeterReader{records: expectedRecords}
    repo := &MeterRepositoryImpl{} // Your actual MeterRepository implementation
    records, err := repo.ReadRecords(mockReader)
    if err != nil {
        t.Errorf("ReadRecords() error = %v", err)
        return
    }
    if !reflect.DeepEqual(records, expectedRecords) {
        t.Errorf("ReadRecords() = %v, want %v", records, expectedRecords)
    }
}
