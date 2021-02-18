package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) ZonesFind(filter *RequestFilter) (*[]model.ZoneObject, error) {
	currentPage := 0
	var data []model.ZoneObject

	getPage := func(pageNum int) (pageData *model.ZonesResult, err error) {
		pageData = new(model.ZonesResult)

		resp, err := c.runRequest("email", "zonesFind", nil, 0, currentPage)
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

func (c *ApiClient) ZoneConfigsFind(filter *RequestFilter) (*[]model.ZoneConfigObject, error) {
	currentPage := 0
	var data []model.ZoneConfigObject

	getPage := func(pageNum int) (pageData *model.ZoneConfigsResult, err error) {
		pageData = new(model.ZoneConfigsResult)

		resp, err := c.runRequest("email", "zoneConfigsFind", nil, 0, currentPage)
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
