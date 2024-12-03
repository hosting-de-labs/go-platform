package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

type Machine struct {
	c *ApiClient
}

func (m *Machine) MachineVirtualMachinesFind(filter *RequestFilter) ([]model.VirtualMachineObject, error) {
	var data []interface{}
	_, err := m.c.Iterate(&data, &model.VirtualMachineObject{}, "machine", "virtualMachinesFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("virtualMachinesFind: %s", err)
	}

	var out []model.VirtualMachineObject
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.VirtualMachineObject))
	}

	return out, nil
}
