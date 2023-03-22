package cotalk

import (
	"context"
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

// Alg7 uses channel to synchronize responses with ordered result
func Alg7(work []*http.Request) (result []*http.Response) {
	type m struct {
		index int
		*http.Response
	}
	c := make(chan m)
	result = make([]*http.Response, len(work))
	for i, r := range work {
		go func(i int, lr *http.Request) {
			resp, _ := http.DefaultClient.Do(lr)
			c <- m{i, resp} // write to channel
		}(i, r)
	}
	for range work {
		v := <-c // read from channel
		result[v.index] = v.Response
	}
	return
}

// Alg8 uses channel to synchronize responses with ordered result
func Alg8(work []*http.Request) (result []*http.Response) {
	type m struct {
		index int
		*http.Response
	}
	c := make(chan m)
	defer close(c) // make sure you clean up when done
	result = make([]*http.Response, len(work))
	for i, r := range work {
		go func(i int, lr *http.Request) {
			resp, _ := http.DefaultClient.Do(lr)
			c <- m{i, resp} // write to channel
		}(i, r)
	}
	for range work {
		v := <-c // read from channel
		result[v.index] = v.Response
	}
	return
}

// Alg9 returns when all work is done or context is cancelled
func Alg9(ctx context.Context, work []*http.Request) (result []*http.Response) {
	type m struct {
		index int
		*http.Response
	}
	c := make(chan m)
	complete := make(chan struct{})
	defer close(c) // make sure you clean up when done
	result = make([]*http.Response, len(work))
	go func() {
		defer close(complete)
		for i, r := range work {
			go func(i int, lr *http.Request) {
				resp, _ := http.DefaultClient.Do(lr)
				c <- m{i, resp} // write to channel
			}(i, r)
		}
		for range work {
			v := <-c // read from channel
			result[v.index] = v.Response
		}
	}()
	select {
	case <-ctx.Done():
		// interrupted
	case <-complete:
	}
	return
}

// Alg10 returns when all work is done or context is cancelled
func Alg10(ctx context.Context, work []*http.Request) (result []*http.Response) {
	type m struct {
		index int
		*http.Response
	}
	c := make(chan m)
	complete := make(chan struct{})
	defer close(c) // make sure you clean up when done
	result = make([]*http.Response, len(work))
	go func() {
		defer close(complete)
		for i, r := range work {
			go func(i int, lr *http.Request) {
				resp, _ := http.DefaultClient.Do(lr)
				select {
				case <-ctx.Done():
				default:
					c <- m{i, resp} // write to channel
				}
			}(i, r)
		}
		for range work {
			v := <-c // read from channel
			result[v.index] = v.Response
		}
	}()
	select {
	case <-ctx.Done():
		// interrupted
	case <-complete:
	}
	return
}
