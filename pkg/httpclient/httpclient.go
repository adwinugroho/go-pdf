package httpclient

import (
	"os"

	"github.com/go-resty/resty/v2"
)

func getClient(URL, APIKEY string) *resty.Client {
	client := resty.New().
		SetBaseURL(URL)

	client.SetDebug(false)

	if APIKEY != "" {
		client = client.SetQueryParam("token", APIKEY)
	}

	return client
}

func FileService() *resty.Client {
	baseURL := os.Getenv("fileservice_url")
	return getClient(baseURL, "")
}
