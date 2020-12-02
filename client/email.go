package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) DomainSettingsFind(filter *RequestFilter) (*[]model.DomainSettingsObject, error) {
	currentPage := 0
	var data []model.DomainSettingsObject

	getPage := func(pageNum int) (pageData *model.DomainSettingsResult, err error) {
		pageData = new(model.DomainSettingsResult)

		resp, err := c.runRequest("email", "domainSettingsFind", nil, 0, currentPage)
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
