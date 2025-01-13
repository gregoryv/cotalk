package main

import (
	_ "embed"
	"fmt"
	"html"
	"os/exec"
	"regexp"
	"strings"
	"sync"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
)

type Deck struct {
	Title  string
	Author string
	Date   string

	AutoCover bool
	cover     *Element

	AutoTOC bool
	toc     *Element

	cards []*Element

	mkuser sync.Once
	user   *CSS

	lastH2 *Element
}

func (p *Deck) CSS() *CSS {
	// vh
	scale := 2

	var footerHeight int = scale
	var headerHeight int = scale * 7
	fontSize := scale

	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
		"background-color: #fff",
	)
	css.Style(".page",
		"height: 100vh",
		"position: relative",
		"margin-bottom: 1vh",
	)
	css.Style(".page .view",
		"position: absolute",
		"top: 0",
		"left: 0",
		"right: 0",
		"bottom: "+vh(footerHeight),
	)
	css.Style(".page .footer",
		"position: absolute",
		"bottom: 0",
		"left: 0",
		"right: 0",
		"text-align: center",
		"height: "+vh(footerHeight),
	)
	css.Style(".page .footer .next",
		"position: absolute",
		"right: 1em",
		"cursor: pointer",
	)
	css.Style(".page .footer .prev",
		"position: absolute",
		"left: 1em",
		"cursor: pointer",
	)
	css.Style(".page .view",
		"font-size: "+vw(fontSize),
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style(".page .view .header",
		"text-align: center",
		"height: "+vh(headerHeight),
		"overflow: hidden",
	)
	css.Style(".header",
		"background-image: url('img/blue_gopher.svg')",
		"background-repeat: no-repeat",
		"background-position: 80px 10px",
	)
	css.Style(".header h1, .header h2, .header .group",
		"font-weight: bold",
		"margin-top: 3vh",
		"font-size: 5vh",
		"margin-bottom: 0vh",
	)
	css.Style(".header h3",
		"margin-top: 0vh",
		"font-size: 3vh",
	)
	css.Style(".page .view .slide",
		"margin: auto",
		"padding: 0.8vw 1.6vw 0.8vw 1.6vw",
		"height: "+vh(100-2*footerHeight-headerHeight-3),
		"overflow: hidden",
	)
	css.Style(".page .view .cover",
		"display: flex",
		"justify-content: center",
		"align-items: center",
		"height: "+vh(100-2*footerHeight-headerHeight),
		"text-align: center",
	)
	// toc
	css.Style(".toc",
		"position: absolute",
		"left: 27vw",
		"width: "+vw(55),
		"padding-top: 1em",
	)

	css.Style(".toc ul",
		"margin-top: 0",
	)
	css.Style("a",
		"text-decoration: none",
		"color: #000",
	)
	css.Style("a:hover",
		"text-decoration: underline",
	)
	css.Style(".h3",
		"margin-left: 5vw",
		"list-style-type: circle",
	)

	css.Style(".trail",
		"position: absolute",
		"height: 6px",
		"width: 6px",
		"border-radius: 3px",
		"background: red",
	)

	css = css.With(layoutView())
	css = css.With(srcfile(fontSize))
	css = css.With(highlightColors())
	if p.user != nil {
		css = css.With(p.user)
	}
	return css
}

func (p *Deck) Style(x string, v ...string) {
	p.mkuser.Do(func() {
		p.user = NewCSS()
	})
	p.user.Style(x, v...)
}

func (p *Deck) NewCard(elements ...any) {
	c := p.newCard(elements...)
	p.cards = append(p.cards, c)
}

func (p *Deck) newCard(elements ...any) *Element {
	header := Div(Class("header"),
		p.headings(elements[0].(*Element)),
	)
	slide := Div(Class("slide"))
	slide.With(elements[1:]...)
	return Wrap(header, slide)
}

func (p *Deck) headings(e *Element) any {
	switch e.Name {
	case "h2":
		p.lastH2 = e
		return e

	case "h3":
		return Wrap(
			Div(Class("group"), p.lastH2.Children[0]),
			e,
		)
	default:
		return e
	}
}

func (p *Deck) Document() *Page {
	body := Body()
	var cards []*Element

	// create cover page if not set
	cover := Wrap(
		Div(Class("cover"),
			Table(
				Tr(
					Td(
						H1(p.Title),
					),
				),
				Tr(
					Td(
						func() *Element {
							if p.cover == nil {
								return Wrap()
							}
							return p.cover
						}(),
					),
				),
				Tr(
					Td(
						Br(), p.Author,
						Br(),
						Span(Attr("style", "font-size: 0.7em"),
							p.Date,
						),
					),
				),
			),
		),
	)

	cards = append(cards, cover)

	// table of deck
	toc := p.toc
	if toc == nil && p.AutoTOC {
		ul := Ul()
		nav := Nav(
			ul,
		)
		for i, root := range p.cards {
			WalkElements(root, func(e *Element) {
				if !(e.Name == "h2" || e.Name == "h3") {
					return
				}
				// +2 skip cover page and toc
				a := A(Href(fmt.Sprintf("#%v", i+3)), e.Text())
				ul.With(Li(Class(e.Name), a))
			})
		}

		toc = p.newCard(
			H2(p.Title),
			Div(Class("toc"), nav),
		)
	}
	if toc != nil {
		cards = append(cards, toc)
	}
	cards = append(cards, p.cards...)

	for i, page := range cards {
		pageIndex := i + 1
		deck := Div(Class("view"))
		deck.With(page.Children...)

		body.With(
			Div(Class("page"), Attr("id", pageIndex),
				deck,
				footer(pageIndex, cards),
			),
		)
	}

	return NewPage(
		Html(
			Head(
				Meta(Charset("utf-8")),
				Title(
					p.Title,
				),
				Style(p.CSS()),
			),
			body,
			Script(
				string(enhance),

				string(trail),
			),
		),
	)
}

func footer(pageIndex int, cards []*Element) *Element {
	return Div(Class("footer"),
		prev(pageIndex-2, cards),
		next(pageIndex, cards),
		pageIndex, "/", len(cards),
	)
}

func prev(i int, cards []*Element) *Element {
	if i >= len(cards) || i <= 0 {
		return Em(Class("prev"), "&nbsp;")
	}
	e := cards[i].Children[0].(*Element)
	return Em(Class("prev"), Attr("onclick", "previousPage()"), "&laquo; ", e.Text())
}

func next(i int, cards []*Element) *Element {
	if i >= len(cards) || i == 1 {
		return Em(Class("next"), "&nbsp;")
	}
	e := cards[i].Children[0].(*Element)
	return Em(Class("next"), Attr("onclick", "nextPage()"), e.Text(), " &raquo;")
}

//go:embed enhance.js
var enhance []byte

//go:embed trail.js
var trail []byte

func vh(i int) string {
	return fmt.Sprintf("%vvh", i)
}

func vw(i int) string {
	return fmt.Sprintf("%vvw", i)
}

func LoadEscaped(filename string) *Element {
	v := files.MustLoad(filename)
	v = strings.TrimSpace(v)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = html.EscapeString(v)
	v = highlight(v)
	return Div(
		Class("srcfile"),
		Code(numLines(v, 1)),
	)
}

var load = Load

func Load(filename string) *Element {
	v := files.MustLoad(filename)
	v = strings.TrimSpace(v)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
	return Div(
		Class("srcfile"),
		Code(numLines(v, 1)),
	)
}

func LoadLinesCustom(filename string, from, to int, fontSize string) *Element {
	return LoadLines(filename, from, to).With(
		Attr("style", "font-size:"+fontSize),
	)
}

func LoadLines(filename string, from, to int) *Element {
	v := files.MustLoadLines(filename, from, to)
	v = strings.TrimSpace(v)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
	return Div(
		Class("srcfile"),
		Code(numLines(v, from)),
	)
}

func numLines(v string, n int) *Element {
	lines := strings.Split(v, "\n")
	ol := Ol()
	if n != 1 {
		ol.With(Attr("start", n))
	}
	for _, l := range lines {
		ol.With(
			Li(l),
		)
	}
	return ol
}

func godoc(pkg, short string) *Element {
	var out []byte
	if short != "" {
		out, _ = exec.Command("go", "doc", short, pkg).Output()
	} else {
		out, _ = exec.Command("go", "doc", pkg).Output()
	}
	v := string(out)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlightGoDoc(v)
	return Wrap(
		Pre(v),
		A(Attr("target", "_blank"),
			Href("https://pkg.go.dev/"+pkg),
			"pkg.go.dev/", pkg,
		),
	)
}

func ShellEscaped(cmd, filename string) *Element {
	v := files.MustLoad(filename)
	v = html.EscapeString(v)
	return Pre(Class("shell dark"), cmd, Br(), v)
}

func Shell(cmd, filename string) *Element {
	v := files.MustLoad(filename)
	return Pre(Class("shell dark"), cmd, Br(), v)
}

// highlight go source code
func highlight(v string) string {
	v = keywords.ReplaceAllString(v, `$1<span class="keyword">$2</span>$3`)
	v = types.ReplaceAllString(v, `$1<span class="type">$2</span>$3`)
	v = comments.ReplaceAllString(v, `<span class="comment">$1</span>`)
	return v
}

// highlightGoDoc output
func highlightGoDoc(v string) string {
	v = docKeywords.ReplaceAllString(v, `$1<span class="keyword">$2</span>$3`)
	// v = types.ReplaceAllString(v, `$1<span class="type">$2</span>$3`)
	v = comments.ReplaceAllString(v, `<span class="comment">$1</span>`)
	return v
}

var (
	types       = regexp.MustCompile(`(\W|\s)(http\.ResponseWriter|http\.Request)(\)|\n|,| \{)`)
	keywords    = regexp.MustCompile(`(\W?)(^package|const|select|defer|import|for|func|range|return|go|var|switch|if|else|case|label|type|struct|interface|default|fallthrough)(\W)`)
	docKeywords = regexp.MustCompile(`(\W?)(^package|func|type|struct|interface)(\W)`)
	comments    = regexp.MustCompile(`(//[^\n]*)`)
)

func highlightColors() *CSS {
	css := NewCSS()
	css.Style(".keyword", "color: darkviolet")
	css.Style(".type", "color: dodgerblue")
	css.Style(".comment, .comment>span", "color: darkgreen")
	return css
}

func srcfile(fontSize int) *CSS {
	css := NewCSS()
	css.Style(".srcfile",
		"margin-top: 1.6em",
		"margin-bottom: 1.6em",
		"background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAlCAMAAAB1cTk3AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAjVBMVEX9/fz4+Pb19fP8/Pz39/Xp6ejr6+rs7Ov29vTt7evo6Oj6+vr7+/vv7+/u7u3v7+37+/ru7uzw8PD////29vbx8fD6+vjq6unz8/L5+ff39/f9/f3y8vHo6Ofx8fHw8O/z8/P6+vn4+Pj09PL+/v729vXp6efy8vDv7+709PTz8/H7+/n8/Pv5+fju7u5JBELWAAAAAWJLR0QB/wIt3gAAAAlwSFlzAAALEwAACxMBAJqcGAAAAAd0SU1FB+cFHw8tIWAQRK8AAAC5SURBVCjPpZDbEoIgEIaXQLPUNNFOWJrRyQ7v/3gJAs2EFzTuBex+8+/uDwDIiglgGyItJJ4/9YM+DxSczcMowvFCFgmk4lpmVEZeiGolYbCmKjZiQgp4i5C305CVPezOPTVxMLD6wrorj79K0U6kpaQxM0/GPIm5gmdilKise9ZI4UW96HoLGcN3ybpF+hPaomxV+oCn/Ukv4DZMgdqQOkMMb0clG4L5uO3/wAGf3Ll9UJk5Lxpr/gOpxRF8cA+lxgAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAyMy0wNS0zMVQxMzo0NTozMyswMjowMJXHbJYAAAAldEVYdGRhdGU6bW9kaWZ5ADIwMjMtMDUtMzFUMTM6NDU6MzMrMDI6MDDkmtQqAAAAAElFTkSuQmCC)",
		"background-repeat: repeat-y",
		"background-size: 20px",
		"padding-left: 16px",
		"background-color: #fafafa",
		"tab-size: 4",
		"-moz-tab-size: 4",
		"font-size: "+vh(fontSize),
	)

	css.Style(".srcfile code",
		"margin-left: 0px",
		"padding-right: 20px", // match the background size of holes
		"padding: 0 0 0 0",
		"background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABQAAAAlCAMAAAB1cTk3AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAjVBMVEX9/fz4+Pb19fP8/Pz39/Xp6ejr6+rs7Ov29vTt7evo6Oj6+vr7+/vv7+/u7u3v7+37+/ru7uzw8PD////29vbx8fD6+vjq6unz8/L5+ff39/f9/f3y8vHo6Ofx8fHw8O/z8/P6+vn4+Pj09PL+/v729vXp6efy8vDv7+709PTz8/H7+/n8/Pv5+fju7u5JBELWAAAAAWJLR0QB/wIt3gAAAAlwSFlzAAALEwAACxMBAJqcGAAAAAd0SU1FB+cFHw8tIWAQRK8AAAC5SURBVCjPpZDbEoIgEIaXQLPUNNFOWJrRyQ7v/3gJAs2EFzTuBex+8+/uDwDIiglgGyItJJ4/9YM+DxSczcMowvFCFgmk4lpmVEZeiGolYbCmKjZiQgp4i5C305CVPezOPTVxMLD6wrorj79K0U6kpaQxM0/GPIm5gmdilKise9ZI4UW96HoLGcN3ybpF+hPaomxV+oCn/Ukv4DZMgdqQOkMMb0clG4L5uO3/wAGf3Ll9UJk5Lxpr/gOpxRF8cA+lxgAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAyMy0wNS0zMVQxMzo0NTozMyswMjowMJXHbJYAAAAldEVYdGRhdGU6bW9kaWZ5ADIwMjMtMDUtMzFUMTM6NDU6MzMrMDI6MDDkmtQqAAAAAElFTkSuQmCC)",
		"background-repeat: repeat-y",
		"background-position: right top",
		"background-size: 20px",
		"display: block",
		"text-align: left",
		"font-family: Inconsolata, monospace",
		"overflow: hidden",
	)
	css.Style(".srcfile ol",
		"margin-left: 0vw",
		"padding-left: 3vw",
	)
	css.Style(".srcfile ol li",
		"white-space: pre",
	)
	css.Style(".srcfile ol li::marker",
		"color: RGB(0,0,0,0.2)",
	)
	css.Style(".srcfile ol li:hover",
		//		"background-color: #b4eeb4",
		"background-color: RGB(180,238,180,0.3)",
	)

	css.Style(".srcfile code i",
		"font-style: normal",
		"color: #a2a2a2",
	)
	return css
}

func Double(e1, e2 any) *Element {
	div := Div(Class("double"))
	div.With(
		Div(
			Class("column left"),
			e1,
		),
	)
	div.With(
		Div(
			Class("column right"),
			e2,
		),
	)
	return div
}

func Middle(width int, element ...any) *Element {
	left := (100 - width) / 2
	div := Div(
		Attr(
			"style",
			fmt.Sprintf(
				"position: absolute; left: %s; width: %s", vw(left), vw(width),
			),
		),
	)
	div.With(element...)
	return div
}

func Center(element ...any) *Element {
	div := Div(Class("center"))
	div.With(element...)
	return div
}

func layoutView() *CSS {
	css := NewCSS()
	css.Style(".outline",
		"border: 1px solid black",
	)

	// center
	css.Style(".center",
		"text-align: center",
	)
	// double
	css.Style(".double")
	css.Style(".column.left",
		"position: absolute",
		"left: 4vw",
		"width: 40vw",
	)
	css.Style(".column.right",
		"position: absolute",
		"left: 50vw",
		"width: 40vw",
	)
	return css
}

func TwoCol(e1, e2 any, width int) *Element {
	leftWidth := 100 - 8 - width

	div := Div(Class("double"))
	div.With(
		Div(
			Class("column left"),
			Attr("style", fmt.Sprintf("width: %vvw", leftWidth)),
			e1,
		),
	)
	div.With(
		Div(
			Class("column right"),
			Attr("style", fmt.Sprintf("left: %vvw; width: %vvw", leftWidth+8, width)),
			e2,
		),
	)
	return div
}
