// tumblrgrid: display a flexible grid of pictures from tumblr, filtered by tags

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ajstarks/gensvg"
)

var (
	localfile, postlink         bool
	ncols, thumbwidth, piclimit int
	gutter                      float64
	filtertag                   string
)

const (
	apiKey = "APIKEY"
	apifmt = "http://api.tumblr.com/v2/blog/%s/posts?api_key=%s&type=photo&limit=%d"
)

// Tumblr is the JSON data descriptions
type Tumblr struct {
	Meta     meta
	Response response
}

type meta struct {
	Msg    string
	Status int
}

type response struct {
	Blog       blog
	Posts      []posts
	TotalPosts int
}

type blog struct {
	Name        string
	Posts       int
	Title       string
	Updated     int
	Description string
	URL         string
}

type posts struct {
	Photos  []photos
	Tags    []string
	Type    string
	LinkURL string
	PostURL string
}

type photos struct {
	AltSizes []altsizes
}

type altsizes struct {
	Height int
	URL    string
	Width  int
}

// resource gives a ReadCloser given a local file or URL
func resource(name string) (io.ReadCloser, error) {
	if len(name) == 0 {
		return os.Stdin, nil
	}
	if localfile {
		return os.Open(name)
	}
	h, err := http.Get(fmt.Sprintf(apifmt, name, apiKey, piclimit))
	return h.Body, err
}

// grid displays tumblr photos in a flexible grid
func grid(canvas *gensvg.SVG, location string, x, y float64, nc int, gutter float64) {
	var (
		t   Tumblr
		r   io.ReadCloser
		err error
		b   []byte
	)

	//get data from the resource, put it into the data structure
	if r, err = resource(location); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer r.Close()
	if b, err = ioutil.ReadAll(r); err != nil {
		fmt.Fprintf(os.Stderr, "%s %v\n", location, err)
		return
	}

	if err = json.Unmarshal(b, &t); err != nil {
		fmt.Fprintf(os.Stderr, "%s, %v\n", location, err)
		return
	}

	// create the linked blog title
	fontsize := float64(thumbwidth / 3)
	title := t.Response.Blog.Title

	if len(title) == 0 {
		title = t.Response.Blog.Name
	}
	canvas.Link(t.Response.Blog.URL, t.Response.Blog.Name)
	if nc < 4 { // if the columns are too narrow, rotate the title text 90 degrees
		canvas.TranslateRotate(x, y/2, 90)
		canvas.Text(fontsize, fontsize, title, "fill:lightgray")
		canvas.Gend()
	} else {
		canvas.Text(x, y/2, title, "fill:lightgray")
	}
	canvas.LinkEnd()

	// walk through the posts, displaying thumbnails, filtered by tags
	np := 0
	xp := x

	for _, posts := range t.Response.Posts {
		if np >= piclimit {
			break
		}
		if !intag(filtertag, posts.Tags) {
			continue
		}
		for _, photos := range posts.Photos {
			for i, p := range photos.AltSizes {
				if i == 0 { // link to the first image in the list
					if postlink {
						canvas.Link(posts.PostURL, "Photo")
					} else {
						canvas.Link(p.URL, "Photo")
					}
				}
				if p.Width == thumbwidth {
					np++
					canvas.Image(xp, y, p.Width, p.Width, p.URL)
					xp += float64(p.Width) + gutter
					if np%nc == 0 {
						xp = x
						y += float64(p.Width) + gutter
					}
				}
			}
			canvas.LinkEnd()
		}
	}
}

// intag searches for tags in list
func intag(tag string, list []string) bool {
	if len(tag) == 0 {
		return true
	}
	for _, s := range list {
		if s == tag {
			return true
		}
	}
	return false
}

// init sets up command flags
func init() {
	flag.BoolVar(&localfile, "f", false, "read from local files")
	flag.BoolVar(&postlink, "p", false, "link to original post")
	flag.IntVar(&ncols, "nc", 5, "number of columns")
	flag.Float64Var(&gutter, "g", 5, "gutter (pixels)")
	flag.IntVar(&thumbwidth, "tw", 75, "thumbnail width")
	flag.IntVar(&piclimit, "n", 30, "picture limit")
	flag.StringVar(&filtertag, "tag", "", "filter tag")
	flag.Parse()
}

func main() {

	np := len(flag.Args())
	if np == 0 {
		np = 1
	}

	thalf := float64(thumbwidth / 2)
	x := thalf
	y := 50.0
	nrows := piclimit / ncols
	fnc := float64(ncols)
	fnr := float64(nrows)
	ftw := float64(thumbwidth)
	fnp := float64(np)
	colincr := (fnc * ftw) + (fnc * gutter) + thalf
	width := (colincr * fnp) + thalf
	height := (ftw * fnr) + (fnr * gutter) + y + thalf
	gstyle := "font-family:Calibri,sans-serif;font-size:18px"

	canvas := gensvg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:rgb(43,62,87)") // tumblr blue
	canvas.Gstyle(gstyle)
	if len(flag.Args()) == 0 {
		grid(canvas, "", x, y, ncols, gutter)
	} else {
		for _, f := range flag.Args() {
			grid(canvas, f, x, y, ncols, gutter)
			x += colincr
		}
	}
	canvas.Gend()
	canvas.End()
}
