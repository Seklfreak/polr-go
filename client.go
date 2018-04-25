package polr

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	pkgVersion = "0.0.1"
	userAgent  = "polr-go/" + pkgVersion
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type PolrError struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"error_code"`
	Message    string `json:"error"`
}

func (e PolrError) Error() string {
	return e.Code + ": " + e.Message
}

// Polr
type Polr struct {
	// HTTP client
	client httpClient

	// API Base URL
	baseURL *url.URL

	// User agent
	UserAgent string

	// ApiKey to authenticate
	ApiKey string

	credentials struct {
		appID  string
		appKey string
	}
}

// Returns a new Polr API Client
func New(baseUrl, apiKey string, hClient httpClient) (*Polr, error) {
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl += "/"
	}
	baseUrl += "v2/"

	baseUrlParsed, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	if hClient == nil {
		hClient = http.DefaultClient
	}

	c := &Polr{}
	c.client = hClient
	c.baseURL = baseUrlParsed
	c.UserAgent = userAgent
	c.ApiKey = apiKey

	return c, nil
}

// creates a new request
func (c *Polr) newRequest(method, path string, body []byte) (*http.Request, error) {
	rel, pErr := url.Parse(path)
	if pErr != nil {
		return nil, pErr
	}

	uri := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, uri.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set up header
	req.Header.Add("User-Agent", c.UserAgent)

	// set key and response_type as GET parameter
	queries, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return nil, err
	}
	queries.Set("key", c.ApiKey)
	queries.Set("response_type", "json")
	req.URL.RawQuery = queries.Encode()

	return req, nil
}

// exectures a request
func (c *Polr) do(req *http.Request) ([]byte, error) {
	resp, doErr := c.client.Do(req)
	if doErr != nil {
		return nil, doErr
	}

	defer resp.Body.Close()

	b, rErr := ioutil.ReadAll(resp.Body)
	if rErr != nil {
		return nil, rErr
	}

	if resp.StatusCode != http.StatusOK {
		var polrE PolrError
		uErr := json.Unmarshal(b, &polrE)
		if uErr != nil {
			return nil, uErr
		}
		return nil, polrE
	}

	return b, nil
}

// creates and executes a request
func (c *Polr) makeRequest(method, path string, params map[string]interface{}) ([]byte, error) {
	b, mErr := json.Marshal(params)
	if mErr != nil {
		return nil, mErr
	}

	req, reqErr := c.newRequest(method, path, b)
	if reqErr != nil {
		return nil, reqErr
	}

	resp, doErr := c.do(req)
	if doErr != nil {
		return nil, doErr
	}

	return resp, nil
}
