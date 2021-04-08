package main

import (
	"os"

	"github.com/ajstarks/gensvg"
)

func main() {
	width, height := 500, 500
	cw, ch := float64(width), float64(height)
	rsize := 100.0
	csize := rsize / 2
	duration := 5.0
	repeat := 5.0
	imw, imh := 100, 144
	canvas := gensvg.New(os.Stdout)
	canvas.Start(cw, ch)
	canvas.Circle(csize, csize, csize, `fill="red"`, `id="circle"`)
	canvas.Image((cw/2)-float64((imw/2)), 0, imw, imh, "gopher.jpg", `id="gopher"`)
	canvas.Square(cw-rsize, 0, rsize, `fill="blue"`, `id="square"`)
	canvas.Animate("#circle", "cx", 0, width, duration, repeat)
	canvas.Animate("#circle", "cy", 0, height, duration, repeat)
	canvas.Animate("#square", "x", width, 0, duration, repeat)
	canvas.Animate("#square", "y", height, 0, duration, repeat)
	canvas.Animate("#gopher", "y", 0, height, duration, repeat)
	canvas.End()
}
