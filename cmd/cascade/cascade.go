// cascade: cascade animation
package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/gensvg"
)

func main() {
	width, height := 800.0, 600.0
	w2, w4, w34 := width/2, width/4, (width*3)/4
	n := 6.0
	dotsize := 10.0
	canvas := gensvg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	canvas.Gstyle("fill:none; stroke:gray")
	id := 0
	for y := 0.0; y < height; y += height / n {
		canvas.Qbez(w2, height, w4, y, 0, y, fmt.Sprintf("id=\"lc%02d\"", id))
		canvas.Qbez(width, y, w34, y, w2, height, fmt.Sprintf("id=\"rc%02d\"", id))
		canvas.Circle(0, 0, dotsize, "fill:red;stroke:none", fmt.Sprintf("id=\"ldot%02d\"", id))
		canvas.Circle(0, 0, dotsize, "fill:blue;stroke:none", fmt.Sprintf("id=\"rdot%02d\"", id))
		id++
	}
	dur := 5.0
	for i := 0; i < id; i++ {
		canvas.AnimateMotion(fmt.Sprintf("#ldot%02d", i), fmt.Sprintf("#lc%02d", i), dur, 0)
		canvas.AnimateMotion(fmt.Sprintf("#rdot%02d", i), fmt.Sprintf("#rc%02d", i), dur, 0)
		dur -= 0.2
	}
	canvas.Gend()
	canvas.End()
}
