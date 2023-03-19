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

type Algorithm func(work []*http.Request) (result []*http.Response)

func Solve(fn Algorithm) error {
	srv := SetupServer()
	defer srv.Close()
	work := createWork(srv.URL)

	// do the work
	result := fn(work)

	// verify, todo move this to test as it depends on the requirements
	return complete(work, result)
}

func SetupServer() *httptest.Server {
	minTaskDuration := 10 * time.Millisecond
	handler := func(w http.ResponseWriter, r *http.Request) {
		<-time.After(minTaskDuration)
		// echo the requested path
		w.Write([]byte(r.URL.Path[1:]))
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}

const exp = "0 1 2 3 4 5 6 7 8 9 a b c d e f"

func createWork(url string) []*http.Request {
	words := strings.Split(exp, " ")
	all := make([]*http.Request, len(words))

	for i, word := range words {
		all[i], _ = http.NewRequest("GET", url+"/"+word, http.NoBody)
	}
	return all
}

func complete(work []*http.Request, result []*http.Response) error {
	if v := len(result); v != len(work) {
		return fmt.Errorf("%v/%v incomplete", v, len(work))
	}
	if err := allOk(result); err != nil {
		return err
	}
	if err := checkOrder(result); err != nil {
		return err
	}
	return nil
}

func allOk(result []*http.Response) error {
	ok := 0
	missing := 0
	for _, r := range result {
		if r == nil {
			missing++
		}
		if r != nil && r.StatusCode < 400 {
			ok++
		}
	}
	if v := len(result) - ok; v > 0 {
		return fmt.Errorf("%v failed %v missing", v, missing)
	}
	return nil
}

func checkOrder(result []*http.Response) error {
	words := make([]string, 0, len(result))
	for _, resp := range result {
		var buf bytes.Buffer
		io.Copy(&buf, resp.Body)
		resp.Body.Close()
		words = append(words, buf.String())
	}
	got := strings.Join(words, " ")
	if exp != got {
		return fmt.Errorf("\nexp: %s\ngot: %s", exp, got)
	}
	return nil
}
