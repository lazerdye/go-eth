package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type HttpError struct {
	err        string
	statusCode int
}

func NewHttpError(err string, statusCode int) error {
	return &HttpError{err, statusCode}
}

func (e *HttpError) Error() string {
	return e.err
}

func (e *HttpError) StatusCode() int {
	return e.statusCode
}

type HttpClient struct {
	client http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (h *HttpClient) Get(ctx context.Context, reqUrl string, params url.Values, data interface{}) error {
	return h.Do(ctx, "GET", reqUrl, params, nil, nil, data)
}

func (h *HttpClient) GetWithHeader(ctx context.Context, reqUrl string, params url.Values, header *http.Header, data interface{}) error {
	return h.Do(ctx, "GET", reqUrl, params, header, nil, data)
}

func (h *HttpClient) PostWithHeader(ctx context.Context, reqUrl string, params url.Values, header *http.Header, data interface{}) error {
	return h.Do(ctx, "POST", reqUrl, params, header, nil, data)
}

func (h *HttpClient) Post(ctx context.Context, reqUrl string, params url.Values, header *http.Header, body interface{}, data interface{}) error {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return h.Do(ctx, "POST", reqUrl, params, header, bytes.NewBuffer(bodyStr), data)
}

func (h *HttpClient) Do(ctx context.Context, method string, reqUrl string, params url.Values, header *http.Header, body io.Reader, data interface{}) error {
	base, err := url.Parse(reqUrl)
	if err != nil {
		return err
	}
	if params != nil {
		base.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, base.String(), body)
	if err != nil {
		return err
	}
	if header != nil {
		req.Header = *header
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return NewHttpError(fmt.Sprintf("Unexpected status code: %d", resp.StatusCode), resp.StatusCode)
	}

	log.Infof("Resp: %+v", resp)

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(data); err != nil {
		return err
	}

	return nil
}
