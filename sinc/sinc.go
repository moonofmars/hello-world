// sinc, surface computes an SVG rendering of a 3-D surface function
package main

import (
	//	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	wid, het = 600, 320          //canvas size in pixels
	cels     = 100               //number of grid cells
	xyrange  = 30.0              //axis ranges(-xy,xy)
	xyscale  = wid / 2 / xyrange //pixels per x or y unit
	zscale   = het * 0.4         //pixels per z uint, z 缩放系数0.4
	angle    = math.Pi / 6       //angle of x,y axes =30°，垂直水平缩放系数
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml") //
	//fmt.Println("Hello World!")
	//io.WriteString(w, "Hello World!")
	svg := "<svg xmlns=\"http://www.w3.org/2000/svg\" style=\"fill:lime;stroke:purple;stroke-width:1\"" +
		" width=\"" + strconv.Itoa(wid) + "\" height=\"" + strconv.Itoa(het) + "\">\n"
	for i := 0; i < cels; i++ {
		for j := 0; j < cels; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			svg += "<polygon points=\"" + ax + "," + ay + " " + bx + "," + by + " " + cx + "," + cy + " " + dx + "," + dy + "\"/>\n"
			//ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svg += "</svg>"
	io.WriteString(w, svg)
}

func corner(i, j int) (string, string) {
	//find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cels - 0.5)
	y := xyrange * (float64(j)/cels - 0.5)
	//compute surface height z
	z := f(x, y)
	//project(x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := wid/2 + (x-y)*cos30*xyscale
	sy := het/2 + (x+y)*sin30*xyscale - z*zscale
	return strconv.FormatFloat(sx, 'g', 16, 64), strconv.FormatFloat(sy, 'g', 16, 64)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //distance from (0,0)
	return math.Sin(r) / r
}
