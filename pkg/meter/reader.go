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
	Time         time.Time
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

		records = append(records, Record{Time: currentTime, Distance: distance})
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
