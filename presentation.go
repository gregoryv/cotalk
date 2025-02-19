package main

import (
	"embed"
	"fmt"
	"strings"

	. "github.com/gregoryv/web"
	"github.com/gregoryv/web/files"
)

func main() {
	d := Deck{
		Title:     "Go Concurrency",
		Author:    "Gregory Vinčić",
		AutoCover: true,
		cover:     Img(Src("cotalk.png")),
		AutoTOC:   true,
		Date:      "2025",
	}
	d.Style(".header", "background-color: #f6f5f4",
		"border-bottom: 1px inset #000",
	)
	d.Style(".page .footer", "bottom: 1vh") // slightly up
	d.Style(".right>*, .right .shell",
		"margin-right: 1em",
	)

	// center toc if short titles
	d.Style(".toc",
		"position: absolute",
		"font-size: 0.9em",
		"left: 13vw",
		"width: "+vw(73),
		"padding-top: 1em",
	)
	d.Style(".shell",
		"padding: 1em",
		"border-radius: 10px",
		"overflow: wrap",
	)
	d.Style(".dark",
		"background-color: #2e2e34",
		"color: aliceblue",
	)
	d.Style("nav",
		"column-count: 2", // Fix this at the end with a manual toc
		"font-size: 0.8em",
	)
	d.Style(".srcfile ol",
		"padding-right: 40px",
	)
	d.Style(".shell",
		"font-size: 1.5vw",
	)
	d.Style("p>a",
		"text-decoration: underline",
	)
	d.Style("quote",
		"display: inline-block",
		"font-style: italic",
		"padding-left: 2vw",
		"padding-right: 2vw",
		"font-size: 1.5vw",
	)
	d.Style("quote>a",
		"font-size: 1vw",
		"float: right",
	)
	d.Style("quote.small",
		"font-size: 1vw",
	)
	d.Style(".srcfile",
		"font-size: 1.3vw",
	)
	d.Style(".filename",
		"display: block",
		"font-size: 0.8vw",
		"text-align: right",
		"margin-top: 2em",
		"margin-bottom: -2em",
	)
	d.Style(".icons",
		"column-count: 3",
		"list-style-type: none",
	)

	d.Style(".small *",
		"font-size: 0.9vw",
	)
	d.Style(".small .filename",
		"margin-bottom: -1.5em",
	)
	d.Style(".stop>img",
		"text-align: center",
		"float:left",
		"margin-right: 20px",
	)
	d.Style(".stop>p",
		"padding-top: 50px",
		"color: red",
	)
	d.Style(".smallerFont srcfile",
		"font-size: 0.8em",
	)
	alg := &algorithms{
		current: 1,
		atLine:  9,
	} // 8 where first func starts

	d.NewCard(H2("Concurrency constructs"),
		Middle(80,
			Img(
				Src("th_small.jpg"),
				Attr("style", "float: left; margin: 1em 1em"),
			),

			P("Sir Charles Antony Richard Hoare (",
				A(
					Href("https://en.wikipedia.org/wiki/Tony_Hoare"),
					"Tony Hoare",
				),
				`). Born 1934 in Sri Lanka, studied at Oxford and in
			Moscow. His research spanned program correctness, sorting
			and programming languages. His work is freely accessible
			online and the channel concept is described in the book `,
				A(
					Href("http://www.7de.se/pub/cspbook.pdf"),
					"Communicating Sequential Processes (CSP) from 1985",
				),
			),
		),
	)

	d.NewCard(H3("Go routines"),
		Middle(50,
			A(Attr("target", "_blank"),
				Href("https://go.dev/tour/concurrency/1"),
				Img(Class("center"), Src("gotour_concurrency_1.png")),
			),
		),
	)

	d.NewCard(H3("Channels"),
		Middle(50,
			A(Attr("target", "_blank"),
				Href("https://go.dev/tour/concurrency/2"),
				Img(Class("center"), Src("gotour_concurrency_2.png")),
			),
		),
	)

	// packages
	// ----------------------------------------
	d.NewCard(H3("Package sync"),
		Middle(40,
			Div(fontSize("1.1vw"),
				godoc("sync", ""),
			),
		),
	)
	d.NewCard(H3("Package context"),
		TwoCol(
			Pre(fontSize("0.8vw"), highlightGoDoc(`package context

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

			Div(fontSize("0.8vw"),
				godoc("context", "-short"),
			),
			55,
		),
	)

	d.NewCard(H2("go test -bench"),
		TwoCol(
			load("benchx_test.go"),
			Shell(
				"$ go test -bench=BenchmarkX -v",
				"testdata/benchx.html",
			),
			55,
		),
	)

	// Problem definition
	// ----------------------------------------
	d.NewCard(H2("Problem"),
		Middle(55, load("problem.go")),
	)

	d.NewCard(H3("The letter challenge"),
		TwoCol(
			LoadLinesCustom("letters.go", 13, 38, "1.1vw"),
			LoadLinesCustom("letters.go", 40, 71, "0.95vw"),
			50,
		),
	)

	d.NewCard(H3("Verification"),
		P(`Each algorithm in these examples is tested like this`),
		LoadLines("alg_test.go", 10, 24),
	)

	d.NewCard(H3("Sequential"),
		P("Simple implementation though very low performance"),
		alg.next(),
	)
	smallFont := Attr("style", "font-size: 1vw")
	// Concurrent
	// ----------------------------------------
	d.NewCard(H2("Concurrent using sync.WaitGroup"),
		alg.next(
			P("Why does it fail?"),
		),
	)
	d.NewCard(H2("Concurrent using sync.WaitGroup Go 1.23"),
		TwoCol(
			alg.last,
			Wrap(
				Shell(
					"$ go1.23 test -benchmem -bench=BenchmarkAlg02",
					"testdata/alg02_go1_23_bench.html",
				).With(smallFont),
				P("Fixed!"),
			),
			45,
		),
	)

	d.NewCard(H2("Arguments are evaluated at calltime"),
		alg.next(
			P(
				"You might get a different result; why does it still fail? and can the tooling help identify the problem, try ",
				Pre(Class("shell dark"), smallFont,
					"$ go test -benchmem -bench=BenchmarkAlg03 -race -count 1",
				),
			),
		),
	)

	d.NewCard(H2("Protect concurrent writes with sync.Mutex"),
		alg.next(
			P("Why does it fail?"),
		),
	)

	d.NewCard(H2("Sort results"),
		alg.next(),
	)

	d.NewCard(H2("Improved performance paid with complexity"),

		P(`Comparing the sequential working algorithm to the working
		concurrent one, tests reveal a substantial improvement.`),

		Shell(
			`$ go test -benchmem -bench="(Alg01|Alg05)$"`,
			"testdata/compare_bench.html",
		),
	)

	d.NewCard(H2("Using channel"),
		alg.next(),
	)

	d.NewCard(H3("Correct order using channel"),
		alg.next(
			P("There is still a bug in this code, do you see it?"),
		),
		// resources are not cleaned up, ie. the channel is leaking
	)
	d.NewCard(H2("Clean up resources"),
		alg.next(),
	)
	d.NewCard(H2("Interrupt"),
		alg.nextCustom("1.0vw", "1.0vw", 45),
	)
	d.NewCard(H2("Respect context cancellation"),
		alg.nextCustom("1.0vw", "1.0vw", 45),
	)
	d.NewCard(H2("Compare all"),

		P(`In this example using channels and sync package primitives
		seem to yield more or less the same result. There performance
		boost would be to try and minimize number of allocations. But
		that is out of scope for this talk.`),

		Shell(
			`$ go test -benchmem -bench="(Alg01|Alg05|Alg08)$"`,
			"testdata/compare_all.html",
		),
	)

	d.NewCard(H2("Summary"),
		Middle(61,
			Ul(
				Li("Concurrency is difficult to get right; even in Go"),
				Li("Tests are invaluable when debugging concurrency issues"),
				Li("Measure, never assume performance optimizations"),
				Li("If performance is good enough with a sequential algorithm, skip the complexity of concurrency"),
			),
		),
		//
	)
	d.Document().SaveAs("docs/index.html")
}

func fontSize(v string) any {
	return Attr("style", "font-size:"+v)
}

//go:embed testdata *.go
var assets embed.FS

//go:embed docs/enhance.js
var enhancejs string

type algorithms struct {
	current int // current algorithm
	atLine  int

	last *Element
}

func (a *algorithms) next(extra ...interface{}) *Element {
	return a.nextCustom("1.1vw", "1.0vw", 45, extra...)
}

func (a *algorithms) nextCustom(fsleft, fsright string, width int, extra ...interface{}) *Element {
	name := fmt.Sprintf("Alg%02v", a.current)
	a.current++
	fn := files.MustLoadFunc("alg.go", name)
	lines := strings.Count(fn, "\n")
	var v *Element
	if fsleft != "" {
		v = LoadLinesCustom("alg.go", a.atLine, a.atLine+lines, fsleft)
	} else {
		v = LoadLines("alg.go", a.atLine, a.atLine+lines)
	}
	a.last = v // because we needed two slides with same algorithm
	// store for next call
	a.atLine += lines
	a.atLine += 2 // spacing to next func

	return TwoCol(
		v,
		Div(
			Shell(
				fmt.Sprintf("$ go test -benchmem -bench=Benchmark%s", name),
				fmt.Sprintf("testdata/%s_bench.html", strings.ToLower(name)),
			).With(func() any {
				if fsright != "" {
					return Attr("style", "font-size: "+fsright)
				}
				return nil
			}()),
			Wrap(extra...),
		),
		width,
	)
}
