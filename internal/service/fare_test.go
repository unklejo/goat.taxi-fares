package service

import (
	"bytes"
	"strings"
	"testing"

	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

type mockMeterRepository struct{}

func (m *mockMeterRepository) ReadRecords(reader *meter.Reader) ([]meter.Record, error) {
	return reader.ReadRecords()
}

func TestFareService_CalculateAndOutputFare(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedError  bool
		expectedOutput string
	}{
		{
			name:           "Basic Input",
			input:          "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			expectedError:  false,
			expectedOutput: "660\n00:02:00.125 1141.2 660.3\n00:03:00.100 1800.8 659.6\n00:01:00.123 480.9 480.9\n00:00:00.000 0.0 0.0\n",
		},
		{
			name:           "Short Distance",
			input:          "00:00:00.000 0.0\n00:00:30.000 500.0\n",
			expectedError:  false,
			expectedOutput: "400\n00:00:30.000 500.0 500.0\n00:00:00.000 0.0 0.0\n",
		},
		{
			name:           "Long Distance",
			input:          "00:00:00.000 0.0\n00:05:00.000 5000.0\n00:10:00.000 12000.0\n",
			expectedError:  false,
			expectedOutput: "1400\n00:10:00.000 12000.0 7000.0\n00:05:00.000 5000.0 5000.0\n00:00:00.000 0.0 0.0\n",
		},
		{
			name:          "Invalid Input Format",
			input:         "invalid\n",
			expectedError: true,
		},
		{
			name:          "Blank Line",
			input:         "00:00:00.000 0.0\n\n00:01:00.000 500.0\n",
			expectedError: true,
		},
		{
			name:          "Past Time",
			input:         "00:01:00.000 500.0\n00:00:00.000 0.0\n",
			expectedError: true,
		},
		{
			name:          "Large Time Gap",
			input:         "00:00:00.000 0.0\n00:06:00.000 500.0\n",
			expectedError: true,
		},
		{
			name:          "Less Than Two Lines",
			input:         "00:00:00.000 0.0\n",
			expectedError: true,
		},
		{
			name:          "Total Mileage Zero",
			input:         "00:00:00.000 0.0\n00:01:00.000 0.0\n",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mockMeterRepository{}
			fareService := NewFareService(mockRepo)
			reader := meter.NewReader(strings.NewReader(tt.input))
			var output bytes.Buffer

			err := fareService.CalculateAndOutputFare(reader, &output)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if !tt.expectedError && output.String() != tt.expectedOutput {
				t.Errorf("expected output: %q, got: %q", tt.expectedOutput, output.String())
			}
		})
	}
}
