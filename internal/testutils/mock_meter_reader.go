package testutils

import "github.com/unklejo/xyz.taxi-fares/pkg/meter"

// MockMeterReader is a mock implementation of the MeterRepository interface
type MockMeterReader struct {
	Records []meter.Record
	Err     error
}

func (m *MockMeterReader) ReadRecords(reader *meter.Reader) ([]meter.Record, error) {
	return m.Records, m.Err
}

// Add this getter method to expose the records field
func (m *MockMeterReader) GetRecords() []meter.Record {
	return m.Records
}
