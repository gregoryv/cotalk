package cotalk

import (
	"embed"
	"fmt"
	"log"
	"os/exec"
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

	nav := &navbar{current: 2}
	alg := &algorithms{current: 1, atLine: 9} // 8 where first func starts

	d.Slide(
		Br(),
		A(Href("#2"), Img(Src("cotalk.png"))),
		Br(), Br(), Br(),
		Span("2023 by Gregory Vinčić"),
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
		Img(Class("left"), Src("th_small.jpg")), Br(),
		P("Sir Charles Antony Richard Hoare (",
			A(Href("https://en.wikipedia.org/wiki/Tony_Hoare"), "Tony Hoare"),

			`). Born 1934 in Sri Lanka, studied at Oxford and in
			Moscow. His research spanned program correctness, sorting
			and programming languages. His work is freely accessible
			online and the Go developers uses his concept of <em>channels</em>
			in the language.`),

		Ul(
			Li(
				A(Href("https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf"), "Communicating Sequential Processes (CSP), paper 1978"),
			),
			Li(
				A(Href("http://www.usingcsp.com/cspbook.pdf"), "CSP, book 1985"),
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

		loadFunc("../alg_test.go", "BenchmarkAlg1"),

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

	d.Slide(H2("Summary"),

		Ul(
			Li("concurrency is difficult to get right even in Go"),
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

//go:embed ex* testdata *.go
var assets embed.FS

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
	name := fmt.Sprintf("Alg%v", a.current)
	a.current++
	fn := files.MustLoadFunc("../alg.go", name)
	lines := strings.Count(fn, "\n")
	v := files.MustLoadLines("../alg.go", a.atLine, a.atLine+lines)
	a.atLine = a.atLine + lines + 2

	v = strings.ReplaceAll(v, "\t", "    ")
	v = highlight(v)
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
