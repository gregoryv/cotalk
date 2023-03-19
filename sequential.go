package cotalk

import (
	"net/http"
)

func Sequential(work []*http.Request) (result []*http.Response) {
	for _, r := range work {
		resp, _ := http.DefaultClient.Do(r)
		result = append(result, resp)
	}
	return
}
