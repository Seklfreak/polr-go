package polr_test

import (
	"github.com/Seklfreak/polr-go"
)

const (
	API_KEY  = "demo-admin"
	BASE_URL = "https://demo.polr.me/api/"
)

var (
	client *polr.Polr
)

func setup() {
	client, _ = polr.New(BASE_URL, API_KEY, nil)
}
