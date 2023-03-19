package cotalk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

type Algorithm func(work []*http.Request) (result []*http.Response)

func CheckAlgorithm(fn Algorithm) error {
	minTaskDuration := 10 * time.Millisecond
	handler := func(w http.ResponseWriter, r *http.Request) {
		<-time.After(minTaskDuration)
	}
	srv := httptest.NewServer(http.HandlerFunc(handler))
	defer srv.Close()

	var (
		numRequests = 30
		work        = createWork(srv.URL, numRequests)
		result      = fn(work)
	)
	return complete(work, result)
}

func createWork(url string, n int) []*http.Request {
	all := make([]*http.Request, n)
	for i, _ := range all {
		all[i], _ = http.NewRequest("GET", url, http.NoBody)
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
