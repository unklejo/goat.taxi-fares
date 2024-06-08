package domain

import (
	"fmt"
	"testing"
)

func TestCalculateFare(t *testing.T) {
	tests := []struct {
		distance float64
		want     int
	}{
		// Examples from Notion
		// 00:00:00.000 0.0
		// 00:01:00.123 480.9
		// 00:02:00.125 1141.2
		// 00:03:00.100 1800.8

		{0, 400},      // No trajectory
		{500, 400},    // Under 1km: ¥400
		{1000, 400},   // Under 1km: ¥400
		{1400, 440},   // 1km price: ¥400 + (¥40 per 400) * 1
		{5000, 800},   // 1km price: ¥400 + (¥40 per 400) * 10
		{9999, 1320},  // 1km price: ¥400 + (¥40 per 400) * 23
		{10000, 1320}, // 1km price: ¥400 + (¥40 per 400) * 23
		{10350, 1360}, // 1km price: ¥400 + ((¥40 per 400) * 23) + ((¥40 per 350) *1)
		{15000, 1920}, // 1km price: ¥400 + ((¥40 per 400) * 23) + ((¥40 per 350) *15)
		{-100, -1},    // No trajectory
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("distance=%.0f", tt.distance), func(t *testing.T) {
			if got := CalculateFare(tt.distance); got != tt.want {
				t.Errorf("CalculateFare() = %v, want %v", got, tt.want)
			}
		})
	}
}
