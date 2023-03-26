package cotalk

import (
	"embed"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
)

func Presentation() *deck {
	d := newDeck()
	d.Title = "Go concurrency"
	d.Styles = append(d.Styles,
		themeOldstyle(),
		highlightColors(),
	)

	nav := &navbar{current: 1}
	alg := &algorithms{current: 1, atLine: 9} // 8 where first func starts

	d.Slide(
		Br(),
		A(Href("#2"), Img(Src("cotalk.png"))),
		Br(), Br(), Br(),
		Span("Gregory Vinčić, 2023"),
		Br(), Br(), Br(),
		nav,
	)

	d.Slide(
		H1("Content"),

		Ul(Class("group"),
			Li("background and history"), Li("goroutines"), Li("channels"),
			Li("package context"), Li("package sync "),
		),
		Ul(Class("group"),
			Li("go test bench"),
			Li("problem definition"),
			Li("sequential solution"),
		),
		Ul(Class("group"),
			Li("concurrent solution"),
			Li("fixing bugs"),
		),
		Ul(Class("group"),
			Li("using channels"),
			Li("sync and interrupt"),
			Li("compare performance"),
		),

		Br(Attr("clear", "all")),

		P(`Follow along by cloning the examples with `),

		Pre(Class("shell dark"),
			"$ git clone git@github.com:preferit/cotalk.git\n",
			"$ cd cotalk",
		),
		Br(Attr("clear", "all")),
		nav,
	)

	d.Slide(H2("Background and history"),
		Table(
			Tr(
				Td(Img(Src("th_small.jpg"))),
				Td(
					P("Sir Charles Antony Richard Hoare (",
						A(Href("https://en.wikipedia.org/wiki/Tony_Hoare"), "Tony Hoare"),

						`). Born 1934 in Sri Lanka, studied at Oxford and in
			Moscow. His research spanned program correctness, sorting
			and programming languages. His work is freely accessible
			online and the Go channel construct is his concept.
			`),

					Ul(
						Li(
							A(Href("https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf"), "Communicating Sequential Processes (CSP), paper 1978"),
						),
						Li(
							A(Href("http://www.usingcsp.com/cspbook.pdf"), "CSP, book 1985"),
						),
					),
				),
			),
		),
		nav,
	)

	d.Slide(H2("Goroutines"),
		A(Attr("target", "_blank"),
			Href("https://go.dev/tour/concurrency/1"),
			Img(Class("center"), Src("gotour_concurrency_1.png")),
		),
		nav,
	)

	d.Slide(H2("channels"),
		A(Attr("target", "_blank"),
			Href("https://go.dev/tour/concurrency/2"),
			Img(Class("center"), Src("gotour_concurrency_2.png")),
		),
		nav,
	)

	// packages
	// ----------------------------------------
	d.Slide(H2("package sync"),
		godoc("sync", ""),
		nav,
	)

	d.Slide(H2("package context"),
		Pre(highlightGoDoc(`package context

...

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

		godoc("context", "-short"),
		nav,
	)

	d.Slide(H2("go test -bench"),
		load("benchx_test.go"),
		shell(
			"$ go test -bench=BenchmarkX -v .",
			"testdata/benchx.html",
		),
		nav,
	)

	// Problem definition
	// ----------------------------------------
	d.Slide(H2("Problem"),
		P("A set of letters ", Code(Letters),

			` are available on a server; the problem is to GET them
			and assemble them in the correct order. Each letter is
			found on /L, ie. GET /3 would return '3'.`,
			//
		),
		P("To simplify the examples we define the Algorithm that does the work as",
			Pre(Code(highlight("type Algorithm func(work []*http.Request) (result []*http.Response)"))),
		),
		nav,
	)

	d.Slide(H2("Verification"),

		P(`Each algorithm in these examples is tested like this`),

		loadFunc("../alg_test.go", "BenchmarkAlg01"),

		P(`The problem and it's verification method Solve is found
in `, A(Href("https://github.com/preferit/cotalk/blob/main/problem.go"), "problem.go"),
		),
		nav,
	)

	d.Slide(H2("Sequential"),
		P("Simple implementation though very low performance"),
		alg,
		nav,
	)

	// Concurrent
	// ----------------------------------------
	d.Slide(H2("Concurrent using sync.WaitGroup"),
		alg,
		P("Why does it fail?"),
		nav,
	)

	d.Slide(H2("Arguments are evaluated at calltime"),
		alg,
		P(
			"You might get a different result; why does it still fail? and can the tooling help identify the problem, try ",
			Pre(Class("shell dark"),
				"$ go test -benchmem -bench=BenchmarkAlg3 -race -count 1",
			),
		),
		nav,
	)

	d.Slide(H2("Protect concurrent writes with sync.Mutex"),
		alg,
		P("Why does it fail?"),
		nav,
	)

	d.Slide(H2("Sort results"),
		alg,
		nav,
	)

	d.Slide(H2("Improved performance paid with complexity"),

		P(`Comparing the sequential working algorithm to the working
		concurrent one, tests reveal a substantial improvement.`),

		shell(
			`$ go test -benchmem -bench="(Alg1|Alg5)$"`,
			"testdata/compare_bench.html",
		),
		nav,
	)

	d.Slide(H2("Using channel"),
		alg,
		nav,
	)

	d.Slide(H2("Correct order using channel"),
		alg,
		P("There is still a bug in this code, do you see it?"),
		nav,
	)
	d.Slide(H2("Clean up resources"),
		alg,
		nav,
	)
	d.Slide(H2("Interrupt"),
		alg,
		nav,
	)
	d.Slide(H2("Respect context cancellation"),
		alg,
		nav,
	)
	d.Slide(H2("Compare all"),

		P(`In this example using channels and sync package primitives
		seem to yield more or less the same result. There performance
		boost would be to try and minimize number of allocations. But
		that is out of scope of this talk.`),

		shell(
			`$ go test -benchmem -bench="(Alg1|Alg5|Alg8)$"`,
			"testdata/compare_all.html",
		),
		nav,
	)

	d.Slide(H2("Go concurrency design summary"),

		Ul(
			Li("concurrency is difficult to get right; even in Go"),
			Li("tests are an invaluable tool when debugging concurrency issues"),
			Li("never assume performance optimizations, always measure"),
			Li("if performance is good enough with a sequential algorithm, skip the complexity of concurrency"),
		),
		//
		nav,
	)

	nav.max = len(d.Slides)
	return d
}

func load(src string) *Element {
	v := mustLoad(src)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
	v = numLines(v, 1)
	return Div(
		Class("srcfile"),
		Pre(Code(v)),
	)
}

func loadFunc(file, src string) *Element {
	v := files.MustLoadFunc(file, src)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
	return Div(
		Class("srcfile"),
		Pre(Code(v)),
	)
}

func numLines(v string, n int) string {
	lines := strings.Split(v, "\n")
	for i, l := range lines {
		lines[i] = fmt.Sprintf("<span class=line><i>%3v</i> %s</span>", n, l)
		n++
	}
	return strings.Join(lines, "\n")
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

//go:embed testdata *.go
var assets embed.FS

//go:embed docs/enhance.js
var enhancejs string

type navbar struct {
	max     int // number of slides
	current int // current slide
}

// BuildElement is called at time of rendering
func (b *navbar) BuildElement() *Element {
	ul := Ul()
	groupDivider := map[int]bool{
		8:  true,
		12: true, // concurrent
		17: true, // channels
	}

	for i := 0; i < b.max; i++ {
		j := i + 1
		hash := fmt.Sprintf("#%v", j)
		li := Li(A(Href(hash), j))
		if j == b.current {
			li.With(Class("current"))
		}
		if groupDivider[j] {
			ul.With(Li(" | "))
		}
		ul.With(li)
	}
	b.current++
	return Nav(ul)
}

type algorithms struct {
	current int // current algorithm
	atLine  int
}

func (a *algorithms) BuildElement() *Element {
	name := fmt.Sprintf("Alg%02v", a.current)
	a.current++
	fn := files.MustLoadFunc("../alg.go", name)
	lines := strings.Count(fn, "\n")
	from := a.atLine
	v := files.MustLoadLines("../alg.go", a.atLine, a.atLine+lines)
	a.atLine = a.atLine + lines + 2

	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
	v = numLines(v, from)
	return Wrap(
		Div(
			Class("srcfile"),
			Pre(Code(v)),
		),
		shell(
			fmt.Sprintf("$ go test -benchmem -bench=Benchmark%s .", name),
			fmt.Sprintf("testdata/%s_bench.html", strings.ToLower(name)),
		),
	)
}

// ----------------------------------------

// newDeck returns a deck with default styling and navigation on bottom
func newDeck() *deck {
	return &deck{
		Styles: []*CSS{},
	}
}

// Had this idea of a deck of slides; turned out to be less
// useful. Leaving it here for now.
type deck struct {
	Title  string // header title
	Slides []*Element
	Styles []*CSS // first one is the deck default styling
}

// Slide appends a new slide to the deck. elements can be anything
// supported by the web package.
func (d *deck) Slide(elements ...interface{}) {
	d.Slides = append(d.Slides, Wrap(elements...))
}

// Page returns a web page ready for use.
func (d *deck) Page() *Page {
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
	body.With(Script(enhancejs))
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

// ----------------------------------------

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
	v = types.ReplaceAllString(v, `$1<span class="type">$2</span>$3`)
	v = comments.ReplaceAllString(v, `<span class="comment">$1</span>`)
	return v
}

var types = regexp.MustCompile(`(\W)(\w+\.\w+)(\)|\n)`)
var keywords = regexp.MustCompile(`(\W?)(^package|select|defer|import|for|func|range|return|go|var|switch|if|case|label|type|struct|interface)(\W)`)
var docKeywords = regexp.MustCompile(`(\W?)(^package|func|type|struct|interface)(\W)`)
var comments = regexp.MustCompile(`(//[^\n]*)`)

func highlightColors() *CSS {
	css := NewCSS()
	css.Style(".keyword", "color: darkviolet")
	css.Style(".type", "color: dodgerblue")
	css.Style(".comment, .comment>span", "color: darkgreen")
	return css
}

// ----------------------------------------

func themeOldstyle() *CSS {
	css := NewCSS()
	css.Style("html, body",
		"margin: 0 0",
		"padding: 0 0",
	)
	css.Style(".slide",
		"padding: 10px 20%",
		"text-align: center",
		"height: calc( 100vh - 50px)",
	)
	css.Style(".slide ul, p, pre",
		"text-align: left",
		"font-size: 1.5vh",
	)

	// navbar
	css.Style(".slide nav",
		"text-align: center",
		"float: left",
		"clear: both",
		"width: 100%",
		"display: block",
		"margin-top: 2em",
	)
	css.Style(".slide nav ul",
		"list-style-type: none",
		"margin: 0 0",
		"padding: 0 0",
		"text-align: center",
	)
	css.Style(".slide nav ul li",
		"margin: 0 1px",
		"display: inline",
	)
	css.Style(".slide nav ul li.current a, .slide nav ul li:hover a",
		"background-color: #e2e2e2",
		"border-radius: 5px",
	)
	css.Style("nav a:link, nav a:visited",
		"color: #727272",
		"padding: 3px 5px",
		"text-decoration: none",
		"cursor: pointer",
	)

	// goish colors
	css.Style("a:link, a:visited",
		"color: #007d9c",
		"text-decoration: none",
	)
	css.Style("a:hover",
		"text-decoration: underline",
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

	css.Style(".srcfile code",
		"padding: .6em 0 2vh 0",
		"background-image: url(\"printout-whole.png\")",
		"background-repeat: repeat-y",
		"background-position: right top",
		"display: block",
		"text-align: left",
	)
	css.Style(".srcfile code .line", // each line
		"display: block",
		"width: 98%",
		"clear: both",
		"margin-bottom: -1.5vh",
	)
	css.Style(".srcfile code span:hover", // each line
		"background-color: wheat",
	)

	css.Style(".srcfile code i",
		"font-style: normal",
		"color: #a2a2a2",
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
		"text-align: left",
		"margin-right: 3em",
	)
	css.Style("td",
		"vertical-align: top",
	)
	css.Style("td:nth-child(2)",
		"padding-left: 2em",
	)

	return css
}
