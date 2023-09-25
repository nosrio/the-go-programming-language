package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
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


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "image/svg+xml")
		svg(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func svg(w io.Writer){
	var color string
	fmt.Fprintf(w,"<svg xmlns='http://www.w3.org/2000/svg' "+
	"style='stroke: grey; fill: white; stroke-width: 0.7' "+
	"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			if math.IsInf(ax, 0) || math.IsInf(ay, 0) || math.IsInf(bx, 0) || math.IsInf(by, 0) || math.IsInf(cx, 0) || math.IsInf(cy, 0) || math.IsInf(dx, 0) || math.IsInf(dy, 0) {
				continue
			}
			z := f(xyrange * (float64(i)/cells - 0.5), xyrange * (float64(j)/cells - 0.5))
			if z > 0 {
				color = "#ff0000"
			} else {
				color = "#0000ff"
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy, color)
	}
	}
	fmt.Fprintln(w, "</svg>")

}
func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y) * cos30 * xyscale
	sy := height/2 + (x+y) * sin30 * xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}