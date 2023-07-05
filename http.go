package c2

// http client defines here
type HTTPClient struct {
}

var defaultHTTPClient *HTTPClient

func init() {
	defaultHTTPClient = &HTTPClient{}
}
