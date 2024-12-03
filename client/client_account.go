package client

import (
	"fmt"

	"github.com/hosting-de-labs/go-platform/model"
)

type Account struct {
	c *ApiClient
}

func (a *Account) AccountsFind(filter *RequestFilter) ([]model.AccountObject, error) {
	var data []interface{}
	_, err := a.c.Iterate(&data, &model.AccountObject{}, "account", "accountsFind", filter, 0)
	if err != nil {
		return nil, fmt.Errorf("accountsFind: %s", err)
	}

	var out []model.AccountObject
	for i := 0; i < len(data); i++ {
		out = append(out, *data[i].(*model.AccountObject))
	}

	return out, nil
}

func (a *Account) AccountUpdate(account *model.AccountObject) (model.AccountObject, error) {
	var data []interface{}
	_, err := a.c.Update("account", "accountUpdate", account)
	if err != nil {
		return *account, fmt.Errorf("accountUpdate: %s", err)
	}

	return out, nil
}
