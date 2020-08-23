package mocking

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type httpClient struct {
	mock.Mock
}

func NewHTTPClient() *httpClient {
	return new(httpClient)
}

func (mock *httpClient) Get(url string) (*http.Response, error) {
	args := mock.Called(url)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*http.Response), args.Error(1)
}

func (mock *httpClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	args := mock.Called(url, contentType, body)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*http.Response), args.Error(1)
}

func MakeSuccessResponse(response string) *http.Response {
	return &http.Response{
		StatusCode: 200,
		Status:     "200 OK",
		Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		Header:     make(http.Header),
	}
}

func MakeNotFoundResponse() *http.Response {
	return &http.Response{
		StatusCode: 404,
		Status:     "404 NOT FOUND",
		Body:       ioutil.NopCloser(bytes.NewBufferString("Not Found")),
		Header:     make(http.Header),
	}
}
