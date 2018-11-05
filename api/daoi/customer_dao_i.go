package daoi

import "gbmchallenge/api/model"

type CustomerDaoI interface {
	CreateAccount(c *model.Customer) (res model.Result, err error)
}
