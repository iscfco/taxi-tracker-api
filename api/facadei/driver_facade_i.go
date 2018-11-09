package facadei

import "taxi-tracker-api/api/model"

type DriverFacadeI interface {
	CreateAccount(driver *model.Driver) model.Result
}
