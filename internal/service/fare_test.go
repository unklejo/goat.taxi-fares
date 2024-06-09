package service

import (
	"bytes"
	"strings"
	"testing"

	"github.com/unklejo/xyz.taxi-fares/internal/repository"
	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

func TestCalculateAndOutputFare(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedFare   string
		expectedOutput string
		wantErr        bool
	}{
		{
			name:           "Basic Input",
			input:          "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			expectedFare:   "520\n", // Calculated based on the domain logic
			expectedOutput: "00:02:00.125 1141.2 660.3\n00:03:00.100 1800.8 659.6\n00:01:00.123 480.9 480.9\n00:00:00.000 0.0 0.0\n",
			wantErr:        false,
		},
		{
			name:           "Short Distance",
			input:          "00:00:00.000 0.0\n00:00:30.000 500.0\n",
			expectedFare:   "400\n", // Base fare for distance <= 1000 meters
			expectedOutput: "00:00:30.000 500.0 500.0\n00:00:00.000 0.0 0.0\n",
			wantErr:        false,
		},
		{
			name:           "Long Distance",
			input:          "00:00:00.000 0.0\n00:05:00.000 5000.0\n00:10:00.000 12000.0\n",
			expectedFare:   "1560\n", // Calculated based on the domain logic
			expectedOutput: "00:10:00.000 12000.0 7000.0\n00:05:00.000 5000.0 5000.0\n00:00:00.000 0.0 0.0\n",
			wantErr:        false,
		},
		{
			name:           "Invalid Input Format",
			input:          "invalid\n00:01:00.123 480.9\n00:02:00.125 1141.2\n",
			expectedFare:   "",
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Blank Line",
			input:          "00:00:00.000 0.0\n\n00:01:00.123 480.9\n",
			expectedFare:   "",
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Past Time",
			input:          "00:01:00.000 0.0\n00:00:30.000 500.0\n",
			expectedFare:   "",
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Large Time Gap",
			input:          "00:00:00.000 0.0\n00:06:00.000 500.0\n",
			expectedFare:   "",
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Less Than Two Lines",
			input:          "00:00:00.000 0.0\n",
			expectedFare:   "",
			expectedOutput: "",
			wantErr:        true,
		},
		{
			name:           "Total Mileage Zero",
			input:          "00:00:00.000 0.0\n00:01:00.000 0.0\n",
			expectedFare:   "",
			expectedOutput: "",
			wantErr:        true,
		},
	}

	repo := repository.NewMeterRepository()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock meter reader
			reader := meter.NewReader(strings.NewReader(tt.input))

			// Create a buffer to capture output
			var buf bytes.Buffer

			// Create the fare service
			fareService := NewFareService(repo)

			// Call the CalculateAndOutputFare function with mock reader and buffer
			err := fareService.CalculateAndOutputFare(reader, &buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateAndOutputFare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			output := buf.String()

			// Split output into fare and records
			outputParts := strings.SplitN(output, "\n", 2)
			if len(outputParts) < 2 {
				if !tt.wantErr {
					t.Fatalf("Unexpected output format: %s", output)
				}
				return
			}

			fare := outputParts[0] + "\n"
			recordsOutput := outputParts[1]

			// Check if the fare matches the expected value
			if fare != tt.expectedFare {
				t.Errorf("Unexpected fare. Got: %s, Want: %s", fare, tt.expectedFare)
			}

			// Check if the output matches the expected value
			if recordsOutput != tt.expectedOutput {
				t.Errorf("Unexpected output. Got: %s, Want: %s", recordsOutput, tt.expectedOutput)
			}
		})
	}
}
