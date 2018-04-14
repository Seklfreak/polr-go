# polr-go

[![Build Status](https://travis-ci.org/Seklfreak/polr-go.svg?branch=master)](https://travis-ci.org/Seklfreak/polr-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Seklfreak/polr-go)](https://goreportcard.com/report/github.com/Seklfreak/polr-go)

```go
client, err = polr.New("https://demo.polr.me/api/", "api-key", nil)
if err != nil {
    panic(err)
}

shortUrl, err := client.Shorten("https://example.org", "", "")
if err != nil {
    panic(err)
}

fmt.Println("Short URL:", shortUrl)
```