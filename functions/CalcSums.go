package linearstats

import "math"

func CalcSums(data []float64) (float64, float64, float64, float64, float64) {
	Sx := float64(0)
	Sy := float64(0)
	Sxy := float64(0)
	P_Sx := float64(0)
	P_Sy := float64(0)
	for i := 0; i < len(data); i++ {
		Sx += float64(i)
		Sy += data[i]
		Sxy += float64(i) * data[i]
		P_Sx += float64(math.Pow(float64(i), 2))
		P_Sy += math.Pow(data[i], 2)
	}
	return Sx, Sy, Sxy, P_Sx, P_Sy
}
