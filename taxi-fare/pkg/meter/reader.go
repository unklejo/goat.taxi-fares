package meter

import (
    "bufio"
    "fmt"
    "io"
    "regexp"
    "strings"
    "time"
)

type Record struct {
    Time          string
    Distance      float64
    DistanceDiff  float64
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

    recordRegex := regexp.MustCompile(`^(\d{2}:\d{2}:\d{2}\.\d{3})\s+(\d+\.?\d*)$`)
    for r.scanner.Scan() {
        line := r.scanner.Text()
        matches := recordRegex.FindStringSubmatch(line)
        if len(matches) != 3 { // 0: full match, 1: time, 2: distance
            return nil, fmt.Errorf("invalid input format: %s", line)
        }
        
        // ... (rest of the logic for parsing time, distance, and error handling)
}
