package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height 	= 600, 320
	cells			= 100
	xyrange			= 30.0
	xyscale			= width / 2 / xyrange
	zscale			= height * 0.4
	angle			= math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type fz func(x, y float64) float64

func main() {
	usage  := "usage: ./main saddle|eggbox"
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	var f fz

	switch os.Args[1] {
	case "saddle":
		f = saddle
	case "eggbox":
		f = eggbox
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	svg(os.Stdout, f)
}

func svg(w io.Writer, f fz){
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
	"style='stroke: grey; fill: white; stroke-width: 0.7' "+
	"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			if math.IsInf(ax, 0) || math.IsInf(ay, 0) || math.IsInf(bx, 0) || math.IsInf(by, 0) || math.IsInf(cx, 0) || math.IsInf(cy, 0) || math.IsInf(dx, 0) || math.IsInf(dy, 0) {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy)
	}
	}
	fmt.Fprintln(w, "</svg>")

}
func corner(i, j int, f fz) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y) * cos30 * xyscale
	sy := height/2 + (x+y) * sin30 * xyscale - z*zscale

	return sx, sy
}

func eggbox(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12 
}

func saddle(x, y float64) float64 {
	return (math.Pow(x,2) - math.Pow(y, 2)) / 560
}
