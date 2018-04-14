package polr_test

import (
	"testing"

	"time"

	"strings"

	"github.com/Seklfreak/polr-go"
	"github.com/stretchr/testify/assert"
)

func Test_Lookup(t *testing.T) {
	setup()

	tests := []struct {
		longUrl string
	}{
		{
			longUrl: "http://example.org/a",
		},
	}

	for _, test := range tests {
	RetryTest:
		result, err := client.Shorten(test.longUrl, "", false)

		if err != nil {
			if errP, ok := err.(*polr.PolrError); ok {
				if errP.Code == "QUOTE_EXCEEDED" {
					time.Sleep(time.Second * 10)
					goto RetryTest
				}
			}
		}

		if err != nil {
			t.Error(err)
			return
		}

		resultParts := strings.Split(result, "/")
		urlEnding := resultParts[len(resultParts)-1]

		lookupRes, err := client.Lookup(urlEnding, "")

		if err != nil {
			if errP, ok := err.(*polr.PolrError); ok {
				if errP.Code == "QUOTE_EXCEEDED" {
					time.Sleep(time.Second * 10)
					goto RetryTest
				}
			}
		}

		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, test.longUrl, lookupRes.LongURL)
	}
}
