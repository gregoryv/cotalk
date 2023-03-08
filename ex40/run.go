package ex40

import (
	"net/http"
	"sync"
)

func Run(work []*http.Request) (result []*http.Response) {
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, r := range work {
		wg.Add(1)
		go func(lr *http.Request) {
			resp, _ := http.DefaultClient.Do(lr)

			// protect result
			m.Lock()
			result = append(result, resp)
			m.Unlock()

			wg.Done()
		}(r)
	}
	wg.Wait()
	return
}
