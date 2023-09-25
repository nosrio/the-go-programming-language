package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x07, 0x1e, 0x22, 0xff},
	color.RGBA{0x1d, 0x78, 0x74, 0xff},
	color.RGBA{0x67, 0x92, 0x89, 0xff},
	color.RGBA{0xf4, 0xc0, 0x95, 0xff},
	color.RGBA{0xf5, 0xc6, 0x3c, 0xff},
	color.RGBA{0xf4, 0x7f, 0x5b, 0xff},
	color.RGBA{0xbb, 0x50, 0x98, 0xff},
	color.RGBA{0x7a, 0x51, 0x97, 0xff},
	color.RGBA{0x53, 0x44, 0xa9, 0xff},
	color.RGBA{0x02, 0x74, 0xbd, 0xff},
	color.RGBA{0xe9, 0xe6, 0xdd, 0xff},
	color.RGBA{0xc4, 0xad, 0x9d, 0xff},
	color.RGBA{0xf5, 0x72, 0x51, 0xff},
}

func main() {
	const (
		xmin, ymin, xmax, ymax 	= -2, -2, +2, +2
		width, height			= 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 6

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return palette[int(n) % len(palette)]
		}
	}
	return color.Black
}