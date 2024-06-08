package domain

import (
	"math"
)

// Fare calculation tiers based on distance
const (
	baseFare  = 400 // Up to 1km
	tier2Fare = 40  // 1km - 10km
	tier3Fare = 40  // Above 10km

	baseDistanceMultiplier  = 1000 // 1km
	tier2DistanceMultiplier = 400  // 1km - 10km
	tier3DistanceMultiplier = 350  // Above 10km

	minDistanceCap      = 1000.0
	maxDistanceCap      = 10000.0
	tier2maxDistanceCap = 9000.0
)

// CalculateFare calculates the taxi fare based on the distance traveled.
func CalculateFare(totalDistance float64) int {
	//Error escape
	if totalDistance <= 0 {
		return -1
	}

	if totalDistance <= minDistanceCap { // Base fare up to 1 km
		return baseFare
	}

	totalFare := baseFare
	remainingDistance := totalDistance - minDistanceCap

	// Tier 2: Up to 10 km, add 40 yen every 400 meters
	if totalDistance <= maxDistanceCap { // Check if it's exceed 3rd tier
		totalFare += int(math.Ceil(remainingDistance/tier2DistanceMultiplier)) * tier2Fare
	} else { // Exceed 10km
		tier3Distance := totalDistance - maxDistanceCap

		totalFare += int(math.Ceil(tier2maxDistanceCap/tier2DistanceMultiplier)) * tier2Fare
		totalFare += int(math.Ceil(tier3Distance/tier3DistanceMultiplier)) * tier3Fare
	}

	return totalFare
}
