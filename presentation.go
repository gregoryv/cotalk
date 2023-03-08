package cotalk

import (
	"embed"
	"fmt"
	"log"
	"os/exec"
	"strings"

	. "github.com/gregoryv/web"
	"github.com/preferit/deck"
)

func Presentation() *deck.Deck {
	d := deck.New()
	d.Title = "Go concurrency"
	d.Styles = append(d.Styles,
		themeOldstyle(),
		deck.HighlightColors(),
	)

	d.Slide(
		H1("Concurrency design in Go"),
		Span("2023 by Gregory Vinčić"),
		P(`Discussion and examples of using concepts related to
		concurrent design.`),

		Ul(Class("group"),
			Li("go routines"),
			Li("channels"),
		),
		Ul(Class("group"),
			Li("package context"),
			Li("package sync "),
			Li("go test bench"),
		),
		Ul(Class("group"),
			Li("examples and training"),
		),

		Br(Attr("clear", "all")),
	)

	d.Slide(H2("Background and history"),
		Img(Class("left"), Src("th_small.jpg")), Br(),
		A(Href("http://www.usingcsp.com/cspbook.pdf"),
			"Communicating Sequential Processes",
		),
		" C. A. R. Hoare, 1985",
		Br(Attr("clear", "all")),
	)

	d.Slide(H2("Goroutines"),
		A(Attr("target", "_blank"),
			Href("https://go.dev/tour/concurrency/1"),
			Img(Class("center"), Src("gotour_concurrency_1.png")),
		),
	)
	d.Slide(H2("channels"),
		A(Attr("target", "_blank"),
			Href("https://go.dev/tour/concurrency/2"),
			Img(Class("center"), Src("gotour_concurrency_2.png")),
		),
	)

	d.Slide(H2("package sync"),
		godoc("sync"),
	)

	d.Slide(H2("package context"),
		godoc("context"),

		Pre(deck.HighlightGoDoc(`
Programs that use Contexts should follow these rules

Do not store Contexts inside a struct type; instead, pass a Context
explicitly to each function that needs it. The Context should be the
first parameter, typically named ctx:

func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}

Do not pass a nil Context, even if a function permits it. Pass
context.TODO if you are unsure about which Context to use.

Use context Values only for request-scoped data that transits
processes and APIs, not for passing optional parameters to functions.
		`)),
	)

	// todo slide 7 test -bench
	d.Slide(H2("go test -bench"),
		flow("ex07/run.go", "go test -benchmark -bench .", "ex07/bench_result.txt"),
	)

	d.Slide(H2("Sequential"),
		P("Simple implementation though very low performance"),
		srcTest(10),
	)

	d.Slide(H2("Concurrent"),
		P("Depending on your concurrent requirements this implementation has couple of bugs"),
		srcTest(20),
	)

	d.Slide(H3("Concurrent - fix bug 1"),
		P("You cannot assume go routines start immediately."),
		srcTest(30),
		P("You might get a different result; but why does it fail?"),
	)

	d.Slide(H3("Concurrent - fix bug 2"),
		P("Unprotected write to slice"),
		srcTest(40),
	)

	d.Slide(H3("Concurrent - orderless"),
		P("For instance counting number of words"),
	)

	d.Slide(H3("Concurrent - ordered"),
		P("Reverse each word and put sequence back together"),
		Pre(`
convert
Aa Bb Cc Dd Ed Ff Gg Hh Ii Jj Kk Ll Mm Nn Oo Pp Qq Rr Ss Tt Uu Vv Xx Yy Zz

into
aA bB cC dD dE fF gG hH iI jJ kK lL mM nN oO pP qQ rR sS tT uU vV xX yY zZ
`),
	)

	d.Slide(H2("Concurrent"),
		P("Using channels"),
		load("ex50/run.go"),
	)

	return d
}

func themeOldstyle() *CSS {
	css := NewCSS()
	css.Style("body",
		"max-width: 40cm",
	)
	css.Style(".header",
		"width: 100%",
		"border-bottom: 1px solid #727272",
		"text-align: right",
		"margin-top: -2em",
		"margin-bottom: 1em",
	)
	css.Style("h1, h2, h3",
		"text-align: center",
	)
	css.Style(".srcfile",
		"margin-top: 1.6em",
		"margin-bottom: 1.6em",
		"background-image: url(\"printout-whole.png\")",
		"background-repeat: repeat-y",
		"padding-left: 36px",
		"background-color: #fafafa",
		"tab-size: 4",
		"-moz-tab-size: 4",
	)

	css.Style(".srcfile a",
		"float: right",
		"margin-right: 1.6em",
		"margin-top: -1.6em",
	)

	css.Style(".srcfile code",
		"padding: .6em 0 .6em 0",
		"background-image: url(\"printout-whole.png\")",
		"background-repeat: repeat-y",
		"background-position: right top",
		"display: block",
	)

	css.Style("code.srcfile",
		"padding-top: 0.6em",
		"padding-bottom: 0.6em",
	)
	css.Style("nav",
		"margin-top: 1em",
		"text-align: center",
	)
	css.Style("nav a:link, nav a:visited",
		"color: #727272",
		"padding: 0 5px",
		"margin: 0 2px",
		"text-decoration: none",
	)
	css.Style("nav a.current, nav a:hover",
		"color: black",
		"border-bottom: 1px solid black",
	)
	// toc
	css.Style("li.h3",
		"margin-left: 2em",
	)
	css.Style(".shell",
		"padding: 1em",
		"border-radius: 10px",
	)
	css.Style(".dark",
		"background-color: #2e2e34",
		"color: aliceblue",
	)
	css.Style(".light",
		"background-color: #ffffff",
		"color: #3b2616",
	)
	css.Style("img.center",
		"display: block",
		"margin-left: auto",
		"margin-right: auto",
	)
	css.Style("img.left",
		"float: left",
		"margin-right: 2em",
	)
	css.Style(".group",
		"float: left",
	)

	return css
}

func flow(src, cmd, result string) *Element {
	return Wrap(
		load(src),
		shell(cmd, result),
	)
}

func srcTest(ex int) *Element {
	return Wrap(
		load(fmt.Sprintf("ex%02v/run.go", ex)),
		shell(
			fmt.Sprintf("$ go test -count 1 -v ./ex%v/", ex),
			fmt.Sprintf("ex%v/test_result.html", ex),
		),
	)
}

func load(src string) *Element {

	v := mustLoad(src)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = deck.Highlight(v)
	return Div(
		Class("srcfile"),
		Pre(Code(v)),
	)
}

func godoc(pkg string) *Element {
	out, _ := exec.Command("go", "doc", "-short", pkg).Output()
	v := string(out)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = deck.HighlightGoDoc(v)
	return Wrap(
		Div(
			Class("srcfile"),
			Pre(Code(v)),
		),
		A(Attr("target", "_blank"),
			Href("https://pkg.go.dev/"+pkg),
			"pkg.go.dev/", pkg,
		),
	)
}

func shell(cmd, filename string) *Element {
	v := mustLoad(filename)
	return Pre(Class("shell dark"), cmd, Br(), v)
}

func mustLoad(src string) string {
	data, err := assets.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

//go:embed ex*
var assets embed.FS
