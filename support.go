package gensvg

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Support functions

// coordpair returns a coordinate pair as a string
func coordpair(x, y float64) string {
	return fmt.Sprintf("%g %g", x, y)
}

// sce makes start, center, end coordinates string for animate transformations
func sce(start, center, end float64) string {
	return fmt.Sprintf("%g %g %g", start, center, end)
}

// repeatString computes the repeat string for animation methods
// repeat <= 0 --> "indefinite", otherwise the integer string
func repeatString(n float64) string {
	if n > 0 {
		return fmt.Sprintf("%g", n)
	}
	return "indefinite"
}

// style returns a style name,attribute string
func style(s string) string {
	if len(s) > 0 {
		return fmt.Sprintf(`style="%s"`, s)
	}
	return s
}

// pp returns a series of polygon points
func (svg *SVG) pp(x []float64, y []float64, tag string) {
	svg.print(tag)
	if len(x) != len(y) {
		svg.print(" ")
		return
	}
	lx := len(x) - 1
	d := svg.Decimals
	for i := 0; i < lx; i++ {
		svg.print(coord(x[i], y[i], d) + " ")
	}
	svg.print(coord(x[lx], y[lx], d))
}

// endstyle modifies an SVG object, with either a series of name="value" pairs,
// or a single string containing a style
func endstyle(s []string, endtag string) string {
	if len(s) > 0 {
		nv := ""
		for i := 0; i < len(s); i++ {
			if strings.Index(s[i], "=") > 0 {
				nv += (s[i]) + " "
			} else {
				nv += style(s[i]) + " "
			}
		}
		return nv + endtag
	}
	return endtag

}

// tt creates a xml element, tag containing s
func (svg *SVG) tt(tag string, s string) {
	svg.print("<" + tag + ">")
	xml.Escape(svg.Writer, []byte(s))
	svg.println("</" + tag + ">")
}

// poly compiles the polygon element
func (svg *SVG) poly(x []float64, y []float64, tag string, s ...string) {
	svg.pp(x, y, "<"+tag+" points=\"")
	svg.print(`" ` + endstyle(s, "/>\n"))
}

// onezero returns "0" or "1"
func onezero(flag bool) string {
	if flag {
		return "1"
	}
	return "0"
}

// pct returns a percetage, capped at 100
func pct(n uint8) uint8 {
	if n > 100 {
		return 100
	}
	return n
}

// islink determines if a string is a script reference
func islink(link string) bool {
	return strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "#") ||
		strings.HasPrefix(link, "../") || strings.HasPrefix(link, "./")
}

// group returns a group element
func group(tag string, value string) string { return fmt.Sprintf(`<g %s="%s">`, tag, value) }

// scale return the scale string for the transform
func scale(n float64) string { return fmt.Sprintf(`scale(%g)`, n) }

// scaleXY return the scale string for the transform
func scaleXY(dx, dy float64) string { return fmt.Sprintf(`scale(%g,%g)`, dx, dy) }

// skewx returns the skewX string for the transform
func skewX(angle float64) string { return fmt.Sprintf(`skewX(%g)`, angle) }

// skewx returns the skewX string for the transform
func skewY(angle float64) string { return fmt.Sprintf(`skewY(%g)`, angle) }

// rotate returns the rotate string for the transform
func rotate(r float64) string { return fmt.Sprintf(`rotate(%g)`, r) }

// translate returns the translate string for the transform
func translate(x, y float64, d int) string { return fmt.Sprintf(`translate(%.*f,%.*f)`, d, x, d, y) }

// coord returns a coordinate string
func coord(x float64, y float64, d int) string { return fmt.Sprintf(`%.*f,%.*f`, d, x, d, y) }

// ptag returns the beginning of the path element
func ptag(x float64, y float64, d int) string { return fmt.Sprintf(`<path d="M%s`, coord(x, y, d)) }

// loc returns the x and y coordinate attributes
func loc(x float64, y float64, d int) string { return fmt.Sprintf(`x="%.*f" y="%.*f"`, d, x, d, y) }

// href returns the href name and attribute
func href(s string) string { return fmt.Sprintf(`xlink:href="%s"`, s) }

// dim returns the dimension string (x, y coordinates and width, height)
func dim(x float64, y float64, w float64, h float64, d int) string {
	return fmt.Sprintf(`x="%.*f" y="%.*f" width="%.*f" height="%.*f"`, d, x, d, y, d, w, d, h)
}

// fsattr returns the XML attribute representation of a filterspec, ignoring empty attributes
func fsattr(s Filterspec) string {
	attrs := ""
	if len(s.In) > 0 {
		attrs += fmt.Sprintf(`in="%s" `, s.In)
	}
	if len(s.In2) > 0 {
		attrs += fmt.Sprintf(`in2="%s" `, s.In2)
	}
	if len(s.Result) > 0 {
		attrs += fmt.Sprintf(`result="%s" `, s.Result)
	}
	return attrs
}

// tablevalues outputs a series of values as a XML attribute
func (svg *SVG) tablevalues(s string, t []float64) {
	svg.printf(` %s="`, s)
	for i := 0; i < len(t)-1; i++ {
		svg.printf("%g ", t[i])
	}
	svg.printf(`%g"%s`, t[len(t)-1], emptyclose)
}

// imgchannel validates the image channel indicator
func imgchannel(c string) string {
	switch c {
	case "R", "G", "B", "A":
		return c
	case "r", "g", "b", "a":
		return strings.ToUpper(c)
	case "red", "green", "blue", "alpha":
		return strings.ToUpper(c[0:1])
	case "Red", "Green", "Blue", "Alpha":
		return c[0:1]
	}
	return "R"
}
