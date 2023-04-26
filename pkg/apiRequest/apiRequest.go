package apiRequest

import "net/http"

func Get(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	return resp
}