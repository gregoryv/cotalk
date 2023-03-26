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

// NewLetterChallenge returns a problem defined as
//
//   - get all letters using a list of requests
//   - in the given order
//
// letters should be a space separated string of letters
func NewLetterChallenge(letters string) *OrderedLetters {
	return &OrderedLetters{
		exp: letters,
	}
}

type OrderedLetters struct {
	url string
	exp string
}

func (p *OrderedLetters) Server() *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		<-time.After(10 * time.Millisecond)
		w.Write([]byte(r.URL.Path[1:]))
	}
	srv := httptest.NewServer(http.HandlerFunc(handler))
	p.url = srv.URL
	return srv
}

func (p *OrderedLetters) Solve(alg Algorithm) error {
	// create the workload, ie. requests
	words := strings.Split(p.exp, " ")
	work := make([]*http.Request, len(words))
	for i, word := range words {
		work[i], _ = http.NewRequest("GET", p.url+"/"+word, http.NoBody)
	}

	// run the algorithm
	result := alg(work)

	// verify the result
	return p.verify(work, result)
}

func (p *OrderedLetters) verify(work []*http.Request, result []*http.Response) error {
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
