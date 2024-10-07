package linearstats

import "math"

func PearsonCorrelationCoefficient(size float64, n ...float64) float64 {
	numerator := (size * n[2]) - (n[0] * n[1])
	denominator := ((size * n[3]) - (n[0] * n[0])) * ((size * n[4]) - (n[1] * n[1]))

	return numerator / math.Sqrt(denominator)
}
