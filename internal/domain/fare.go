package domain

func CalculateFare(distance float64) int {
	if distance <= 1000 {
		return 400
	}

	fare := 400
	remainingDistance := distance - 1000

	fare += int(min(remainingDistance, 9000)/400) * 40
	remainingDistance -= min(remainingDistance, 9000)

	fare += int(ceil(remainingDistance/350)) * 40

	return fare
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func ceil(a float64) float64 {
	return float64(int(a) + 1)
}
