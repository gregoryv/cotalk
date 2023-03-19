package cotalk

import (
	"net/http"
	"sync"
)

// Alg1 solves the work sequentially
func Alg1(work []*http.Request) (result []*http.Response) {
	for _, r := range work {
		resp, _ := http.DefaultClient.Do(r)
		result = append(result, resp)
	}
	return
}

// Alg2 uses sync.WaitGroup to wait for all responses
func Alg2(work []*http.Request) (result []*http.Response) {
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

// Alg3 fixes reference problem inside loop
func Alg3(work []*http.Request) (result []*http.Response) {
	var wg sync.WaitGroup
	for _, r := range work {
		wg.Add(1)
		go func(lr *http.Request) {
			// use local argument
			resp, _ := http.DefaultClient.Do(lr)
			result = append(result, resp)
			wg.Done()
		}(r) // make a copy of pointer with argument
	}
	wg.Wait()
	return
}

// Alg4 synchronizes writes accross go routines
func Alg4(work []*http.Request) (result []*http.Response) {
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

// Alg5 fix order
func Alg5(work []*http.Request) (result []*http.Response) {
	var wg sync.WaitGroup
	var m sync.Mutex
	result = make([]*http.Response, len(work))
	for i, r := range work {
		wg.Add(1)
		go func(i int, lr *http.Request) {
			resp, _ := http.DefaultClient.Do(lr)

			// protect result
			m.Lock()
			result[i] = resp
			m.Unlock()

			wg.Done()
		}(i, r)
	}
	wg.Wait()
	return
}

// Alg6 uses channel to synchronize responses
func Alg6(work []*http.Request) (result []*http.Response) {
	c := make(chan *http.Response)
	for _, r := range work {
		go func(lr *http.Request) {
			resp, _ := http.DefaultClient.Do(lr)
			c <- resp // write to channel
		}(r)
	}
	for range work {
		resp := <-c // read from channel
		result = append(result, resp)
	}
	return
}
