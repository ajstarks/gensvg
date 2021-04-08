// funnel draws a funnel-like shape

package main

import (
	"os"

	"github.com/ajstarks/gensvg"
)

var canvas = gensvg.New(os.Stdout)
var width = 320.0
var height = 480.0

func funnel(bg int, fg int, grid float64, dim float64) {
	h := dim / 2
	canvas.Rect(0, 0, width, height, canvas.RGB(bg, bg, bg))
	for size := grid; size < width; size += grid {
		canvas.Ellipse(h, size, size/2, size/2, canvas.RGBA(fg, fg, fg, 0.2))
	}
}

func main() {
	canvas.Start(width, height)
	canvas.Title("Funnel")
	funnel(0, 255, 25, width)
	canvas.End()
}
