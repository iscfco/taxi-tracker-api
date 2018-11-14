package daoi

import (
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/driver"
)

type DriverDaoI interface {
	CreateAccount(c *driver.Driver) (res model.Result, err error)
	GetByEmail(email *string) (error, driver.Driver)
}
