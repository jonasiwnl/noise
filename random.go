package noise

import "math/rand"

// -1D------------------------------------------------------------------------

func RandomNoise1(opts *NoiseOptions, dim int) *[]float64 {
	noise := make([]float64, dim)

	for x := range noise {
		noise[x] = rand.Float64()*float64(opts.Amplitude) + (opts.Zero - 0.5)
	}

	return &noise
}

// -2D------------------------------------------------------------------------

func RandomNoise2(opts *NoiseOptions, dim ...int) *[][]float64 {
	if len(dim) != 2 {
		panic("RandomNoise2: invalid dimensions")
	}

	noise := make([][]float64, dim[0])

	for x := range noise {
		noise[x] = make([]float64, dim[1])
		for y := range noise[x] {
			noise[x][y] =
				rand.Float64()*float64(opts.Amplitude) + (opts.Zero - 0.5)
		}
	}

	return &noise
}

// -3D------------------------------------------------------------------------

func RandomNoise3(opts *NoiseOptions, dim ...int) *[][][]float64 {
	if len(dim) != 3 {
		panic("RandomNoise3: invalid dimensions")
	}

	noise := make([][][]float64, dim[0])

	for x := range noise {
		noise[x] = make([][]float64, dim[1])
		for y := range noise[x] {
			noise[x][y] = make([]float64, dim[2])
			for z := range noise[x][y] {
				noise[x][y][z] =
					rand.Float64()*float64(opts.Amplitude) + (opts.Zero - 0.5)
			}
		}
	}

	return &noise
}

// -4D------------------------------------------------------------------------

func RandomNoise4(opts *NoiseOptions, dim ...int) *[][][][]float64 {
	if len(dim) != 4 {
		panic("RandomNoise4: invalid dimensions")
	}

	noise := make([][][][]float64, dim[0])

	for x := range noise {
		noise[x] = make([][][]float64, dim[1])
		for y := range noise[x] {
			noise[x][y] = make([][]float64, dim[2])
			for z := range noise[x][y] {
				noise[x][y][z] = make([]float64, dim[3])
				for w := range noise[x][y][z] {
					noise[x][y][z][w] =
						rand.Float64()*
							float64(opts.Amplitude) +
							(opts.Zero - 0.5)
				}
			}
		}
	}

	return &noise
}
