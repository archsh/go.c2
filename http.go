package c2

import (
	"fmt"
	"io"
	"net/http"
)

// http client defines here
type HTTPClient struct {
	client *http.Client
}

func (c HTTPClient) Get(s string, w io.Writer) (int64, error) {
	if resp, e := c.client.Get(s); nil != e {
		return 0, e
	} else {
		if resp.StatusCode/100 > 3 {
			return 0, fmt.Errorf("%s", resp.Status)
		}
		defer func() { _ = resp.Body.Close() }()
		return io.Copy(w, resp.Body)
	}
}

func HttpGet(s string, w io.Writer) (int64, error) {
	return defaultHTTPClient.Get(s, w)
}

var defaultHTTPClient *HTTPClient

func init() {
	defaultHTTPClient = &HTTPClient{client: http.DefaultClient}
}
