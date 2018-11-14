package daoi

import (
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/customer"
)

type CustomerDaoI interface {
	CreateAccount(c *customer.Customer) (res model.Result, err error)
	GetByEmail(email *string) (error, customer.Customer)
}
