package main

import (
	"log"

	n "github.com/jonasiwnl/noise"
)

func main() {
	dims := []int{8, 8}

	opts := &n.NoiseOptions{
		Dimensions: &dims,
		Amplitude:  255,
		Zero:       0,
		Seed:       0,
	}

	noise, err := n.PerlinNoise(opts)

	if err != nil {
		log.Fatal(err)
	}

	for i := range *noise {
		if i%dims[0] == 0 {
			print("\n")
		}

		print(int((*noise)[i]), " ")
	}

	print("\n\n")
}
