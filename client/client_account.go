package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

type Account struct {
	c *ApiClient
}

type AccountRequest struct {
	Account model.AccountObject `json:"account"`
}

type AccountResponse struct {
	EmptyResponse
	Response model.AccountObject
}

func (a *Account) AccountsFind(filter *RequestFilter) ([]model.AccountObject, error) {
	var data []interface{}
	_, err := a.c.Find(&data, &model.AccountObject{}, "account", "accountsFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("accountsFind: %s", err)
	}

	var out []model.AccountObject
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.AccountObject))
	}

	return out, nil
}

func (a *Account) AccountUpdate(account model.AccountObject) (*model.AccountObject, error) {
	accOut := &model.AccountObject{}
	err := a.c.ParsedRequest("account", "accountUpdate", AccountRequest{account}, accOut, &model.AccountObject{})
	if err != nil {
		return nil, fmt.Errorf("accountUpdate: %s", err)
	}

	return accOut, nil
}
