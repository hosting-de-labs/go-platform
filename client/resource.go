package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) ResourceVirtualMachineHostsFind(filter *RequestFilter) (*[]model.VirtualMachineHostObject, error) {
	resp, err := c.runRequest("resource", "virtualMachineHostsFind", filter, 0, 0)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	hosts := &model.VirtualMachineHostsResult{}

	err = json.Unmarshal(body, hosts)
	if err != nil {
		panic(err)
	}

	return &hosts.Response.Data, nil
}
