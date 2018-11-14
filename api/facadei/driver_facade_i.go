package facadei

import (
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/driver"
)

type DriverFacadeI interface {
	CreateAccount(driver *driver.Driver) model.Result
}
