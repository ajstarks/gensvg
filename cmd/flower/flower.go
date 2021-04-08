// flower - draw random flowers, inspired by Evelyn Eastmond's DesignBlocks gererated "grain2"

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/gensvg"
)

var (
	canvas    = gensvg.New(os.Stdout)
	niter     = flag.Int("n", 200, "number of iterations")
	width     = flag.Float64("w", 500, "width")
	height    = flag.Float64("h", 500, "height")
	thickness = flag.Float64("t", 10, "max petal thinkness")
	np        = flag.Float64("p", 15, "max number of petals")
	psize     = flag.Float64("s", 30, "max length of petals")
	opacity   = flag.Float64("o", 50, "max opacity (10-100)")
)

const flowerfmt = `stroke:rgb(%d,%d,%d); stroke-opacity:%.2f; stroke-width:%v`

func radial(xp, yp, n, r float64, style ...string) {
	var x, y, t, limit float64
	limit = 2.0 * math.Pi
	canvas.Gstyle(style[0])
	for t = 0.0; t < limit; t += limit / n {
		x = r * math.Cos(t)
		y = r * math.Sin(t)
		canvas.Line(xp, yp, xp+x, yp+y)
	}
	canvas.Gend()
}

func random(small, big float64) float64 {
	if small >= big {
		return small
	}
	r := (big - small)
	return float64(rand.Intn(int(r))) + small
}

func randrad(w, h float64, n int) {
	var x, y, o, s, t, p float64
	var r, g, b int
	for i := 0; i < n; i++ {
		x = random(0, w)
		y = random(0, h)
		r = rand.Intn(255)
		g = rand.Intn(255)
		b = rand.Intn(255)
		o = random(10, *opacity)
		s = random(10, *psize)
		t = random(2, *thickness)
		p = random(10, *np)
		radial(x, y, p, s, fmt.Sprintf(flowerfmt, r, g, b, o/100.0, t))
	}
}

func background(v int) { canvas.Rect(0, 0, *width, *height, canvas.RGB(v, v, v)) }

func init() {
	flag.Parse()
	rand.Seed(int64(time.Now().Nanosecond()) % 1e9)
}

func main() {
	canvas.Start(*width, *height)
	canvas.Title("Random Flowers")
	background(255)
	randrad(*width, *height, *niter)
	canvas.End()
}
