package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) MachineVirtualMachinesFind(filter *RequestFilter) ([]model.VirtualMachineObject, error) {
	var data []interface{}
	_, err := c.Iterate(&data, &model.VirtualMachineObject{}, "machine", "virtualMachinesFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("domain settings find: %s", err)
	}

	out := []model.VirtualMachineObject{}
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.VirtualMachineObject))
	}

	return out, nil
}
