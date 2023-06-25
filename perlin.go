package noise

import "math"

// -Internal-----------------------------------------------------------------

func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

func lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}

func grad(hash int, coords ...int) float64 {
	h := hash & 15
	u := float64(coords[0])
	v := float64(coords[1])

	if h < 8 {
		u = -u
	}
	if h%4 == 0 || h%4 == 1 {
		v = -v
	} else {
		u, v = -v, u
	}

	return u + v
}

func hash(coords ...int) int {
	const prime = 1103515245
	const prime2 = 1979
	const prime3 = 269

	h := 0
	for _, coord := range coords {
		h = h*prime ^ coord*prime2
	}
	h ^= len(coords) * prime3
	return h & 0x7fffffff
}

func perlin(indices []int) float64 {
	dim := len(indices)
	res := 0.0
	freq := 1.0
	ampl := 1.0

	for i := 0; i < dim; i++ {
		x := float64(indices[i]) * freq
		intX := int(math.Floor(x))
		t := x - float64(intX)
		// u := fade(t)

		// Hash
		indices[i] = intX % 256
		nextIndices := make([]int, dim)
		copy(nextIndices, indices)
		nextIndices[i] = (intX + 1) % 256

		// Grad
		grad0 := grad(hash(indices...), indices...)
		grad1 := grad(hash(nextIndices...), nextIndices...)

		// Interpolate
		res += ampl * lerp(t, grad0, grad1)

		indices[i] = indices[i] % 255

		freq *= 2.0
		ampl *= 0.5
	}

	return res
}

// -Internal----------------------------------------------------------------

func PerlinNoise(opts *NoiseOptions, dim ...int) (*[]float64, error) {
	size := 1
	for i := range dim {
		size *= dim[i]
	}
	noise := make([]float64, size)

	var coords []int

	for i := range noise {
		i0 := i
		n := len(dim)
		coords = make([]int, n)

		for j := n - 1; j >= 0; j-- {
			coords[j] = i0 % dim[j]
			i0 /= dim[j]
		}

		noise[i] = perlin(coords)
	}

	return &noise, nil
}
