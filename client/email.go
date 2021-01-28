package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) DomainSettingsFind(filter *RequestFilter) ([]model.DomainSettingsObject, error) {
	var data []interface{}
	_, err := c.Iterate(&data, &model.DomainSettingsObject{}, "email", "domainSettingsFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("domain settings find: %s", err)
	}

	out := []model.DomainSettingsObject{}
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.DomainSettingsObject))
	}

	return out, nil
}
