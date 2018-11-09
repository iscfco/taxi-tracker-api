package daoi

import "taxi-tracker-api/api/model"

type CustomerDaoI interface {
	CreateAccount(c *model.Customer) (res model.Result, err error)
	GetByEmail(email *string) (error, model.Customer)
}
