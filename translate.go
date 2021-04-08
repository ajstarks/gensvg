package gensvg

// Gtransform begins a group, with the specified transform
// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (svg *SVG) Gtransform(s string) { svg.println(group("transform", s)) }

// Translate begins coordinate translation, end with Gend()
// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (svg *SVG) Translate(x, y float64) { svg.Gtransform(translate(x, y, svg.Decimals)) }

// Scale scales the coordinate system by n, end with Gend()
// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (svg *SVG) Scale(n float64) { svg.Gtransform(scale(n)) }

// ScaleXY scales the coordinate system by dx and dy, end with Gend()
// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (svg *SVG) ScaleXY(dx, dy float64) { svg.Gtransform(scaleXY(dx, dy)) }

// SkewX skews the x coordinate system by angle a, end with Gend()
// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (svg *SVG) SkewX(a float64) { svg.Gtransform(skewX(a)) }

// SkewY skews the y coordinate system by angle a, end with Gend()
// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (svg *SVG) SkewY(a float64) { svg.Gtransform(skewY(a)) }

// SkewXY skews x and y coordinates by ax, ay respectively, end with Gend()
// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (svg *SVG) SkewXY(ax, ay float64) {
	svg.Gtransform(skewX(ax) + " " + skewY(ay))
}

// Rotate rotates the coordinate system by r degrees, end with Gend()
// Standard Reference: http://www.w3.org/TR/SVG11/coords.html#TransformAttribute
func (svg *SVG) Rotate(r float64) { svg.Gtransform(rotate(r)) }

// TranslateRotate translates the coordinate system to (x,y), then rotates to r degrees, end with Gend()
func (svg *SVG) TranslateRotate(x, y float64, r float64) {
	d := svg.Decimals
	svg.Gtransform(translate(x, y, d) + " " + rotate(r))
}

// RotateTranslate rotates the coordinate system r degrees, then translates to (x,y), end with Gend()
func (svg *SVG) RotateTranslate(x, y float64, r float64) {
	d := svg.Decimals
	svg.Gtransform(rotate(r) + " " + translate(x, y, d))
}
