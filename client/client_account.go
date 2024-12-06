package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

type Account struct {
	c *ApiClient
}

type AccountRequest struct {
	Account model.AccountObject `json:"subaccount"`
}

type AccountResponse struct {
	EmptyResponse
	Response model.AccountObject
}

func (a *Account) SubaccountsFind(filter *RequestFilter) ([]model.AccountObject, error) {
	var data []interface{}
	_, err := a.c.Find(&data, &model.AccountObject{}, "account", "subaccountsFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("subaccountsFind: %s", err)
	}

	var out []model.AccountObject
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.AccountObject))
	}

	return out, nil
}

func (a *Account) SubaccountUpdate(account model.AccountObject) (*model.AccountObject, error) {
	accOut := &model.AccountObject{}
	err := a.c.Request("account", "subaccountUpdate", AccountRequest{account}, accOut, &model.AccountObject{})
	if err != nil {
		return nil, fmt.Errorf("subaccountUpdate: %s", err)
	}

	return accOut, nil
}
