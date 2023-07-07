package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	client http.Client
}

type Response struct {
	statusCode int
	content    []byte
}

type OptionFunc func(r *http.Request)

type Reader interface {
	Read() io.Reader
}

type FormData map[string]string

func (data FormData) Read() io.Reader {
	var postData []string
	for key, val := range data {
		postData = append(postData, key+"="+val)
	}
	return strings.NewReader(strings.Join(postData, "&"))
}

type Json map[string]interface{}

func (j Json) Read() io.Reader {
	s, _ := json.Marshal(j)
	return strings.NewReader(string(s))
}

func (c Client) Post(url string, data Reader, opts ...OptionFunc) (*Response, error) {
	req, err := http.NewRequest("POST", url, data.Read())

	switch data.(type) {
	case Json:
		SetHeader("Content-Type", "application/json")(req)
	case FormData:
		SetHeader("Content-Type", "application/x-www-form-urlencoded")(req)
	}

	for _, opt := range opts {
		opt(req)
	}

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return c.response(response)
}

func (c Client) Get(url string, opts ...OptionFunc) (*Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	for _, opt := range opts {
		opt(req)
	}
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return c.response(response)
}

func SetHeaders(headers map[string]string) OptionFunc {
	return func(r *http.Request) {
		for key, value := range headers {
			r.Header.Add(key, value)
		}
	}
}

func SetHeader(key, value string) OptionFunc {
	return func(r *http.Request) {
		r.Header.Add(key, value)
	}
}

func (c Client) response(resp *http.Response) (*Response, error) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	return &Response{
		statusCode: resp.StatusCode,
		content:    body,
	}, err
}

func (r Response) StatusCode() int {
	return r.statusCode
}

func (r Response) Content() string {
	return string(r.content)
}

func (r Response) Unmarshal(v any) error {
	return json.Unmarshal(r.content, v)
}

func EntityToJson[T interface{}](data T) (Json, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	jsonObj := Json{}
	if err := json.Unmarshal(jsonBytes, &jsonObj); err != nil {
		// 处理错误
		return nil, err
	}
	return jsonObj, nil
}
