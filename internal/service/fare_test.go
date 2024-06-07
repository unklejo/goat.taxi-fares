package service

import (
	"bytes"
	"strings"
	"xyz.taxi-fares/internal/repository"
	"xyz.taxi-fares/pkg/meter"
	"testing"
)

// Helper function to parse meter records from string input
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
		name    string
		input   string
		want    string
		wantErr bool
	}{
		// Valid inputs
		{
			name:    "Basic Input",
			input:   "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			want:    `1240 00:02:00.125 1141.2 660.3 00:03:00.100 1800.8 659.6 00:01:00.123 480.9 480.9 00:00:00.000 0.0 0.0`,
			wantErr: false,
		},
		// Error scenarios (invalid format, time order, gaps, etc.)
		{
            name:  "Invalid Input Format",
            input: "00:00:00 0.0\n00:01:00.123 480.9",
            want:    "", // No output for invalid format
            wantErr: true,
        },
        {
            name:  "Blank Line",
            input: "00:00:00.000 0.0\n\n00:01:00.123 480.9",
            want:    "",
            wantErr: true,
        },
        {
            name:  "Invalid Time Order",
            input: "00:01:00.123 480.9\n00:00:00.000 0.0",
            want:    "",
            wantErr: true,
        },
        {
            name:  "Time Gap Too Large",
            input: "00:00:00.000 0.0\n00:06:00.123 480.9", 
            want:    "",
            wantErr: true,
        },
        {
            name:  "Insufficient Data",
            input: "00:00:00.000 0.0", 
            want:    "",
            wantErr: true,
        },
        {
            name:  "Zero Total Distance",
            input: "00:00:00.000 0.0\n00:01:00.123 0.0",
            want:    "",
            wantErr: true,
        },
        {
            name:  "Invalid Distance Format",
            input: "00:00:00.000 0.0\n00:01:00.123 480.9a",
            want:    "",
            wantErr: true,
        }
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{} // Capture standard output

			// Create a mock MeterRepository that returns records from tt.input
			mockRepo := &mockMeterRepository{
				records: parseRecords(t, tt.input),
			}
			service := NewFareService(mockRepo)

			err := service.CalculateAndOutputFare(meter.NewReader(strings.NewReader(tt.input)))
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateAndOutputFare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := out.String(); got != tt.want {
				t.Errorf("CalculateAndOutputFare() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Mock MeterRepository implementation
type mockMeterRepository struct {
	records []meter.Record
}

func (m *mockMeterRepository) ReadRecords(reader meter.Reader) ([]meter.Record, error) {
	return m.records, nil
}
