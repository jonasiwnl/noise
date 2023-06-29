package main

import (
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	n "github.com/jonasiwnl/noise"
)

func handler(w http.ResponseWriter, r *http.Request) {
	graph := charts.NewGraph()

	graph.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "noise chart",
	}))

	graph.Render(w)

	opts := &n.NoiseOptions{
		Amplitude: 5,
		Zero:      0,
		Seed:      0,
	}

	x := 8
	y := 8

	noise, err := n.OpenSimplexNoise2(opts, x, y)

	if err != nil {
		panic(err)
	}

	for i := range *noise {
		for j := range (*noise)[i] {
			print((*noise)[i][j], " ")
		}
		print("\n")
	}

	print("\n\n")
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
