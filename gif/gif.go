// gif
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

//复合类型
var palette = []color.Color{color.White, color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}}

const (
	whiteIndex = 0 //1st color in palette
	blackIndex = 1 //next color in palette
)

func main() {
	//fmt.Println("Hello World!")
	//lissajous(os.Stdout)
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {
	cycles, _ := strconv.Atoi(r.FormValue("cycles"))
	//cycles := float64(cycle)
	fmt.Println(cycles)
	const (
		//cycles  = 5     //number of complete x oscillator resolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas covers[-size,size]
		nframes = 8     //number of animation frames
		delay   = 80    //delay between frams in 10ms units
	)

	freq := rand.Float64() * 3.0        //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} //struct
	phase := 0.0                        //phase difference
	for i := 0; i < nframes; i++ {      //每一次形成一个动画帧
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i%3+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim) //ignore encoding error
	//fmt.Println("err:", err)
	if err != nil {
		os.Stdout.WriteString(err.Error())
	}
}
