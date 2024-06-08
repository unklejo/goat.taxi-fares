package repository

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/unklejo/xyz.taxi-fares/pkg/meter"
)

func parseTime(timeStr string) time.Time {
	timeLayout := "15:04:05.000"
	parsedTime, _ := time.Parse(timeLayout, timeStr)
	return parsedTime
}

func TestMeterRepository_ReadRecords(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []meter.Record
		wantErr bool
	}{
		{
			name:  "Valid Input",
			input: "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			want: []meter.Record{
				{Time: parseTime("00:00:00.000"), Distance: 0.0},
				{Time: parseTime("00:01:00.123"), Distance: 480.9},
				{Time: parseTime("00:02:00.125"), Distance: 1141.2},
				{Time: parseTime("00:03:00.100"), Distance: 1800.8},
			},
		},
		{
			name:    "Invalid Input Format",
			input:   "00:00:00 0.0\n00:01:00.123 480.9",
			wantErr: true,
		},
		{
			name:    "Blank Line",
			input:   "00:00:00.000 0.0\n\n00:01:00.123 480.9",
			wantErr: true,
		},
		{
			name:    "Invalid Time Order",
			input:   "00:01:00.123 480.9\n00:00:00.000 0.0",
			wantErr: true,
		},
		{
			name:    "Time Gap Too Large",
			input:   "00:00:00.000 0.0\n00:06:00.123 480.9",
			wantErr: true,
		},
		{
			name:    "Insufficient Data",
			input:   "00:00:00.000 0.0",
			wantErr: true,
		},
		{
			name:    "Zero Total Distance",
			input:   "00:00:00.000 0.0\n00:01:00.123 0.0",
			wantErr: true,
		},
		{
			name:    "Invalid Distance Format",
			input:   "00:00:00.000 0.0\n00:01:00.123 480.9a",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMeterRepository()

			reader := meter.NewReader(strings.NewReader(tt.input))

			got, err := repo.ReadRecords(reader)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// We only need to verify the return records if no error is expected
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
