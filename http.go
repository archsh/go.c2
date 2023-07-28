package c2

// http client defines here
type HTTPClient struct {
}

func (c HTTPClient) Get(filename string, output string) error {
	return nil
}

func HttpGet(filename string, output string) error {
	return defaultHTTPClient.Get(filename, output)
}

var defaultHTTPClient *HTTPClient

func init() {
	defaultHTTPClient = &HTTPClient{}
}
