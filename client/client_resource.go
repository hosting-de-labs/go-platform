package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

type Resource struct {
	c *ApiClient
}

func (r *Resource) ResourceVirtualMachineHostsFind(filter *RequestFilter) ([]model.VirtualMachineHostObject, error) {
	var data []interface{}
	_, err := r.c.Iterate(&data, &model.VirtualMachineHostObject{}, "resource", "virtualMachineHostsFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("domain settings find: %s", err)
	}

	var out []model.VirtualMachineHostObject
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.VirtualMachineHostObject))
	}

	return out, nil
}
