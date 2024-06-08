package service

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unklejo/xyz.taxi-fares/internal/repository"
)

func TestFareService_CalculateAndOutputFare(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedFare   string
		expectedOutput string
	}{
		{
			name:           "Basic Input",
			input:          "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			expectedFare:   "660",
			expectedOutput: "00:02:00.125 1141.2 660.3\n00:03:00.100 1800.8 659.6\n00:01:00.123 480.9 480.9\n00:00:00.000 0.0 0.0\n",
		},
		{
			name:           "Short Distance",
			input:          "00:00:00.000 0.0\n00:00:30.000 500.0\n",
			expectedFare:   "400",
			expectedOutput: "00:00:30.000 500.0 500.0\n00:00:00.000 0.0 0.0\n",
		},
		{
			name:           "Long Distance",
			input:          "00:00:00.000 0.0\n00:05:00.000 5000.0\n00:10:00.000 12000.0\n",
			expectedFare:   "1400",
			expectedOutput: "00:10:00.000 12000.0 7000.0\n00:05:00.000 5000.0 5000.0\n00:00:00.000 0.0 0.0\n",
		},
		{
			name:           "Invalid Input Format",
			input:          "invalid\n",
			expectedFare:   "",
			expectedOutput: "",
		},
		{
			name:           "Blank Line",
			input:          "00:00:00.000 0.0\n\n00:01:00.000 500.0\n",
			expectedFare:   "",
			expectedOutput: "",
		},
		{
			name:           "Past Time",
			input:          "00:01:00.000 500.0\n00:00:00.000 0.0\n",
			expectedFare:   "",
			expectedOutput: "",
		},
		{
			name:           "Large Time Gap",
			input:          "00:00:00.000 0.0\n00:06:00.000 500.0\n",
			expectedFare:   "",
			expectedOutput: "",
		},
		{
			name:           "Less Than Two Lines",
			input:          "00:00:00.000 0.0\n",
			expectedFare:   "",
			expectedOutput: "",
		},
		{
			name:           "Total Mileage Zero",
			input:          "00:00:00.000 0.0\n00:01:00.000 0.0\n",
			expectedFare:   "",
			expectedOutput: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewMeterRepository()
			records, err := repo.ReadRecords(tt.input)
			if tt.expectedFare == "" {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			var sb strings.Builder
			sb.WriteString(tt.expectedFare + "\n")
			for _, record := range records {
				sb.WriteString(record.Time.Format("15:04:05.000") + " ")
				sb.WriteString(fmt.Sprintf("%.1f", record.Distance) + " ")
				sb.WriteString(fmt.Sprintf("%.1f", record.DistanceDiff) + "\n")
			}
			assert.Equal(t, tt.expectedOutput, sb.String())
		})
	}
}
