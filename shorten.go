package polr

import (
	"encoding/json"
	"net/url"
)

type ShortenResult struct {
	Action string `json:"action"`
	Result string `json:"result"`
}

// Shorten shortens an url.
func (c *Polr) Shorten(longUrl, customEnding string, secret bool) (shortUrl string, err error) {
	p := url.Values{}
	p.Set("url", longUrl)

	if secret {
		p.Set("is_secret", "true")
	}

	if customEnding != "" {
		p.Set("custom_ending", customEnding)
	}

	resp, err := c.makeRequest("POST", "action/shorten?"+p.Encode(), nil)
	if err != nil {
		return "", err
	}

	re := &ShortenResult{}

	uErr := json.Unmarshal(resp, &re)
	if uErr != nil {
		return "", uErr
	}

	return re.Result, nil
}
