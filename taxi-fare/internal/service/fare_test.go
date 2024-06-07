package service

import (
	"bytes"
	"strings"
	"taxi-fare-calculator/internal/repository"
	"taxi-fare-calculator/pkg/meter"
	"testing"
)

func TestFareService_CalculateAndOutputFare(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		// Valid inputs
		{
			name:  "Basic Input",
			input: "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			want: `1240
00:02:00.125 1141.2 660.3
00:03:00.100 1800.8 659.6
00:01:00.123 480.9 480.9
00:00:00.000 0.0 0.0
`,
			wantErr: false,
		},
		// Error scenarios (invalid format, time order, gaps, etc.)
		// ... (add more test cases to cover error conditions)
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

// Helper function to parse meter records from string input
func parseRecords(t *testing.T, input string) []meter.Record {
	// ... (implement parsing logic)
}

// Mock MeterRepository implementation
type mockMeterRepository struct {
	records []meter.Record
}

func (m *mockMeterRepository) ReadRecords(reader meter.Reader) ([]meter.Record, error) {
	return m.records, nil
}
