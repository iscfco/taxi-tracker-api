package daoi

import "taxi-tracker-api/api/model"

type DriverDaoI interface {
	CreateAccount(c *model.Driver) (res model.Result, err error)
	GetByEmail(email *string) (error, model.Driver)
}
