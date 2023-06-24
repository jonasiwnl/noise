package noise

type NoiseOptions struct {
	Dimensions *[]int
	Amplitude  float64
	Zero       float64
	Seed       int
}

type Noise interface {
	Generate(opts *NoiseOptions) (*[]float64, error)
}
