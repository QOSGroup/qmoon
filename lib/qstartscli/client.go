package qstatcli

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"

	"net/http"
	"net/url"
)

const tmDefaultServer = "http://18.188.103.180:26657"

// 代码结构和部分实现的灵感来源于 https://github.com/google/go-github

// Option 可选参数
type option struct {
	host string
	// HTTP Client used to communicate with the API.
	// Base URL for API requests. Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	baseURL *url.URL
}
type SetOption func(options *option) error

// NewOption 创建可选参数
func NewOption(fs ...SetOption) (*option, error) {
	opt := &option{
		host: tmDefaultServer,
	}
	baseURL, err := url.Parse(opt.host)
	if err != nil {
		return nil, err
	}
	opt.baseURL = baseURL

	if fs != nil {
		for _, f := range fs {
			if err := f(opt); err != nil {
				return nil, err
			}
		}
	}

	return opt, nil
}

// SetOptionHost 设置可选参数host
func SetOptionHost(host string) SetOption {
	return func(opt *option) error {
		baseURL, err := url.Parse(host)
		if err != nil {
			return err
		}

		opt.host = host
		opt.baseURL = baseURL

		return nil
	}
}

// Client 客户端结构
type Client struct {
	host    string
	baseURL *url.URL
	client  *http.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Test int
	aaa  int
	//APIGroup *APIGroupService
}

type service struct {
	client *Client
}

// NewClient 创建tendermint Client
func NewClient(opt *option) *Client {
	c := &Client{
		host:    opt.host,
		baseURL: opt.baseURL,
	}

	c.common.client = c

	//c.APIGroup = (*APIGroupService)(&c.common)

	return c
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

	if c.host[:7] != "http://" {
		c.host = "http://" + c.host
	}

	baseURL, err := url.Parse(c.host)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(baseURL.Path, "/") {
		baseURL.Path = baseURL.Path + "/"
	}
	c.BaseURL = baseURL

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.token == "" {
		r, _, err := c.Login.Post(context.Background(), &LoginReq{
			Username: c.username,
			Password: c.password,
		})
		if err != nil {
			return nil, err
		}
		c.token = r.Token
	}
	req.Header.Set("X-Resthub-Token", c.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// Response is a GitHub API response. This wraps the standard http.Response
// returned from GitHub and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

type AuthError struct{}

func (*AuthError) Error() string {
	return "未登录"
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
//
// The error type will be *RateLimitError for rate limit exceeded errors,
// *AcceptedError for 202 Accepted status codes,
// and *TwoFactorAuthError for two-factor authentication errors.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	if c := r.StatusCode; c == http.StatusUnauthorized || c == http.StatusForbidden {
		return &AuthError{}
	}

	return nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}

		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			d, _ := ioutil.ReadAll(resp.Body)
			logrus.Debugf("[http] req:%+v, body:%s", req, string(d))
			err = json.Unmarshal(d, v)

			//err = json.NewDecoder(resp.Body).Decode(v)
			//if err == io.EOF {
			//	err = nil // ignore EOF errors caused by empty response body
			//}
			//
			if d, err := json.Marshal(v); err == nil {
				logrus.Infof("Client.Do req:%+v,  resp:%d.", req.URL.String(), response.StatusCode)
				logrus.Debugf("Client.Do req:%+v,  body:%s.", req, string(d))
			}
		}
	}

	return response, err
}
