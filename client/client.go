package client

import "github.com/go-resty/resty/v2"

type ApiClient struct {
	Dns      Dns
	Email    Email
	Machine  Machine
	Resource Resource

	url   string
	token string
	limit int

	client *resty.Client
}

func NewApiClient(url string, token string, limit int) *ApiClient {
	c := &ApiClient{}

	c.url = url
	c.token = token

	c.limit = 25
	if limit > 0 {
		c.limit = limit
	}

	c.client = resty.New()

	c.Dns.c = c
	c.Email.c = c
	c.Machine.c = c
	c.Resource.c = c

	return c
}
