package ex50

import (
	"net/http"
)

func Run(work []*http.Request) (result []*http.Response) {
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
