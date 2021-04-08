// fontcompare: compare two fonts

package main

import (
	"fmt"
	"os"

	"github.com/ajstarks/gensvg"
)

var (
	canvas = gensvg.New(os.Stdout)
	width  = 1000.0
	height = 600.0
	chars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789(){}[],.:;-+*/\\&_^%$#@!~`'\"<>"
	gstyle = "font-family:%s;font-size:%dpt;text-anchor:middle;fill:%s;fill-opacity:%.2f"
)

func letters(top, left float64, font, color string, opacity float32) {
	rows := 7
	cols := 13
	glyph := 0
	fontsize := 32.0
	spacing := fontsize * 2
	x := left
	y := top
	canvas.Gstyle(fmt.Sprintf(gstyle, font, fontsize, color, opacity))
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			canvas.Text(x, y, chars[glyph:glyph+1])
			glyph++
			x += spacing
		}
		x = left
		y += spacing
	}
	canvas.Gend()
}

func main() {
	var font1, font2 string
	if len(os.Args) < 2 {
		font1 = "sans"
		font2 = "sans-serif"
	} else {
		font1 = os.Args[1]
		font2 = os.Args[2]
	}
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")
	canvas.Text(80, 540, font1, "font-size:14pt; fill:blue; font-family:"+font1)
	canvas.Text(80, 560, font2, "font-size:14pt; fill:red;  font-family:"+font2)
	letters(100, 100, font1, "blue", 0.5)
	letters(100, 100, font2, "red", 0.5)
	canvas.End()
}
