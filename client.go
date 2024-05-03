package each

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	ErrMissingCredential = errors.New("missing credential - please provide your API key")

	defaultBaseURL     = "https://api.eachlabs.ai/v1"
	defaultFlowBaseURL = "https://flows.eachlabs.ai/api/v1"

	defaultUserAgent = "each/go"
)

type BackendService string

const (
	BackendServiceFlow      BackendService = "flow"
	BackendServiceInference BackendService = "inference"
)

type options struct {
	credential  string
	baseURL     string
	flowBaseURL string
	httpClient  *http.Client
	userAgent   *string
}

type Client struct {
	options    *options
	httpClient *http.Client
}

type ClientOptions func(*options) error

func NewClient(opts ...ClientOptions) (*Client, error) {
	c := &Client{
		options: &options{
			httpClient:  http.DefaultClient,
			baseURL:     defaultBaseURL,
			flowBaseURL: defaultFlowBaseURL,
			userAgent:   &defaultUserAgent,
		},
	}

	var errs []error
	for _, o := range opts {
		if err := o(c.options); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	if c.options.credential == "" {
		return nil, errors.New("missing credential - please provide your API key")
	}

	c.httpClient = c.options.httpClient

	return c, nil
}

func WithCredential(credential string) ClientOptions {
	return func(o *options) error {
		o.credential = credential
		return nil
	}
}

func WithCredentialFromEnv() ClientOptions {
	return func(o *options) error {
		credential := os.Getenv("EACH_API_KEY")
		if credential == "" {
			return errors.New("missing credential - please provide your EACH_API_KEY env var")
		}

		o.credential = credential
		return nil
	}
}

func WithBaseURL(baseURL string) ClientOptions {
	return func(o *options) error {
		o.baseURL = baseURL
		return nil
	}
}

func WithFlowBaseURL(flowBaseURL string) ClientOptions {
	return func(o *options) error {
		o.flowBaseURL = flowBaseURL
		return nil
	}
}

func (e *Client) newRequest(ctx context.Context, backendType BackendService, method, path string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", e.options.flowBaseURL, path)
	if backendType == BackendServiceInference {
		url = fmt.Sprintf("%s%s", e.options.baseURL, path)
	}
	fmt.Println(url)
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create request failed w err: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", e.options.credential)
	if e.options.userAgent != nil {
		request.Header.Set("User-Agent", *e.options.userAgent)
	}

	return request, nil
}

func (e *Client) doRequest(ctx context.Context, req *http.Request, out interface{}) error {
	resp, err := e.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request failed w err: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("do request failed w status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(out)
}
