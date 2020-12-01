package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hosting-de-labs/go-platform/model"
)

func (c *ApiClient) DomainSettingsFind(filter *RequestFilter) (*[]model.DomainSettingObject, error) {
	currentPage := 0
	var data []model.DomainSettingObject

	for {
		resp, err := c.runRequest("email", "domainSettingsFind", nil, 0, currentPage)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		domainSettingsResult := new(model.DomainSettingsResult)

		err = json.Unmarshal(body, domainSettingsResult)
		if err != nil {
			return nil, err
		}

		data = append(data, domainSettingsResult.Response.Data...)
		currentPage = domainSettingsResult.Response.Page

		if currentPage >= domainSettingsResult.Response.TotalPages {
			break
		}

		currentPage++
	}

	return &data, nil
}
