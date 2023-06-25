package main

import n "github.com/jonasiwnl/noise"

func main() {
	opts := &n.NoiseOptions{
		Amplitude: 5,
		Zero:      0,
		Seed:      0,
	}

	x := 8
	y := 8
	z := 8

	noise := n.RandomNoise3(opts, x, y, z)

	for i := range *noise {
		for j := range (*noise)[i] {
			for p := range (*noise)[i][j] {
				print((*noise)[i][j][p], " ")
			}
			print("\n")
		}
		print("\n\n")
	}

	print("\n\n")
}
