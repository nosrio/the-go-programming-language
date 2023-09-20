// The next program demon strates basic usage of Go’s stand ard image packages, which we’ll use
// to cre ate a sequence of bit-mapped images and then encode the sequence as a GIF animat ion.
// The images, cal le d Li ssajous figure s, were a staple visual effec t in sci-fi films of the 1960s.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
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
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles 	= 5
		res		= 0.001
		size	= 100
		nframes	= 64
		delay	= 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((i % (len(palette)-1))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}