// Mo dif y the Lissajous ser ver to read parameter values fro m the URL. For exam-
// ple, you mig ht arrange it so that a URL like http://localhost:8000/?cycles=20 sets the
// numb er of cyc les to 20 ins tead of the defau lt 5. Use the strconv.Atoi func tion to convert the
// st ring parameter into an int eger. You can see its document ation wit h go doc strconv.Atoi .

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(w http.ResponseWriter, r *http.Request) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	cycles := 5.0

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Failed to parse the form: %v\n", err)
		return 
	}
	if value, found := r.Form["cycles"]; found {
		token := value[0]
		cyclesInt, err := strconv.Atoi(token)
		if err != nil {
			fmt.Fprintf(w, "Failed to convert cycles: %v\n", err)
			return 
		}
		cycles = float64(cyclesInt)
	}
		
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size*0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}