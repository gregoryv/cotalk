package cotalk

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

// Setup returns running test server and the letters problem to solve.
// The problem can be solved multiple times.
func Setup() (*httptest.Server, *LettersProblem) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		<-time.After(10 * time.Millisecond)
		w.Write([]byte(r.URL.Path[1:]))
	}
	srv := httptest.NewServer(http.HandlerFunc(handler))
	prob := &LettersProblem{
		srv: srv,
		exp: Letters,
	}
	return srv, prob
}

const Letters = "0 1 2 3 4 5 6 7 8 9 a b c d e f"

type LettersProblem struct {
	srv *httptest.Server
	exp string
}

func (p *LettersProblem) Solve(alg Algorithm) error {
	// create the workload, ie. requests
	words := strings.Split(p.exp, " ")
	work := make([]*http.Request, len(words))
	for i, word := range words {
		work[i], _ = http.NewRequest("GET", p.srv.URL+"/"+word, http.NoBody)
	}

	// run the algorithm
	result := alg(work)

	// verify the result
	return p.complete(work, result)
}

func (p *LettersProblem) complete(work []*http.Request, result []*http.Response) error {
	got := make([]string, 0, len(result))
	for _, resp := range result {
		if resp != nil {
			var buf bytes.Buffer
			io.Copy(&buf, resp.Body)
			resp.Body.Close()
			got = append(got, buf.String())
		} else {
			got = append(got, " ") // makes it easier to see
		}
	}
	if got := strings.Join(got, " "); p.exp != got {
		return fmt.Errorf("\nexp: %s\ngot: %s", p.exp, got)
	}
	return nil
}
