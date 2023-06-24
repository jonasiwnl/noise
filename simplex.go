package noise

import "math"

// -Internal-----------------------------------------------------------------

// Skew and Simplicial subdivision
func skewAndSS(length int) *[]float64 {
	sum := 0.0
	icoords := make([]float64, length)

	for _, val := range icoords {
		sum += val
		val += (math.Sqrt(val+1.0) - 1) / val
	}

	for _, val := range icoords {
		val += sum
	}

	return &icoords
}

func gradient() *[]float64 {
	return &[]float64{0.0}
}

func kernelSummation() *[]float64 {
	return &[]float64{0.0}
}

// -Internal-----------------------------------------------------------------

func SimplexNoise(opts *NoiseOptions) (*[]float64, error) {
	return &[]float64{0.0}, nil
}
