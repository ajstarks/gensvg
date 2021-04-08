package gensvg

import "fmt"
// Animation

// Animate animates the specified link, using the specified attribute
// The animation starts at coordinate from, terminates at to, and repeats as specified
func (svg *SVG) Animate(link, attr string, from, to int, duration float64, repeat float64, s ...string) {
	svg.printf(`<animate %s attributeName="%s" from="%d" to="%d" dur="%gs" repeatCount="%s" %s`,
		href(link), attr, from, to, duration, repeatString(repeat), endstyle(s, emptyclose))
}

// AnimateMotion animates the referenced object along the specified path
func (svg *SVG) AnimateMotion(link, path string, duration float64, repeat float64, s ...string) {
	svg.printf(`<animateMotion %s dur="%gs" repeatCount="%s" %s<mpath %s/></animateMotion>
`, href(link), duration, repeatString(repeat), endstyle(s, ">"), href(path))
}

// AnimateTransform animates in the context of SVG transformations
func (svg *SVG) AnimateTransform(link, ttype, from, to string, duration float64, repeat float64, s ...string) {
	svg.printf(`<animateTransform %s attributeName="transform" type="%s" from="%s" to="%s" dur="%gs" repeatCount="%s" %s`,
		href(link), ttype, from, to, duration, repeatString(repeat), endstyle(s, emptyclose))
}

// AnimateTranslate animates the translation transformation
func (svg *SVG) AnimateTranslate(link string, fx, fy, tx, ty float64, duration float64, repeat float64, s ...string) {
	svg.AnimateTransform(link, "translate", coordpair(fx, fy), coordpair(tx, ty), duration, repeat, s...)
}

// AnimateRotate animates the rotation transformation
func (svg *SVG) AnimateRotate(link string, fs, fc, fe, ts, tc, te float64, duration float64, repeat float64, s ...string) {
	svg.AnimateTransform(link, "rotate", sce(fs, fc, fe), sce(ts, tc, te), duration, repeat, s...)
}

// AnimateScale animates the scale transformation
func (svg *SVG) AnimateScale(link string, from, to, duration float64, repeat float64, s ...string) {
	svg.AnimateTransform(link, "scale", fmt.Sprintf("%g", from), fmt.Sprintf("%g", to), duration, repeat, s...)
}

// AnimateSkewX animates the skewX transformation
func (svg *SVG) AnimateSkewX(link string, from, to, duration float64, repeat float64, s ...string) {
	svg.AnimateTransform(link, "skewX", fmt.Sprintf("%g", from), fmt.Sprintf("%g", to), duration, repeat, s...)
}

// AnimateSkewY animates the skewY transformation
func (svg *SVG) AnimateSkewY(link string, from, to, duration float64, repeat float64, s ...string) {
	svg.AnimateTransform(link, "skewY", fmt.Sprintf("%g", from), fmt.Sprintf("%g", to), duration, repeat, s...)
}

// Utility

// Grid draws a grid at the specified coordinate, dimensions, and spacing, with optional style.
func (svg *SVG) Grid(x float64, y float64, w float64, h float64, n float64, s ...string) {

	if len(s) > 0 {
		svg.Gstyle(s[0])
	}
	for ix := x; ix <= x+w; ix += n {
		svg.Line(ix, y, ix, y+h)
	}

	for iy := y; iy <= y+h; iy += n {
		svg.Line(x, iy, x+w, iy)
	}
	if len(s) > 0 {
		svg.Gend()
	}

}
