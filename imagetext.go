package gensvg

import (
	"encoding/xml"
	"fmt"
)
// Image places at x,y (upper left hand corner), the image with
// width w, and height h, referenced at link, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#ImageElement
func (svg *SVG) Image(x float64, y float64, w int, h int, link string, s ...string) {
	svg.printf(`<image %s %s %s`, dim(x, y, float64(w), float64(h), svg.Decimals), href(link), endstyle(s, emptyclose))
}

// Text places the specified text, t at x,y according to the style specified in s
// Standard Reference: http://www.w3.org/TR/SVG11/text.html#TextElement
func (svg *SVG) Text(x float64, y float64, t string, s ...string) {
	svg.printf(`<text %s %s`, loc(x, y, svg.Decimals), endstyle(s, ">"))
	xml.Escape(svg.Writer, []byte(t))
	svg.println(`</text>`)
}

// Textspan begins text, assuming a tspan will be included, end with TextEnd()
// Standard Reference: https://www.w3.org/TR/SVG11/text.html#TSpanElement
func (svg *SVG) Textspan(x float64, y float64, t string, s ...string) {
	svg.printf(`<text %s %s`, loc(x, y, svg.Decimals), endstyle(s, ">"))
	xml.Escape(svg.Writer, []byte(t))
}

// Span makes styled spanned text, should be proceeded by Textspan
// Standard Reference: https://www.w3.org/TR/SVG11/text.html#TSpanElement
func (svg *SVG) Span(t string, s ...string) {
	if len(s) == 0 {
		xml.Escape(svg.Writer, []byte(t))
		return
	}
	svg.printf(`<tspan %s`, endstyle(s, ">"))
	xml.Escape(svg.Writer, []byte(t))
	svg.printf(`</tspan>`)
}

// TextEnd ends spanned text
// Standard Reference: https://www.w3.org/TR/SVG11/text.html#TSpanElement
func (svg *SVG) TextEnd() {
	svg.println(`</text>`)
}

// Textpath places text optionally styled text along a previously defined path
// Standard Reference: http://www.w3.org/TR/SVG11/text.html#TextPathElement
func (svg *SVG) Textpath(t string, pathid string, s ...string) {
	svg.printf("<text %s<textPath xlink:href=\"%s\">", endstyle(s, ">"), pathid)
	xml.Escape(svg.Writer, []byte(t))
	svg.println(`</textPath></text>`)
}

// Textlines places a series of lines of text starting at x,y, at the specified size, fill, and alignment.
// Each line is spaced according to the spacing argument
func (svg *SVG) Textlines(x, y float64, s []string, size, spacing float64, fill, align string) {
	d := svg.Decimals
	svg.Gstyle(fmt.Sprintf("font-size:%.*fpx;fill:%s;text-anchor:%s", d, size, fill, align))
	for _, t := range s {
		svg.Text(x, y, t)
		y += spacing
	}
	svg.Gend()
}
