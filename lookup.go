package polr

import (
	"encoding/json"
	"net/url"
)

type LookupResultContainer struct {
	Action string       `json:"action"`
	Result LookupResult `json:"result"`
}

type LookupResult struct {
	LongURL   string `json:"long_url"`
	CreatedAt struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"created_at"`
	Clicks    string `json:"clicks"`
	UpdatedAt struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"updated_at"`
}

// Shorten shortens an url.
func (c *Polr) Lookup(urlEnding, urlKey string) (re *LookupResult, err error) {
	p := url.Values{}
	p.Set("url_ending", urlEnding)

	if urlKey != "" {
		p.Set("url_key", urlKey)
	}

	resp, err := c.makeRequest("POST", "action/lookup?"+p.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var reContainer LookupResultContainer
	uErr := json.Unmarshal(resp, &reContainer)
	if uErr != nil {
		return nil, uErr
	}

	return &reContainer.Result, nil
}
