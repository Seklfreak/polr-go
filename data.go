package polr

import (
	"encoding/json"
	"net/url"
)

type DataResultContainer struct {
	Action string     `json:"action"`
	Result DataResult `json:"result"`
}

type DataResult struct {
	URLEnding string `json:"url_ending"`
	Data      []struct {
		Label  string `json:"label"`
		Clicks int    `json:"clicks"`
		X      string `json:"x"`
		Y      int    `json:"y"`
	} `json:"data"`
}

// Shorten shortens an url.
func (c *Polr) Data(urlEnding, statsType, leftBound, rightBound string) (re *DataResult, err error) {
	p := url.Values{}
	p.Set("url_ending", urlEnding)
	p.Set("stats_type", statsType)

	if leftBound != "" {
		p.Set("left_bound", leftBound)
	}

	if rightBound != "" {
		p.Set("right_bound", rightBound)
	}

	resp, err := c.makeRequest("POST", "data/link?"+p.Encode(), nil)
	if err != nil {
		return nil, err
	}

	var reContainer DataResultContainer
	uErr := json.Unmarshal(resp, &reContainer)
	if uErr != nil {
		return nil, uErr
	}

	return &reContainer.Result, nil
}
