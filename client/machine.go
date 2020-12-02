package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) MachineVirtualMachinesFind(filter *RequestFilter) (*[]model.VirtualMachineObject, error) {
	currentPage := 0
	var data []model.VirtualMachineObject

	getPage := func(pageNum int) (pageData *model.VirtualMachineResult, err error) {
		resp, err := c.runRequest("machine", "virtualMachinesFind", nil, 0, currentPage)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, pageData)
		if err != nil {
			return nil, err
		}

		return pageData, nil
	}

	for {
		pageData, err := getPage(currentPage)
		if err != nil {
			return nil, err
		}

		data = append(data, pageData.Response.Data...)
		currentPage = pageData.Response.Page

		if currentPage >= pageData.Response.TotalPages {
			break
		}

		currentPage++
	}

	return &data, nil
}
