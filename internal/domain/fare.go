package domain

import "math"

// Fare calculation tiers based on distance
const (
	baseFare         = 400 // Up to 1km
	tier2FarePer400m = 40  // 1km - 10km
	tier3FarePer350m = 40  // Above 10km
)

// CalculateFare calculates the taxi fare based on the distance traveled.
func CalculateFare(distance float64) int {
	if distance <= 1000 { // Base fare up to 1 km
		return baseFare
	}

	fare := baseFare
	remainingDistance := distance - 1000

	// Tier 2: Up to 10 km, add 40 yen every 400 meters
	tier2Distance := math.Min(remainingDistance, 9000) // Maximum distance in tier 2 is 9000 meters
	fare += int(tier2Distance/400) * tier2FarePer400m
	remainingDistance -= tier2Distance

	// Tier 3: Above 10 km, add 40 yen every 350 meters
	fare += int(math.Ceil(remainingDistance/350)) * tier3FarePer350m

	return fare
}
