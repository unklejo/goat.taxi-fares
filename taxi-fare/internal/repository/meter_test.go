package repository

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"taxi-fare/pkg/meter"
)

type mockMeterReader struct {
	records []meter.Record
	err     error
}

func (m *mockMeterReader) ReadRecords() ([]meter.Record, error) {
	return m.records, m.err
}

func TestMeterRepository_ReadRecords(t *testing.T) {
	tests := []struct {
		name      string
		reader    meter.Reader
		want      []meter.Record
		wantErr   bool
		errorText string // Added to check for specific error messages
	}{
		{
			name: "Valid Input",
			reader: &mockMeterReader{
				records: []meter.Record{
					{Time: "00:00:00.000", Distance: 0.0},
					{Time: "00:01:00.123", Distance: 480.9},
					{Time: "00:02:00.125", Distance: 1141.2},
					{Time: "00:03:00.100", Distance: 1800.8},
				},
			},
			want: []meter.Record{
				{Time: "00:00:00.000", Distance: 0.0},
				{Time: "00:01:00.123", Distance: 480.9},
				{Time: "00:02:00.125", Distance: 1141.2},
				{Time: "00:03:00.100", Distance: 1800.8},
			},
			wantErr:   false,
			errorText: "",
		},
		{
			name:      "Invalid Input Format",
			reader:    &mockMeterReader{err: fmt.Errorf("invalid input format: 00:00:00 0.0")},
			want:      nil,
			wantErr:   true,
			errorText: "invalid input format: 00:00:00 0.0",
		},
		{
			name:      "Invalid Time Order",
			reader:    &mockMeterReader{err: fmt.Errorf("invalid time order: 00:01:00.000 0.0")},
			want:      nil,
			wantErr:   true,
			errorText: "invalid time order: 00:01:00.000 0.0",
		},
		{
			name:      "Time Gap Too Large",
			reader:    &mockMeterReader{err: fmt.Errorf("time gap too large: 00:10:00.000 0.0")},
			want:      nil,
			wantErr:   true,
			errorText: "time gap too large: 00:10:00.000 0.0",
		},
		{
			name:      "Insufficient Data",
			reader:    &mockMeterReader{records: []meter.Record{}}, // Empty slice
			want:      nil,
			wantErr:   true,
			errorText: "insufficient or invalid data",
		},
		{
			name: "Zero Total Distance",
			reader: &mockMeterReader{
				records: []meter.Record{
					{Time: "00:00:00.000", Distance: 0.0},
					{Time: "00:01:00.123", Distance: 0.0},
				},
			},
			want:      nil,
			wantErr:   true,
			errorText: "insufficient or invalid data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewMeterRepository()
			got, err := repo.ReadRecords(tt.reader)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errorText {
				t.Errorf("ReadRecords() error message = %v, want %v", err.Error(), tt.errorText)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
