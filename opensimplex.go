package noise

import (
	"math"
	"rand"
)

// -Internal-----------------------------------------------------------------

var gradients2 = []int8{
	5, 2, 2, 5,
	-5, 2, -2, 5,
	5, -2, 2, -5,
	-5, -2, -2, -5,
}

func extrapolate2(xsb float64, ysb float64, dx float64, dy float64, perm *[]int) float64 {
	index := (*perm)[int((*perm)[int(xsb)&0xFF]+int(ysb))&0xFF] & 0x0E
	return float64(gradients2[index])*dx + float64(gradients2[index+1])*dy
}

func openSimplexNoise2(x, y float64, perm *[]int) float64 {
	// I have no idea what I'm doing.
	stretch := float64(x+y) * F2
	xs := float64(x) + stretch
	ys := float64(y) + stretch

	xsb := math.Floor(xs)
	ysb := math.Floor(ys)

	squish := (xsb + ysb) * G2
	xb := xsb + squish
	yb := ysb + squish

	xins := xs - xsb
	yins := ys - ysb
	inSum := xins + yins

	dx0 := x - xb
	dy0 := y - yb

	value := 0.0

	dx1 := dx0 - 1 - G2
	dy1 := dy0 - G2
	attn1 := 2 - dx1*dx1 - dy1*dy1
	if attn1 > 0 {
		attn1 *= attn1
		value += attn1 * attn1 * extrapolate2(xsb+1, ysb+0, dx1, dy1, perm)
	}

	dx2 := dx0 - G2
	dy2 := dy0 - 1 - G2
	attn2 := 2 - dx2*dx2 - dy2*dy2
	if attn2 > 0 {
		attn2 *= attn2
		value += attn2 * attn2 * extrapolate2(xsb+0, ysb+1, dx2, dy2, perm)
	}

	if inSum <= 1 {
		zins := 1 - inSum
		if zins > xins || zins > yins {
			if xins > yins {
				xsvExt := xsb + 1
				ysvExt := ysb - 1
				dxExt := dx0 - 1
				dyExt := dy0 + 1
				value += zins *
					zins *
					extrapolate2(xsvExt, ysvExt, dxExt, dyExt, perm)
			} else {
				xsvExt := xsb - 1
				ysvExt := ysb + 1
				dxExt := dx0 + 1
				dyExt := dy0 - 1
				value += zins *
					zins *
					extrapolate2(xsvExt, ysvExt, dxExt, dyExt, perm)
			}
		}
	} else {
		zins := 2 - inSum
		if zins < xins || zins < yins {
			if xins > yins {
				xsvExt := xsb + 2
				ysvExt := ysb + 0
				dxExt := dx0 - 2
				dyExt := dy0 + 0
				value += zins *
					zins *
					extrapolate2(xsvExt, ysvExt, dxExt, dyExt, perm)
			} else {
				xsvExt := xsb + 0
				ysvExt := ysb + 2
				dxExt := dx0 + 0
				dyExt := dy0 - 2
				value += zins *
					zins *
					extrapolate2(xsvExt, ysvExt, dxExt, dyExt, perm)
			}
		}

		xsvExt := xsb + 1
		ysvExt := ysb + 1
		dxExt := dx0 - 1 - 2*G2
		dyExt := dy0 - 1 - 2*G2
		value += zins * zins * extrapolate2(xsvExt, ysvExt, dxExt, dyExt, perm)
	}
	return value
}

// -Internal-----------------------------------------------------------------

func OpenSimplexNoise2(opts *NoiseOptions, dim ...int) (*[][]float64, error) {
	if len(dim) != 2 {
		panic("RandomNoise2: invalid dimensions")
	}

	var seed int64
	if opts.Seed == 0 {
		seed = rand.Int()
	} else {
		seed = opts.Seed
	}

	perm := make([]int16, 256)
	source := make([]int16, 256)
	for i := range source {
		source[i] = int16(i)
	}

	// WTF?
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407

	noise := make([][]float64, dim[0])

	for x := range noise {
		noise[x] = make([]float64, dim[1])
		for y := range noise[x] {
			noise[x][y] = openSimplexNoise2(float64(x), float64(y))
		}
	}

	return &noise, nil
}
