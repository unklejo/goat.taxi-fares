package service

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/unklejo/xyz.taxi-fares/internal/repository"
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

const timeLayout = "15:04:05.000"

func parseRecords(t *testing.T, input string) []meter.Record {
	records := make([]meter.Record, 0) // Initialize empty slice
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) != 2 {
			t.Fatalf("Invalid record format: %s", line)
		}

		timeStr := fields[0]
		distanceStr := fields[1]

		_, err := time.Parse(timeLayout, timeStr)
		if err != nil {
			t.Fatalf("Invalid time format: %s", timeStr)
		}

		distance, err := strconv.ParseFloat(distanceStr, 64)
		if err != nil {
			t.Fatalf("Invalid distance format: %s", distanceStr)
		}

		records = append(records, meter.Record{Time: timeStr, Distance: distance})
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Error scanning input: %v", err)
	}

	return records
}

func TestFareService_CalculateAndOutputFare(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		mockRepo       repository.MeterRepository
		expectedOutput string
		wantErr        bool
	}{
		{
			name:  "Basic Input",
			input: "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			mockRepo: &mockMeterReader{
				records: []meter.Record{
					{Time: "00:00:00.000", Distance: 0.0},
					{Time: "00:01:00.123", Distance: 480.9},
					{Time: "00:02:00.125", Distance: 1141.2},
					{Time: "00:03:00.100", Distance: 1800.8},
				},
				err: nil,
			},
			expectedOutput: `1240 00:02:00.125 1141.2 660.3 00:03:00.100 1800.8 659.6 00:01:00.123 480.9 480.9 00:00:00.000 0.0 0.0`,
			wantErr:        false,
		},
		// Error scenarios (invalid format, time order, gaps, etc.)
		{
			name:           "Invalid Input Format",
			input:          "00:00:00 0.0\n00:01:00.123 480.9",
			mockRepo:       &mockMeterReader{records: nil, err: fmt.Errorf("invalid input format")},
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Blank Line",
			input:          "00:00:00.000 0.0\n\n00:01:00.123 480.9",
			mockRepo:       &mockMeterReader{records: nil, err: fmt.Errorf("blank line encountered")},
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Invalid Time Order",
			input:          "00:01:00.123 480.9\n00:00:00.000 0.0",
			mockRepo:       &mockMeterReader{records: nil, err: fmt.Errorf("invalid time order")},
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Time Gap Too Large",
			input:          "00:00:00.000 0.0\n00:06:00.123 480.9",
			mockRepo:       &mockMeterReader{records: nil, err: fmt.Errorf("time gap too large")},
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Insufficient Data",
			input:          "00:00:00.000 0.0",
			mockRepo:       &mockMeterReader{records: nil, err: fmt.Errorf("insufficient or invalid data")},
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Zero Total Distance",
			input:          "00:00:00.000 0.0\n00:01:00.123 0.0",
			mockRepo:       &mockMeterReader{records: nil, err: fmt.Errorf("invalid data: total distance is zero")},
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Invalid Distance Format",
			input:          "00:00:00.000 0.0\n00:01:00.123 480.9a",
			mockRepo:       &mockMeterReader{records: nil, err: fmt.Errorf("invalid distance format: 480.9a")},
			expectedOutput: "",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}

			mockRepo := &mockMeterReader{
				records: parseRecords(t, tt.input),
			}
			service := NewFareService(mockRepo)

			// Dereference the reader
			err := service.CalculateAndOutputFare(*meter.NewReader(strings.NewReader(tt.input)), out)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateAndOutputFare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			expectedOutput := strings.ReplaceAll(tt.expectedOutput, " ", "\n")
			if got := out.String(); got != expectedOutput {
				t.Errorf("CalculateAndOutputFare() = %v, want %v", got, expectedOutput)
			}
		})
	}
}

// Mock MeterRepository implementation
type mockMeterReader struct {
	records []meter.Record
	err     error
}

// Implement ReadRecords to satisfy the MeterRepository interface
func (m *mockMeterReader) ReadRecords(reader meter.Reader) ([]meter.Record, error) {
	return m.records, m.err
}
