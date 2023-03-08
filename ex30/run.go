package ex30

import (
	"net/http"
	"sync"
)

func Run(work []*http.Request) (result []*http.Response) {
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
