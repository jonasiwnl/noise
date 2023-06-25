package noise_test

import (
	"testing"

	n "github.com/jonasiwnl/noise"
)

func TestPerlin(t *testing.T) {
	opts := &n.NoiseOptions{
		Dimensions: n.NewDimensions(8, 8),
		Amplitude:  255,
		Zero:       0,
		Seed:       0,
	}

	data, err := n.PerlinNoise(opts)
	if err != nil {
		t.Error(err)
	}

	if len(*data) != 8 {
		t.Errorf("Expected %d, got %d", 8, len(*data))
	}
}
