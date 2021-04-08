package gensvg

import "fmt"
// Colors

// RGB specifies a fill color in terms of a (r)ed, (g)reen, (b)lue triple.
// Standard reference: http://www.w3.org/TR/css3-color/
func (svg *SVG) RGB(r int, g int, b int) string {
	return fmt.Sprintf(`fill:rgb(%d,%d,%d)`, r, g, b)
}

// RGBA specifies a fill color in terms of a (r)ed, (g)reen, (b)lue triple and opacity.
func (svg *SVG) RGBA(r int, g int, b int, a float64) string {
	return fmt.Sprintf(`fill-opacity:%.2f; %s`, a, svg.RGB(r, g, b))
}

// Gradients

// LinearGradient constructs a linear color gradient identified by id,
// along the vector defined by (x1,y1), and (x2,y2).
// The stop color sequence defined in sc. Coordinates are expressed as percentages.
func (svg *SVG) LinearGradient(id string, x1, y1, x2, y2 uint8, sc []Offcolor) {
	svg.printf("<linearGradient id=\"%s\" x1=\"%d%%\" y1=\"%d%%\" x2=\"%d%%\" y2=\"%d%%\">\n",
		id, pct(x1), pct(y1), pct(x2), pct(y2))
	svg.stopcolor(sc)
	svg.println("</linearGradient>")
}

// RadialGradient constructs a radial color gradient identified by id,
// centered at (cx,cy), with a radius of r.
// (fx, fy) define the location of the focal point of the light source.
// The stop color sequence defined in sc.
// Coordinates are expressed as percentages.
func (svg *SVG) RadialGradient(id string, cx, cy, r, fx, fy uint8, sc []Offcolor) {
	svg.printf("<radialGradient id=\"%s\" cx=\"%d%%\" cy=\"%d%%\" r=\"%d%%\" fx=\"%d%%\" fy=\"%d%%\">\n",
		id, pct(cx), pct(cy), pct(r), pct(fx), pct(fy))
	svg.stopcolor(sc)
	svg.println("</radialGradient>")
}

// stopcolor is a utility function used by the gradient functions
// to define a sequence of offsets (expressed as percentages) and colors
func (svg *SVG) stopcolor(oc []Offcolor) {
	for _, v := range oc {
		svg.printf("<stop offset=\"%d%%\" stop-color=\"%s\" stop-opacity=\"%.2f\"/>\n",
			pct(v.Offset), v.Color, v.Opacity)
	}
}
