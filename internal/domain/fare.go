package domain

import (
	"math"
)

// Fare calculation tiers based on distance
const (
	baseFare           = 400 // Up to 1km
	tier2Fare          = 40  // 1km - 10km
	tier3Fare          = 40  // Above 10km
	baseDistanceRange  = 1000.0
	tier2DistanceRange = 9000.0
	tier2DistanceUnit  = 400.0
	tier3DistanceUnit  = 350.0
)

// CalculateFare calculates the fare based on distance.
func CalculateFare(distance float64) int {
	if distance <= 0 { // Check if distance zero or negative
		return baseFare
	}

	distance -= baseDistanceRange

	if distance <= 0 { // Check if distance exceed tier 2
		return baseFare
	}

	totalFare := baseFare
	if distance > (tier2DistanceRange) { // Exceed distance of tier 3
		totalFare += int(math.Ceil(tier2DistanceRange/tier2DistanceUnit)) * tier2Fare
		distance -= tier2DistanceRange
		totalFare += int(math.Ceil(distance/tier3DistanceUnit)) * tier3Fare
	} else { // Distance in range of tier 2
		totalFare += int(math.Ceil(distance/tier2DistanceUnit)) * tier2Fare
	}

	return totalFare
}
