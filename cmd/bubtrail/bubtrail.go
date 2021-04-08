// bubtrail draws a randmonized trail of bubbles

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/gensvg"
)

var (
	width   = 1200.0
	height  = 600.0
	opacity = 0.5
	size    = 40.0
	niter   = 200
	canvas  = gensvg.New(os.Stdout)
)

func init() {
	flag.Float64Var(&size, "s", 40, "bubble size")
	flag.IntVar(&niter, "n", 200, "number of iterations")
	flag.Float64Var(&opacity, "o", 0.5, "opacity")
	flag.Parse()
	rand.Seed(int64(time.Now().Nanosecond()) % 1e9)
}

func background(v int) { canvas.Rect(0, 0, width, height, canvas.RGB(v, v, v)) }

func random(small, big float64) float64 {
	if small >= big {
		return small
	}
	r := (big - small)
	return float64(rand.Intn(int(r))) + small
}

func main() {
	var style string

	canvas.Start(width, height)
	canvas.Title("Bubble Trail")
	background(200)
	canvas.Gstyle(fmt.Sprintf("fill-opacity:%.2f;stroke:none", opacity))
	for i := 0; i < niter; i++ {
		x := float64(random(0, width))
		y := float64(random(height/3, (height*2)/3))
		r := float64(random(0, 10000))
		switch {
		case r >= 0 && r <= 2500:
			style = "fill:rgb(255,255,255)"
		case r > 2500 && r <= 5000:
			style = "fill:rgb(127,0,0)"
		case r > 5000 && r <= 7500:
			style = "fill:rgb(127,127,127)"
		case r > 7500 && r <= 10000:
			style = "fill:rgb(0,0,0)"
		}
		canvas.Circle(x, y, size, style)
	}
	canvas.Gend()
	canvas.End()
}
