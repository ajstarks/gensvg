package gensvg

import (
	"encoding/xml"
	"fmt"
)

// Structure, Metadata, Scripting, Transformation, and Links

// Start begins the SVG document with the width w and height h.
// Other attributes may be optionally added, for example viewbox or additional namespaces
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#SVGElement
func (svg *SVG) Start(w float64, h float64, ns ...string) {
	d := svg.Decimals
	svg.printf(svginitfmt, svgtop, d, w, "", d, h, "")
	svg.genattr(ns)
}

// Startunit begins the SVG document, with width and height in the specified units
// Other attributes may be optionally added, for example viewbox or additional namespaces
func (svg *SVG) Startunit(w float64, h float64, unit string, ns ...string) {
	d := svg.Decimals
	svg.printf(svginitfmt, svgtop, d, w, unit, d, h, unit)
	svg.genattr(ns)
}

// Startpercent begins the SVG document, with width and height as percentages
// Other attributes may be optionally added, for example viewbox or additional namespaces
func (svg *SVG) Startpercent(w float64, h float64, ns ...string) {
	d := svg.Decimals
	svg.printf(svginitfmt, svgtop, d, w, "%", d, h, "%")
	svg.genattr(ns)
}

// Startview begins the SVG document, with the specified width, height, and viewbox
// Other attributes may be optionally added, for example viewbox or additional namespaces
func (svg *SVG) Startview(w, h, minx, miny, vw, vh float64) {
	d := svg.Decimals
	svg.Start(w, h, fmt.Sprintf(vbfmt, d, minx, d, miny, d, vw, d, vh))
}

// StartviewUnit begins the SVG document with the specified unit
func (svg *SVG) StartviewUnit(w, h float64, unit string, minx, miny, vw, vh float64) {
	d := svg.Decimals
	svg.Startunit(w, h, unit, fmt.Sprintf(vbfmt, d, minx, d, miny, d, vw, d, vh))
}

// Startraw begins the SVG document, passing arbitrary attributes
func (svg *SVG) Startraw(ns ...string) {
	svg.printf(svgtop)
	svg.genattr(ns)
}

// End the SVG document
func (svg *SVG) End() { svg.println("</svg>") }

// linkembed defines an element with a specified type,
// (for example "application/javascript", or "text/css").
// if the first variadic argument is a link, use only the link reference.
// Otherwise, treat those arguments as the text of the script (marked up as CDATA).
// if no data is specified, just close the element
func (svg *SVG) linkembed(tag string, scriptype string, data ...string) {
	svg.printf(`<%s type="%s"`, tag, scriptype)
	switch {
	case len(data) == 1 && islink(data[0]):
		svg.printf(" %s/>\n", href(data[0]))

	case len(data) > 0:
		svg.printf(">\n<![CDATA[\n")
		for _, v := range data {
			svg.println(v)
		}
		svg.printf("]]>\n</%s>\n", tag)

	default:
		svg.println(`/>`)
	}
}

// Script defines a script with a specified type, (for example "application/javascript").
func (svg *SVG) Script(scriptype string, data ...string) {
	svg.linkembed("script", scriptype, data...)
}

// Style defines the specified style (for example "text/css")
func (svg *SVG) Style(scriptype string, data ...string) {
	svg.linkembed("style", scriptype, data...)
}

// Gstyle begins a group, with the specified style.
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#GElement
func (svg *SVG) Gstyle(s string) { svg.println(group("style", s)) }

// Group begins a group with arbitrary attributes
func (svg *SVG) Group(s ...string) { svg.printf("<g %s\n", endstyle(s, `>`)) }

// Gid begins a group, with the specified id
func (svg *SVG) Gid(s string) {
	svg.print(`<g id="`)
	xml.Escape(svg.Writer, []byte(s))
	svg.println(`">`)
}

// Gend ends a group (must be paired with Gsttyle, Gtransform, Gid).
func (svg *SVG) Gend() { svg.println(`</g>`) }

// ClipPath defines a clip path
func (svg *SVG) ClipPath(s ...string) { svg.printf(`<clipPath %s`, endstyle(s, `>`)) }

// ClipEnd ends a ClipPath
func (svg *SVG) ClipEnd() {
	svg.println(`</clipPath>`)
}

// Def begins a defintion block.
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#DefsElement
func (svg *SVG) Def() { svg.println(`<defs>`) }

// DefEnd ends a defintion block.
func (svg *SVG) DefEnd() { svg.println(`</defs>`) }

// Marker defines a marker
// Standard reference: http://www.w3.org/TR/SVG11/painting.html#MarkerElement
func (svg *SVG) Marker(id string, x, y, width, height float64, s ...string) {
	d := svg.Decimals
	svg.printf(`<marker id="%s" refX="%.*f" refY="%.*f" markerWidth="%.*f" markerHeight="%.*f" %s`,
		id, d, x, d, y, d, width, d, height, endstyle(s, ">\n"))
}

// MarkerEnd ends a marker
func (svg *SVG) MarkerEnd() { svg.println(`</marker>`) }

// Pattern defines a pattern with the specified dimensions.
// The putype can be either "user" or "obj", which sets the patternUnits
// attribute to be either userSpaceOnUse or objectBoundingBox
// Standard reference: http://www.w3.org/TR/SVG11/pservers.html#Patterns
func (svg *SVG) Pattern(id string, x, y, width, height float64, putype string, s ...string) {
	puattr := "userSpaceOnUse"
	if putype != "user" {
		puattr = "objectBoundingBox"
	}
	d := svg.Decimals
	svg.printf(`<pattern id="%s" x="%.*f" y="%.*f" width="%.*f" height="%.*f" patternUnits="%s" %s`,
		id, d, x, d, y, d, width, d, height, puattr, endstyle(s, ">\n"))
}

// PatternEnd ends a marker
func (svg *SVG) PatternEnd() { svg.println(`</pattern>`) }

// Desc specified the text of the description tag.
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#DescElement
func (svg *SVG) Desc(s string) { svg.tt("desc", s) }

// Title specified the text of the title tag.
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#TitleElement
func (svg *SVG) Title(s string) { svg.tt("title", s) }

// Link begins a link named "name", with the specified title.
// Standard Reference: http://www.w3.org/TR/SVG11/linking.html#Links
func (svg *SVG) Link(href string, title string) {
	svg.printf("<a xlink:href=\"%s\" xlink:title=\"", href)
	xml.Escape(svg.Writer, []byte(title))
	svg.println("\">")
}

// LinkEnd ends a link.
func (svg *SVG) LinkEnd() { svg.println(`</a>`) }

// Use places the object referenced at link at the location x, y, with optional style.
// Standard Reference: http://www.w3.org/TR/SVG11/struct.html#UseElement
func (svg *SVG) Use(x float64, y float64, link string, s ...string) {
	svg.printf(`<use %s %s %s`, loc(x, y, svg.Decimals), href(link), endstyle(s, emptyclose))
}

// Mask creates a mask with a specified id, dimension, and optional style.
func (svg *SVG) Mask(id string, x float64, y float64, w float64, h float64, s ...string) {
	d := svg.Decimals
	svg.printf(`<mask id="%s" x="%.*f" y="%.*f" width="%.*f" height="%.*f" %s`, id, d, x, d, y, d, w, d, h, endstyle(s, `>`))
}

// MaskEnd ends a Mask.
func (svg *SVG) MaskEnd() { svg.println(`</mask>`) }
