package client

import "github.com/go-resty/resty/v2"

type ApiClient struct {
	url   string
	token string
	limit int

	client *resty.Client
}

func NewApiClient(url string, token string, limit int) *ApiClient {
	c := new(ApiClient)

	c.url = url
	c.token = token

	c.limit = 25
	if limit > 0 {
		c.limit = limit
	}

	c.client = resty.New()

	return c
}
