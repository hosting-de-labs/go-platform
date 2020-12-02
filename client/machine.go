package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) MachineVirtualMachinesFind(filter *RequestFilter) (*[]model.VirtualMachineObject, error) {
	currentPage := 0
	var data []model.VirtualMachineObject

	for {
		resp, err := c.runRequest("machine", "virtualMachinesFind", nil, 0, currentPage)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		virtualMachineResult := new(model.VirtualMachineResult)

		err = json.Unmarshal(body, virtualMachineResult)
		if err != nil {
			return nil, err
		}

		data = append(data, virtualMachineResult.Response.Data...)
		currentPage = virtualMachineResult.Response.Page

		if currentPage >= virtualMachineResult.Response.TotalPages {
			break
		}

		currentPage++
	}

	return &data, nil
}
