package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) ZonesFind(filter *RequestFilter) ([]model.ZoneObject, error) {
	var data []interface{}
	_, err := c.Iterate(&data, &model.ZoneObject{}, "dns", "zonesFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("zones find: %s", err)
	}

	out := []model.ZoneObject{}
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.ZoneObject))
	}

	return out, nil
}

func (c *ApiClient) ZoneConfigsFind(filter *RequestFilter) ([]model.ZoneConfigObject, error) {
	var data []interface{}
	_, err := c.Iterate(&data, &model.ZoneConfigObject{}, "dns", "zoneConfigsFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("zones find: %s", err)
	}

	out := []model.ZoneConfigObject{}
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.ZoneConfigObject))
	}

	return out, nil
}
