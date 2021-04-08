package gensvg

// Paths

// Path draws an arbitrary path, the caller is responsible for structuring the path data
func (svg *SVG) Path(d string, s ...string) {
	svg.printf(`<path d="%s" %s`, d, endstyle(s, emptyclose))
}

// Arc draws an elliptical arc, with optional style, beginning coordinate at sx,sy, ending coordinate at ex, ey
// width and height of the arc are specified by ax, ay, the x axis rotation is r
// if sweep is true, then the arc will be drawn in a "positive-angle" direction (clockwise), if false,
// the arc is drawn counterclockwise.
// if large is true, the arc sweep angle is greater than or equal to 180 degrees,
// otherwise the arc sweep is less than 180 degrees
// http://www.w3.org/TR/SVG11/paths.html#PathDataEllipticalArcCommands
func (svg *SVG) Arc(sx float64, sy float64, ax float64, ay float64, r float64, large bool, sweep bool, ex float64, ey float64, s ...string) {
	d := svg.Decimals
	svg.printf(`%s A%s %.*f %s %s %s" %s`,
		ptag(sx, sy, d), coord(ax, ay, d), d, r, onezero(large), onezero(sweep), coord(ex, ey, d), endstyle(s, emptyclose))
}

// Bezier draws a cubic bezier curve, with optional style, beginning at sx,sy, ending at ex,ey
// with control points at cx,cy and px,py.
// Standard Reference: http://www.w3.org/TR/SVG11/paths.html#PathDataCubicBezierCommands
func (svg *SVG) Bezier(sx float64, sy float64, cx float64, cy float64, px float64, py float64, ex float64, ey float64, s ...string) {
	d := svg.Decimals
	svg.printf(`%s C%s %s %s" %s`,
		ptag(sx, sy, d), coord(cx, cy, d), coord(px, py, d), coord(ex, ey, d), endstyle(s, emptyclose))
}

// Qbez draws a quadratic bezier curver, with optional style
// beginning at sx,sy, ending at ex, sy with control points at cx, cy
// Standard Reference: http://www.w3.org/TR/SVG11/paths.html#PathDataQuadraticBezierCommands
func (svg *SVG) Qbez(sx float64, sy float64, cx float64, cy float64, ex float64, ey float64, s ...string) {
	d := svg.Decimals
	svg.printf(`%s Q%s %s" %s`,
		ptag(sx, sy, d), coord(cx, cy, d), coord(ex, ey, d), endstyle(s, emptyclose))
}

// Qbezier draws a Quadratic Bezier curve, with optional style, beginning at sx, sy, ending at tx,ty
// with control points are at cx,cy, ex,ey.
// Standard Reference: http://www.w3.org/TR/SVG11/paths.html#PathDataQuadraticBezierCommands
func (svg *SVG) Qbezier(sx float64, sy float64, cx float64, cy float64, ex float64, ey float64, tx float64, ty float64, s ...string) {
	d := svg.Decimals
	svg.printf(`%s Q%s %s T%s" %s`,
		ptag(sx, sy, d), coord(cx, cy, d), coord(ex, ey, d), coord(tx, ty, d), endstyle(s, emptyclose))
}

// Lines

// Line draws a straight line between two points, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#LineElement
func (svg *SVG) Line(x1 float64, y1 float64, x2 float64, y2 float64, s ...string) {
	d := svg.Decimals
	svg.printf(`<line x1="%.*f" y1="%.*f" x2="%.*f" y2="%.*f" %s`, d, x1, d, y1, d, x2, d, y2, endstyle(s, emptyclose))
}

// Polyline draws connected lines between coordinates, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/shapes.html#PolylineElement
func (svg *SVG) Polyline(x []float64, y []float64, s ...string) {
	svg.poly(x, y, "polyline", s...)
}
