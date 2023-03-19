package cotalk

import (
	"net/http"
	"sync"
)

func Sequential(work []*http.Request) (result []*http.Response) {
	for _, r := range work {
		resp, _ := http.DefaultClient.Do(r)
		result = append(result, resp)
	}
	return
}

func ConcurrentWaitGroup(work []*http.Request) (result []*http.Response) {
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

func ConcurrentWaitGroup_FixReference(work []*http.Request) (result []*http.Response) {
	var wg sync.WaitGroup
	for _, r := range work {
		wg.Add(1)
		go func(lr *http.Request) {
			resp, _ := http.DefaultClient.Do(lr) // use argument
			result = append(result, resp)
			wg.Done()
		}(r) // make a copy of pointer with argument
	}
	wg.Wait()
	return
}
