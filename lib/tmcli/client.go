package tmcli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"reflect"

	"github.com/QOSGroup/qmoon/lib/tmcli/tmmock"
	"github.com/google/go-querystring/query"
	"github.com/tendermint/go-amino"
	tmltypes "github.com/tendermint/tendermint/rpc/lib/types"

	"net/http"
	"net/url"
)

var TmDefaultServer = "http://127.0.0.1:26657"

var cdc = amino.NewCodec()

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
		host: TmDefaultServer,
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

//func MakeCodec() *amino.Codec {
//	cdc := amino.NewCodec()
//
//	cdc.RegisterInterface((*account.Account)(nil), nil)
//	cdc.RegisterConcrete(&qosacc.QOSAccount{}, "qbase/account/QOSAccount", nil)
//
//	return cdc
//}

// Client 客户端结构
type Client struct {
	host    string
	baseURL *url.URL
	client  *http.Client
	cdc     *amino.Codec

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Test int
	aaa  int

	Status             *statusService
	AbciInfo           *abciInfoService
	ConsensusState     *consensusStateService
	DumpConsensusState *dumpConsensusStateService
	Genesis            *genesisService
	Health             *healthService
	NetInfo            *netInfoService
	NumUnconfirmedTxs  *numUnconfirmedTxsService
	AbciQuery          *abciQueryService
	Block              *blockService
	BlockResults       *blockResultsService
	Blockchain         *blockchainService
	BroadcastTxAsync   *broadcastTxAsyncService
	BroadcastTxCommit  *broadcastTxCommitService
	BroadcastTxSync    *broadcastTxSyncService
	Subscribe          *subscribeService
	Tx                 *txService
	TxSearch           *txSearchService
	UnconfirmedTxs     *unconfirmedTxsService
	Unsubscribe        *unsubscribeService
	Validators         *validatorsService
}

type service struct {
	client *Client
}

// NewClient 创建tendermint Client
func NewClient(opt *option) *Client {
	if opt == nil {
		opt, _ = NewOption()
	}

	c := &Client{
		host:    opt.host,
		baseURL: opt.baseURL,
		client:  http.DefaultClient,
		cdc:     cdc,
	}

	c.common.client = c

	c.Status = (*statusService)(&c.common)
	c.AbciInfo = (*abciInfoService)(&c.common)
	c.ConsensusState = (*consensusStateService)(&c.common)
	c.DumpConsensusState = (*dumpConsensusStateService)(&c.common)
	c.Genesis = (*genesisService)(&c.common)
	c.Health = (*healthService)(&c.common)
	c.NetInfo = (*netInfoService)(&c.common)
	c.NumUnconfirmedTxs = (*numUnconfirmedTxsService)(&c.common)
	c.AbciQuery = (*abciQueryService)(&c.common)
	c.Block = (*blockService)(&c.common)
	c.BlockResults = (*blockResultsService)(&c.common)
	c.Blockchain = (*blockchainService)(&c.common)
	c.BroadcastTxAsync = (*broadcastTxAsyncService)(&c.common)
	c.BroadcastTxCommit = (*broadcastTxCommitService)(&c.common)
	c.BroadcastTxSync = (*broadcastTxSyncService)(&c.common)
	c.Subscribe = (*subscribeService)(&c.common)
	c.Tx = (*txService)(&c.common)
	c.TxSearch = (*txSearchService)(&c.common)
	c.UnconfirmedTxs = (*unconfirmedTxsService)(&c.common)
	c.Unsubscribe = (*unsubscribeService)(&c.common)
	c.Validators = (*validatorsService)(&c.common)

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

// newRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
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

	return req, nil
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
		return errors.New("没有授权")
	}

	return nil
}

// do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// do returns *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
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

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			var tmresp tmltypes.RPCResponse
			err = json.NewDecoder(resp.Body).Decode(&tmresp)
			if err != nil {
				if err == io.EOF {
					err = nil // ignore EOF errors caused by empty response body
				} else {
					return resp, err
				}
			}

			if tmresp.Error != nil {
				return resp, tmresp.Error
			}

			err = c.cdc.UnmarshalJSON(tmresp.Result, v)
			if err != nil {
				return resp, err
			}
		}
	}

	return resp, err
}

type testTmServer struct {
	s *httptest.Server
}

func (tts *testTmServer) Close() {
	tts.s.Close()
}

func (tts *testTmServer) URL() string {
	return tts.s.URL
}

func NewTestTmServer() *testTmServer {
	tts := &testTmServer{}

	tm := tmmock.NewTendermintMock()
	s := httptest.NewServer(tm)
	TmDefaultServer = s.URL

	tts.s = s

	return tts
}
