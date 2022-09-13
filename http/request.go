package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var client = &http.Client{}

// Request
type Request struct {
	URL          string `json:"url"`
	ResponseBody bool   `json:"response_body"`
	ResponseCode int    `json:"response_code"`
}

// Get
func (r Request) Get() (result []byte, err error) {
	request, errs := http.NewRequest("GET", r.URL, nil)
	if err != nil {
		err = errs
		return
	}

	if response, errs := client.Do(request); err != nil {
		err = errs
		return
	} else {
		defer response.Body.Close()
		if r.ResponseBody {
			result, err = ioutil.ReadAll(response.Body)
		}
		if response.StatusCode != r.ResponseCode {
			err = fmt.Errorf("status code is %v", response.StatusCode)
		}
	}
	return
}
