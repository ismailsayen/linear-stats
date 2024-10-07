package linearstats

import "math"

func CalcSlope(data []float64, size float64, nbrs ...float64) float64 {
	numerator := (size * nbrs[2]) - (nbrs[0] * nbrs[1])
	denominator := (size * nbrs[3]) - math.Pow(nbrs[0], 2)
	return numerator / denominator
}
