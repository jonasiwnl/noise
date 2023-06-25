package noise

type NoiseOptions struct {
	Amplitude float64
	Zero      float64
	Seed      int
}

type Noise interface {
	Generate(opts *NoiseOptions) (*[]float64, error)
}

type NoiseReturn interface {
	*[][][][]float64 | *[][][]float64 | *[][]float64 | *[]float64
}
