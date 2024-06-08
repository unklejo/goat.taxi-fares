package domain

import (
	"math"
)

// Fare calculation tiers based on distance
const (
	baseFare           = 400 // Up to 1km
	tier2Fare          = 40  // 1km - 10km
	tier3Fare          = 40  // Above 10km
	baseDistanceLimit  = 1000.0
	tier2DistanceLimit = 10000.0
	tier2DistanceUnit  = 400.0
	tier3DistanceUnit  = 350.0
)

// CalculateFare calculates the taxi fare based on the distance traveled.
func CalculateFare(totalDistance float64) int {
	if totalDistance <= 0 {
		return -1
	}

	totalFare := baseFare

	if totalDistance <= baseDistanceLimit {
		return totalFare
	}

	remainingDistance := totalDistance - baseDistanceLimit

	if totalDistance <= tier2DistanceLimit {
		totalFare += int(math.Ceil(remainingDistance/tier2DistanceUnit)) * tier2Fare
		return totalFare
	}

	tier2Distance := tier2DistanceLimit - baseDistanceLimit
	tier3Distance := totalDistance - tier2DistanceLimit

	totalFare += int(math.Ceil(tier2Distance/tier2DistanceUnit)) * tier2Fare
	totalFare += int(math.Ceil(tier3Distance/tier3DistanceUnit)) * tier3Fare

	return totalFare
}
