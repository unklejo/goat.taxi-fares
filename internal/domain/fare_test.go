ackage domain

import "testing"

func TestCalculateFare(t *testing.T) {
    tests := []struct {
        distance float64
        want     int
    }{
        {0, 400},
        {500, 400},
        {1000, 400},
        {1400, 440},
        {5000, 1200},
        {9999, 2720},
        {10000, 2720},
        {10350, 2760},
        {15000, 3520},
    }
    for _, tt := range tests {
        t.Run(fmt.Sprintf("distance=%.0f", tt.distance), func(t *testing.T) {
            if got := CalculateFare(tt.distance); got != tt.want {
                t.Errorf("CalculateFare() = %v, want %v", got, tt.want)
            }
        })
    }
}
