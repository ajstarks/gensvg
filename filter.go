package gensvg

// Filter Effects:
// Most functions have common attributes (in, in2, result) defined in type Filterspec
// used as a common first argument.

// Filter begins a filter set
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#FilterElement
func (svg *SVG) Filter(id string, s ...string) {
	svg.printf(`<filter id="%s" %s`, id, endstyle(s, ">\n"))
}

// Fend ends a filter set
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#FilterElement
func (svg *SVG) Fend() {
	svg.println(`</filter>`)
}

// FeBlend specifies a Blend filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feBlendElement
func (svg *SVG) FeBlend(fs Filterspec, mode string, s ...string) {
	switch mode {
	case "normal", "multiply", "screen", "darken", "lighten":
		break
	default:
		mode = "normal"
	}
	svg.printf(`<feBlend %s mode="%s" %s`,
		fsattr(fs), mode, endstyle(s, emptyclose))
}

// FeColorMatrix specifies a color matrix filter primitive, with matrix values
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feColorMatrixElement
func (svg *SVG) FeColorMatrix(fs Filterspec, values [20]float64, s ...string) {
	svg.printf(`<feColorMatrix %s type="matrix" values="`, fsattr(fs))
	for _, v := range values {
		svg.printf(`%g `, v)
	}
	svg.printf(`" %s`, endstyle(s, emptyclose))
}

// FeColorMatrixHue specifies a color matrix filter primitive, with hue rotation values
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feColorMatrixElement
func (svg *SVG) FeColorMatrixHue(fs Filterspec, value float64, s ...string) {
	if value < -360 || value > 360 {
		value = 0
	}
	svg.printf(`<feColorMatrix %s type="hueRotate" values="%g" %s`,
		fsattr(fs), value, endstyle(s, emptyclose))
}

// FeColorMatrixSaturate specifies a color matrix filter primitive, with saturation values
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feColorMatrixElement
func (svg *SVG) FeColorMatrixSaturate(fs Filterspec, value float64, s ...string) {
	if value < 0 || value > 1 {
		value = 1
	}
	svg.printf(`<feColorMatrix %s type="saturate" values="%g" %s`,
		fsattr(fs), value, endstyle(s, emptyclose))
}

// FeColorMatrixLuminence specifies a color matrix filter primitive, with luminence values
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feColorMatrixElement
func (svg *SVG) FeColorMatrixLuminence(fs Filterspec, s ...string) {
	svg.printf(`<feColorMatrix %s type="luminenceToAlpha" %s`,
		fsattr(fs), endstyle(s, emptyclose))
}

// FeComponentTransfer begins a feComponent filter element
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feComponentTransferElement
func (svg *SVG) FeComponentTransfer() {
	svg.println(`<feComponentTransfer>`)
}

// FeCompEnd ends a feComponent filter element
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feComponentTransferElement
func (svg *SVG) FeCompEnd() {
	svg.println(`</feComponentTransfer>`)
}

// FeComposite specifies a feComposite filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feCompositeElement
func (svg *SVG) FeComposite(fs Filterspec, operator string, k1, k2, k3, k4 int, s ...string) {
	switch operator {
	case "over", "in", "out", "atop", "xor", "arithmetic":
		break
	default:
		operator = "over"
	}
	svg.printf(`<feComposite %s operator="%s" k1="%d" k2="%d" k3="%d" k4="%d" %s`,
		fsattr(fs), operator, k1, k2, k3, k4, endstyle(s, emptyclose))
}

// FeConvolveMatrix specifies a feConvolveMatrix filter primitive
// Standard referencd: http://www.w3.org/TR/SVG11/filters.html#feConvolveMatrixElement
func (svg *SVG) FeConvolveMatrix(fs Filterspec, matrix [9]int, s ...string) {
	svg.printf(`<feConvolveMatrix %s kernelMatrix="%d %d %d %d %d %d %d %d %d" %s`,
		fsattr(fs),
		matrix[0], matrix[1], matrix[2],
		matrix[3], matrix[4], matrix[5],
		matrix[6], matrix[7], matrix[8], endstyle(s, emptyclose))
}

// FeDiffuseLighting specifies a diffuse lighting filter primitive,
// a container for light source elements, end with DiffuseEnd()
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feComponentTransferElement
func (svg *SVG) FeDiffuseLighting(fs Filterspec, scale, constant float64, s ...string) {
	svg.printf(`<feDiffuseLighting %s surfaceScale="%g" diffuseConstant="%g" %s`,
		fsattr(fs), scale, constant, endstyle(s, `>`))
}

// FeDiffEnd ends a diffuse lighting filter primitive container
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feDiffuseLightingElement
func (svg *SVG) FeDiffEnd() {
	svg.println(`</feDiffuseLighting>`)
}

// FeDisplacementMap specifies a feDisplacementMap filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feDisplacementMapElement
func (svg *SVG) FeDisplacementMap(fs Filterspec, scale float64, xchannel, ychannel string, s ...string) {
	svg.printf(`<feDisplacementMap %s scale="%g" xChannelSelector="%s" yChannelSelector="%s" %s`,
		fsattr(fs), scale, imgchannel(xchannel), ychannel, endstyle(s, emptyclose))
}

// FeDistantLight specifies a feDistantLight filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feDistantLightElement
func (svg *SVG) FeDistantLight(fs Filterspec, azimuth, elevation float64, s ...string) {
	svg.printf(`<feDistantLight %s azimuth="%g" elevation="%g" %s`,
		fsattr(fs), azimuth, elevation, endstyle(s, emptyclose))
}

// FeFlood specifies a flood filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feFloodElement
func (svg *SVG) FeFlood(fs Filterspec, color string, opacity float64, s ...string) {
	svg.printf(`<feFlood %s flood-color="%s" flood-opacity="%g" %s`,
		fsattr(fs), color, opacity, endstyle(s, emptyclose))
}

// FeFunc{linear|Gamma|Table|Discrete} specify various types of feFunc{R|G|B|A} filter primitives
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feComponentTransferElement

// FeFuncLinear specifies a linear style function for the feFunc{R|G|B|A} filter element
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feComponentTransferElement
func (svg *SVG) FeFuncLinear(channel string, slope, intercept float64) {
	svg.printf(`<feFunc%s type="linear" slope="%g" intercept="%g"%s`,
		imgchannel(channel), slope, intercept, emptyclose)
}

// FeFuncGamma specifies the curve values for gamma correction for the feFunc{R|G|B|A} filter element
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feComponentTransferElement
func (svg *SVG) FeFuncGamma(channel string, amplitude, exponent, offset float64) {
	svg.printf(`<feFunc%s type="gamma" amplitude="%g" exponent="%g" offset="%g"%s`,
		imgchannel(channel), amplitude, exponent, offset, emptyclose)
}

// FeFuncTable specifies the table of values for the feFunc{R|G|B|A} filter element
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feComponentTransferElement
func (svg *SVG) FeFuncTable(channel string, tv []float64) {
	svg.printf(`<feFunc%s type="table"`, imgchannel(channel))
	svg.tablevalues(`tableValues`, tv)
}

// FeFuncDiscrete specifies the discrete values for the feFunc{R|G|B|A} filter element
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feComponentTransferElement
func (svg *SVG) FeFuncDiscrete(channel string, tv []float64) {
	svg.printf(`<feFunc%s type="discrete"`, imgchannel(channel))
	svg.tablevalues(`tableValues`, tv)
}

// FeGaussianBlur specifies a Gaussian Blur filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feGaussianBlurElement
func (svg *SVG) FeGaussianBlur(fs Filterspec, stdx, stdy float64, s ...string) {
	if stdx < 0 {
		stdx = 0
	}
	if stdy < 0 {
		stdy = 0
	}
	svg.printf(`<feGaussianBlur %s stdDeviation="%g %g" %s`,
		fsattr(fs), stdx, stdy, endstyle(s, emptyclose))
}

// FeImage specifies a feImage filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feImageElement
func (svg *SVG) FeImage(href string, result string, s ...string) {
	svg.printf(`<feImage xlink:href="%s" result="%s" %s`,
		href, result, endstyle(s, emptyclose))
}

// FeMerge specifies a feMerge filter primitive, containing feMerge elements
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feMergeElement
func (svg *SVG) FeMerge(nodes []string, s ...string) {
	svg.println(`<feMerge>`)
	for _, n := range nodes {
		svg.printf("<feMergeNode in=\"%s\"/>\n", n)
	}
	svg.println(`</feMerge>`)
}

// FeMorphology specifies a feMorphologyLight filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feMorphologyElement
func (svg *SVG) FeMorphology(fs Filterspec, operator string, xradius, yradius float64, s ...string) {
	switch operator {
	case "erode", "dilate":
		break
	default:
		operator = "erode"
	}
	svg.printf(`<feMorphology %s operator="%s" radius="%g %g" %s`,
		fsattr(fs), operator, xradius, yradius, endstyle(s, emptyclose))
}

// FeOffset specifies the feOffset filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feOffsetElement
func (svg *SVG) FeOffset(fs Filterspec, dx, dy int, s ...string) {
	svg.printf(`<feOffset %s dx="%d" dy="%d" %s`,
		fsattr(fs), dx, dy, endstyle(s, emptyclose))
}

// FePointLight specifies a fePpointLight filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#fePointLightElement
func (svg *SVG) FePointLight(x, y, z float64, s ...string) {
	svg.printf(`<fePointLight x="%g" y="%g" z="%g" %s`,
		x, y, z, endstyle(s, emptyclose))
}

// FeSpecularLighting specifies a specular lighting filter primitive,
// a container for light source elements, end with SpecularEnd()
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feSpecularLightingElement
func (svg *SVG) FeSpecularLighting(fs Filterspec, scale, constant float64, exponent int, color string, s ...string) {
	svg.printf(`<feSpecularLighting %s surfaceScale="%g" specularConstant="%g" specularExponent="%d" lighting-color="%s" %s`,
		fsattr(fs), scale, constant, exponent, color, endstyle(s, ">\n"))
}

// FeSpecEnd ends a specular lighting filter primitive container
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feSpecularLightingElement
func (svg *SVG) FeSpecEnd() {
	svg.println(`</feSpecularLighting>`)
}

// FeSpotLight specifies a feSpotLight filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feSpotLightElement
func (svg *SVG) FeSpotLight(fs Filterspec, x, y, z, px, py, pz float64, s ...string) {
	svg.printf(`<feSpotLight %s x="%g" y="%g" z="%g" pointsAtX="%g" pointsAtY="%g" pointsAtZ="%g" %s`,
		fsattr(fs), x, y, z, px, py, pz, endstyle(s, emptyclose))
}

// FeTile specifies the tile utility filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feTileElement
func (svg *SVG) FeTile(fs Filterspec, in string, s ...string) {
	svg.printf(`<feTile %s %s`, fsattr(fs), endstyle(s, emptyclose))
}

// FeTurbulence specifies a turbulence filter primitive
// Standard reference: http://www.w3.org/TR/SVG11/filters.html#feTurbulenceElement
func (svg *SVG) FeTurbulence(fs Filterspec, ftype string, bfx, bfy float64, octaves int, seed int64, stitch bool, s ...string) {
	if bfx < 0 || bfx > 1 {
		bfx = 0
	}
	if bfy < 0 || bfy > 1 {
		bfy = 0
	}
	switch ftype[0:1] {
	case "f", "F":
		ftype = "fractalNoise"
	case "t", "T":
		ftype = "turbulence"
	default:
		ftype = "turbulence"
	}

	var ss string
	if stitch {
		ss = "stitch"
	} else {
		ss = "noStitch"
	}
	svg.printf(`<feTurbulence %s type="%s" baseFrequency="%.2f %.2f" numOctaves="%d" seed="%d" stitchTiles="%s" %s`,
		fsattr(fs), ftype, bfx, bfy, octaves, seed, ss, endstyle(s, emptyclose))
}

// Filter Effects convenience functions, modeled after CSS versions

// Blur emulates the CSS blur filter
func (svg *SVG) Blur(p float64) {
	svg.FeGaussianBlur(Filterspec{}, p, p)
}

// Brightness emulates the CSS brightness filter
func (svg *SVG) Brightness(p float64) {
	svg.FeComponentTransfer()
	svg.FeFuncLinear("R", p, 0)
	svg.FeFuncLinear("G", p, 0)
	svg.FeFuncLinear("B", p, 0)
	svg.FeCompEnd()
}

// Contrast emulates the CSS contrast filter
//func (svg *SVG) Contrast(p float64) {
//}

// Dropshadow emulates the CSS dropshadow filter
//func (svg *SVG) Dropshadow(p float64) {
//}

// Grayscale eumulates the CSS grayscale filter
func (svg *SVG) Grayscale() {
	svg.FeColorMatrixSaturate(Filterspec{}, 0)
}

// HueRotate eumulates the CSS huerotate filter
func (svg *SVG) HueRotate(a float64) {
	svg.FeColorMatrixHue(Filterspec{}, a)
}

// Invert eumulates the CSS invert filter
func (svg *SVG) Invert() {
	svg.FeComponentTransfer()
	svg.FeFuncTable("R", []float64{1, 0})
	svg.FeFuncTable("G", []float64{1, 0})
	svg.FeFuncTable("B", []float64{1, 0})
	svg.FeCompEnd()
}

// Saturate eumulates the CSS saturate filter
func (svg *SVG) Saturate(p float64) {
	svg.FeColorMatrixSaturate(Filterspec{}, p)
}

// Sepia applies a sepia tone, emulating the CSS sepia filter
func (svg *SVG) Sepia() {
	var sepiamatrix = [20]float64{
		0.280, 0.450, 0.05, 0, 0,
		0.140, 0.390, 0.04, 0, 0,
		0.080, 0.280, 0.03, 0, 0,
		0, 0, 0, 1, 0,
	}
	svg.FeColorMatrix(Filterspec{}, sepiamatrix)
}
