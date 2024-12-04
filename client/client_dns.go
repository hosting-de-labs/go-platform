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
	_, err := c.c.Find(&data, &model.ZoneConfigObject{}, "dns", "zoneConfigsFind", filter, 0)
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
	_, err := c.c.Find(&data, &model.ZoneObject{}, "dns", "zonesFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("zones find: %s", err)
	}

	var out []model.ZoneObject
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.ZoneObject))
	}

	return out, nil
}

type ZoneResponse struct {
	Response model.ZoneObject
	EmptyResponse
}

type ZoneCreateRecreateRequest struct {
	ZoneConfig              model.ZoneConfigObject `json:"zoneConfig,omitempty"`
	Records                 []model.RecordObject   `json:"records,omitempty"`
	NameserverSetID         string                 `json:"nameserverSetId,omitempty"`
	UseDefaultNameserverSet bool                   `json:"useDefaultNameserverSet,omitempty"`
}

func (c *Dns) ZoneCreate(zreq ZoneCreateRecreateRequest) (*ZoneResponse, error) {
	resp, err := c.c.RawRequest("dns", "zoneCreate", zreq)
	if err != nil {
		return nil, fmt.Errorf("zone create: %s", err)
	}

	var r ZoneResponse
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, fmt.Errorf("zone create: %s", err)
	}

	return &r, nil
}

func (c *Dns) ZoneRecreate(zreq ZoneCreateRecreateRequest) (*ZoneResponse, error) {
	resp, err := c.c.RawRequest("dns", "zoneRecreate", zreq)
	if err != nil {
		return nil, fmt.Errorf("zone recreate: %s", err)
	}

	var r ZoneResponse
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, fmt.Errorf("zone recreate: %s", err)
	}

	return &r, nil
}

type ZoneUpdateRequest struct {
	ZoneConfig      model.ZoneConfigObject `json:"zoneConfig,omitempty"`
	RecordsToAdd    []model.RecordObject   `json:"recordsToAdd,omitempty"`
	RecordsToModify []model.RecordObject   `json:"recordsToModify,omitempty"`
	RecordsToDelete []model.RecordObject   `json:"recordsToDelete,omitempty"`
}

func (c *Dns) ZoneUpdate(zreq ZoneUpdateRequest) (*ZoneResponse, error) {
	resp, err := c.c.RawRequest("dns", "zoneUpdate", zreq)
	if err != nil {
		return nil, fmt.Errorf("zone update: %s", err)
	}

	var r ZoneResponse
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, fmt.Errorf("zone update: %s", err)
	}

	return &r, nil
}

type ZoneDeleteRequest struct {
	ZoneConfigID string `json:"zoneConfigId,omitempty"`
	ZoneName     string `json:"zoneName,omitempty"`
}

func (c *Dns) ZoneDelete(zreq ZoneDeleteRequest) (*EmptyResponse, error) {
	resp, err := c.c.RawRequest("dns", "zoneDelete", zreq)
	if err != nil {
		return nil, fmt.Errorf("zone delete: %s", err)
	}

	var r EmptyResponse
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		return nil, fmt.Errorf("zone delete: %s", err)
	}

	return &r, nil
}
