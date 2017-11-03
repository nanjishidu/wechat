package small

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpRequest struct {
	req *http.Request
}

func NewRequest(method, url string, body io.Reader) *HttpRequest {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("httplib NewRequest:", err)
	}
	return &HttpRequest{
		req: req,
	}
}

func HttpGet(url string, data ...io.Reader) *HttpRequest {
	if len(data) > 0 {
		return NewRequest("GET", url, data[0])
	}
	return NewRequest("GET", url, nil)
}
func (m *HttpRequest) DoRequest() (*http.Response, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(m.req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (m *HttpRequest) Bytes() ([]byte, error) {
	resp, err := m.DoRequest()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, errors.New("response body is null")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, err

}
func (m *HttpRequest) ToJson(v interface{}) error {
	b, err := m.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}
