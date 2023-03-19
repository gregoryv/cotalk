package cotalk

import (
	"fmt"

	"github.com/gregoryv/web"
	. "github.com/gregoryv/web"
)

// New returns a deck with default styling and navigation on bottom
func New() *Deck {
	return &Deck{
		Styles: []*CSS{
			DefaultStyle(),
		},
		NavShow: Bottom,
	}
}

func DefaultStyle() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style("nav", "text-align: center")
	css.Style("nav ul",
		"list-style-type: none",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style("nav ul li",
		"margin: 0 4px",
		"cursor: pointer",
		"display: inline",
	)
	css.Style("nav ul li.current",
		"text-decoration: underline",
	)
	css.Style(".slide",
		"background-color: black",
		"padding: 10px",
		"text-align: center",
		"height: calc( 100vh - 50px)",
	)
	// goish colors
	css.Style("a:link, a:visited",
		"color: #007d9c",
		"text-decoration: none",
	)
	css.Style("a:hover",
		"text-decoration: underline",
	)
	return css
}

type Deck struct {
	Title   string // header title
	Slides  []*Element
	Styles  []*CSS // first one is the deck default styling
	NavShow Position
}

type Position int

const (
	Hidden Position = iota
	Top
	Bottom
)

// Slide appends a new slide to the deck. elements can be anything
// supported by the web package.
func (d *Deck) Slide(elements ...interface{}) {
	d.Slides = append(d.Slides, Wrap(elements...))
}

// Page returns a web page ready for use.
func (d *Deck) Page() *web.Page {
	styles := Style()
	for _, s := range d.Styles {
		styles.With(s)
	}
	body := Body()
	for i, content := range d.Slides {
		j := i + 1
		id := fmt.Sprintf("%v", j)
		slide := Div(Class("slide"), Id(id))

		slide.With(content)
		body.With(slide)
	}

	return NewFile("index.html",
		Html(
			Head(
				Title(d.Title),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				styles,
			),
			body,
		),
	)
}

type navbar struct {
	max     int // number of slides
	current int // current slide
}

// BuildElement is called at time of rendering
func (b *navbar) BuildElement() *Element {
	ul := Ul()
	for i := 0; i < b.max; i++ {
		j := i + 1
		hash := fmt.Sprintf("#%v", j)
		li := Li(A(Href(hash), j))
		if j == b.current {
			li.With(Class("current"))
		}
		ul.With(li)
	}
	b.current++
	return Nav(ul)
}
