package meter

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestReaderReadRecords(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []Record
		wantErr bool
	}{
		{
			name:  "Valid Input",
			input: "00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			want: []Record{
				{Time: parseTime("00:00:00.000"), Distance: 0.0},
				{Time: parseTime("00:01:00.123"), Distance: 480.9},
				{Time: parseTime("00:02:00.125"), Distance: 1141.2},
				{Time: parseTime("00:03:00.100"), Distance: 1800.8},
			},
		},
		{
			name:    "Invalid Input Format",
			input:   "00:00:00.000 0.0\n00:01:00.123480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			wantErr: true,
		},
		{
			name:    "Invalid Time Order",
			input:   "00:01:00.123 480.9\n00:00:00.000 0.0\n00:02:00.125 1141.2\n00:03:00.100 1800.8\n",
			wantErr: true,
		},
		{
			name:    "Large Time Gap",
			input:   "00:00:00.000 0.0\n00:05:00.000 500.0\n00:12:00.000 1200.0\n00:13:00.000 1500.0\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := NewReader(strings.NewReader(tt.input))
			got, err := reader.ReadRecords()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parseTime(timeStr string) time.Time {
	timeLayout := "15:04:05.000"
	parsedTime, _ := time.Parse(timeLayout, timeStr)
	return parsedTime
}
