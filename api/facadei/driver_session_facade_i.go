package facadei

import (
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/driver"
)

type DriverSessionFacadeI interface {
	Authorize(user *model.User) driver.DriverSession
}
