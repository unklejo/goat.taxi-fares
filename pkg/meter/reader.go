package meter

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	Time         string
	Distance     float64
	DistanceDiff float64
}

type Reader struct {
	scanner *bufio.Scanner
}

func NewReader(r io.Reader) *Reader {
	return &Reader{scanner: bufio.NewScanner(r)}
}

func (r *Reader) ReadRecords() ([]Record, error) {
	var records []Record
	var lastTime *time.Time
	const timeLayout = "15:04:05.000"

	recordRegex := regexp.MustCompile(`^(\d{2}:\d{2}:\d{2}\.\d{3})\s+(\d+\.?\d*)$`)
	for r.scanner.Scan() {
		line := r.scanner.Text()

		// Check for blank lines explicitly
		if strings.TrimSpace(line) == "" {
			return nil, fmt.Errorf("blank line encountered")
		}

		matches := recordRegex.FindStringSubmatch(line)
		if len(matches) != 3 { // 0: full match, 1: time, 2: distance
			return nil, fmt.Errorf("invalid input format: %s", line)
		}

		currentTime, err := time.Parse(timeLayout, matches[1])
		if err != nil {
			return nil, fmt.Errorf("invalid time format: %s", matches[1])
		}

		if lastTime != nil && currentTime.Before(*lastTime) {
			return nil, fmt.Errorf("invalid time order: %s", line)
		}

		if lastTime != nil && currentTime.Sub(*lastTime) > 5*time.Minute {
			return nil, fmt.Errorf("time gap too large: %s", line)
		}

		distance, err := strconv.ParseFloat(matches[2], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid distance format: %s", matches[2])
		}

		records = append(records, Record{Time: matches[1], Distance: distance})
		lastTime = &currentTime
	}

	if err := r.scanner.Err(); err != nil {
		return nil, err
	}

	if len(records) < 2 || records[len(records)-1].Distance == 0.0 {
		return nil, fmt.Errorf("insufficient or invalid data")
	}

	return records, nil
}

// parseFloat64 converts a string to a float64, ensuring it has at most one decimal point.
func parseFloat64(s string) (float64, error) {
	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return 0, fmt.Errorf("invalid floating-point format: %s", s)
	}
	whole, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	if len(parts) == 1 {
		return float64(whole), nil
	}
	fracStr := parts[1]
	if len(fracStr) > 1 {
		return 0, fmt.Errorf("too many digits after decimal point: %s", s)
	}
	frac, err := strconv.Atoi(fracStr)
	if err != nil {
		return 0, err
	}
	return float64(whole) + float64(frac)/10, nil
}
