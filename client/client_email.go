package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

type Email struct {
	c *ApiClient
}

func (e *Email) DomainSettingsFind(filter *RequestFilter) ([]model.DomainSettingsObject, error) {
	var data []interface{}
	_, err := e.c.Find(&data, model.DomainSettingsObject{}, "email", "domainSettingsFind", filter, 0)

	if err != nil {
		return nil, fmt.Errorf("domain settings find: %s", err)
	}

	var out []model.DomainSettingsObject
	for i := 0; i < len(data); i++ {
		out = append(out, data[i].(model.DomainSettingsObject))
	}

	return out, nil
}
