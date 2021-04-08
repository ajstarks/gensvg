package gensvg

// Shapes

// Circle centered at x,y, with radius r, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#CircleElement
func (svg *SVG) Circle(x float64, y float64, r float64, s ...string) {
	d := svg.Decimals
	svg.printf(`<circle cx="%.*f" cy="%.*f" r="%.*f" %s`, d, x, d, y, d, r, endstyle(s, emptyclose))
}

// Ellipse centered at x,y, centered at x,y with radii w, and h, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#EllipseElement
func (svg *SVG) Ellipse(x float64, y float64, w float64, h float64, s ...string) {
	d := svg.Decimals
	svg.printf(`<ellipse cx="%.*f" cy="%.*f" rx="%.*f" ry="%.*f" %s`,
		d, x, d, y, d, w, d, h, endstyle(s, emptyclose))
}

// Polygon draws a series of line segments using an array of x, y coordinates, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#PolygonElement
func (svg *SVG) Polygon(x []float64, y []float64, s ...string) {
	svg.poly(x, y, "polygon", s...)
}

// Rect draws a rectangle with upper left-hand corner at x,y, with width w, and height h, with optional style
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#RectElement
func (svg *SVG) Rect(x float64, y float64, w float64, h float64, s ...string) {
	svg.printf(`<rect %s %s`, dim(x, y, w, h, svg.Decimals), endstyle(s, emptyclose))
}

// CenterRect draws a rectangle with its center at x,y, with width w, and height h, with optional style
func (svg *SVG) CenterRect(x float64, y float64, w float64, h float64, s ...string) {
	svg.Rect(x-(w/2), y-(h/2), w, h, s...)
}

// Roundrect draws a rounded rectangle with upper the left-hand corner at x,y,
// with width w, and height h. The radii for the rounded portion
// are specified by rx (width), and ry (height).
// Style is optional.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#RectElement
func (svg *SVG) Roundrect(x float64, y float64, w float64, h float64, rx float64, ry float64, s ...string) {
	d := svg.Decimals
	svg.printf(`<rect %s rx="%.*f" ry="%.*f" %s`, dim(x, y, w, h, svg.Decimals), d, rx, d, ry, endstyle(s, emptyclose))
}

// Square draws a square with upper left corner at x,y with sides of length l, with optional style.
func (svg *SVG) Square(x float64, y float64, l float64, s ...string) {
	svg.Rect(x, y, l, l, s...)
}
