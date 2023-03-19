package ex20

import (
	"net/http"
	"sync"
)

func Run(work []*http.Request) (result []*http.Response) {
	var wg sync.WaitGroup
	for _, r := range work {
		wg.Add(1)
		go func() {
			resp, _ := http.DefaultClient.Do(r)
			result = append(result, resp)
			wg.Done()
		}()
	}
	wg.Wait()
	return
}
