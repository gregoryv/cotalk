package cotalk

import (
	"embed"
	"fmt"
	"log"
	"os/exec"
	"strings"

	. "github.com/gregoryv/web"
)

func Presentation() *deck {
	d := newDeck()
	d.Title = "Go concurrency"
	d.Styles = append(d.Styles,
		themeOldstyle(),
		highlightColors(),
	)

	nav := &navbar{current: 1}

	d.Slide(
		H1("Concurrency design in Go"),
		Span("2023 by Gregory Vinčić"),
		P(`Discussion and examples of using concepts related to
		concurrent design.`),

		Ul(Class("group"), Li("go routines"), Li("channels")),
		Ul(Class("group"), Li("package context"), Li("package sync "), Li("go test bench")),
		Ul(Class("group"), Li("examples and training")),
		Br(Attr("clear", "all")),
		nav,
	)

	d.Slide(H2("Background and history"),
		Img(Class("left"), Src("th_small.jpg")), Br(),
		A(Href("http://www.usingcsp.com/cspbook.pdf"),
			"Communicating Sequential Processes",
		),
		" C. A. R. Hoare, 1985",
		Br(Attr("clear", "all")),

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
		godoc("sync"),
		nav,
	)

	d.Slide(H2("package context"),
		Pre(highlightGoDoc(`
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

		godoc("context"),
		nav,
	)

	d.Slide(H2("go test -bench"),
		load("ex07/run.go"),
		load("ex07/run_test.go"),
		shell(
			"$ cd ex07; go test -benchmark -bench .",
			"testdata/sequential_benchmark.html",
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
		load("ex10/run_test.go"),
		nav,
	)

	d.Slide(H2("Sequential"),
		P("Simple implementation though very low performance"),
		srcTest(10),
		nav,
	)

	// Concurrent
	// ----------------------------------------
	d.Slide(H2("Concurrent"),
		P("This algorithm uses the sync.WaitGroup to wait for all requests to complete; however it has several bugs."),
		srcTest(20),
		nav,
	)

	d.Slide(H2("Concurrent"),
		P("Bug 2; Unprotected write to slice"),
		srcTest(40),
		nav,
	)

	d.Slide(H2("Concurrent"),
		P("Bug 1; you cannot assume go routines start immediately."),
		srcTest(30),
		P("You might get a different result; but why does it fail?"),
		nav,
	)

	d.Slide(H2("Concurrent"),
		P("Using channels"),
		load("ex50/run.go"),
		nav,
	)

	nav.max = len(d.Slides)
	return d
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
	v = highlight(v)
	return Div(
		Class("srcfile"),
		Pre(Code(v)),
	)
}

func godoc(pkg string) *Element {
	out, _ := exec.Command("go", "doc", "-short", pkg).Output()
	v := string(out)
	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlightGoDoc(v)
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

//go:embed ex* testdata
var assets embed.FS
