package meter

import (
	"reflect"
	"strings"
	"testing"
)

func TestReader_ReadRecords(t *testing.T) {
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
				{Time: "00:00:00.000", Distance: 0.0},
				{Time: "00:01:00.123", Distance: 480.9},
				{Time: "00:02:00.125", Distance: 1141.2},
				{Time: "00:03:00.100", Distance: 1800.8},
			},
		},
		// Add more test cases here
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
