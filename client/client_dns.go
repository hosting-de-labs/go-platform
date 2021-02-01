package client

import (
	"encoding/json"
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

type Dns struct {
	c *ApiClient
}

func (c *Dns) ZoneConfigsFind(filter *RequestFilter) ([]model.ZoneConfigObject, error) {
	var data []interface{}
	_, err := c.c.Iterate(&data, &model.ZoneConfigObject{}, "dns", "zoneConfigsFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("zones find: %s", err)
	}

	var out []model.ZoneConfigObject
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.ZoneConfigObject))
	}

	return out, nil
}

func (c *Dns) ZonesFind(filter *RequestFilter) ([]model.ZoneObject, error) {
	var data []interface{}
	_, err := c.c.Iterate(&data, &model.ZoneObject{}, "dns", "zonesFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("zones find: %s", err)
	}

	var out []model.ZoneObject
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.ZoneObject))
	}

	return out, nil
}

type ZoneUpdateRequest struct {
	ZoneConfig      model.ZoneConfigObject `json:"zoneConfig,omitempty"`
	RecordsToAdd    []model.RecordObject   `json:"recordsToAdd,omitempty"`
	RecordsToModify []model.RecordObject   `json:"recordsToModify,omitempty"`
	RecordsToDelete []model.RecordObject   `json:"recordsToDelete,omitempty"`
}

type ZoneUpdateResponse struct {
	Response model.ZoneObject

	Metadata Metadata
	Status   string

	Errors   []Error
	Warnings []Error
}

func (c *Dns) ZoneUpdate(zreq ZoneUpdateRequest) (*ZoneUpdateResponse, error) {
	resp, err := c.c.Update("dns", "zoneUpdate", zreq)
	if err != nil {
		return nil, fmt.Errorf("update: %s", err)
	}

	var r ZoneUpdateResponse
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, fmt.Errorf("zone update: %s", err)
	}

	return &r, nil
}
