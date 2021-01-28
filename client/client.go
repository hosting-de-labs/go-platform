package client

type ApiClient struct {
	url   string
	token string
	limit int
}

func NewApiClient(url string, token string, limit int) *ApiClient {
	c := new(ApiClient)

	c.url = url
	c.token = token

	c.limit = 25
	if limit > 0 {
		c.limit = limit
	}

	return c
}
