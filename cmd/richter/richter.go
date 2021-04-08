// richter -- inspired by Gerhard Richter's 256 colors, 1974

package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/ajstarks/gensvg"
)

var canvas = gensvg.New(os.Stdout)

var width = 700.0
var height = 400.0

func main() {
	rand.Seed(int64(time.Now().Nanosecond()) % 1e9)
	canvas.Start(width, height)
	canvas.Title("Richter")
	canvas.Rect(0, 0, width, height, "fill:white")
	rw := 32.0
	rh := 18.0
	margin := 5.0
	for i, x := 0, 20.0; i < 16; i++ {
		x += (rw + margin)
		for j, y := 0, 20.0; j < 16; j++ {
			canvas.Rect(x, y, rw, rh, canvas.RGB(rand.Intn(255), rand.Intn(255), rand.Intn(255)))
			y += (rh + margin)
		}
	}
	canvas.End()
}
