package polr_test

import (
	"testing"

	"strconv"
	"time"

	"github.com/Seklfreak/polr-go"
	"github.com/stretchr/testify/assert"
)

func Test_Shorten(t *testing.T) {
	setup()

	endingSuffix := strconv.FormatInt(time.Now().UnixNano(), 10)

	tests := []struct {
		longUrl       string
		secret        bool
		customEnding  string
		shouldContain []string
	}{
		{
			longUrl: "http://example.org/a",
			shouldContain: []string{
				"http://demo.polr.me/",
			},
		},
		{
			longUrl: "http://example.org/b",
			secret:  true,
			shouldContain: []string{
				"http://demo.polr.me/",
			},
		},
		{
			longUrl:      "http://example.org/c",
			customEnding: "example" + endingSuffix,
			shouldContain: []string{
				"http://demo.polr.me/",
				"example" + endingSuffix,
			},
		},
		{
			longUrl:      "http://example.org/c",
			secret:       true,
			customEnding: "secret-example" + endingSuffix,
			shouldContain: []string{
				"http://demo.polr.me/",
				"secret-example" + endingSuffix,
			},
		},
	}

	for _, test := range tests {
	RetryTest:
		result, err := client.Shorten(test.longUrl, test.customEnding, test.secret)

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

		for _, shouldContain := range test.shouldContain {
			assert.Contains(t, result, shouldContain)
		}
	}
}
