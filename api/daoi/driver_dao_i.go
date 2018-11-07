package daoi

import "gbmchallenge/api/model"

type DriverDaoI interface {
	CreateAccount(c *model.Driver) (res model.Result, err error)
}
