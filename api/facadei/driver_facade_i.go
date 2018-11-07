package facadei

import "gbmchallenge/api/model"

type DriverFacadeI interface {
	CreateAccount(driver *model.Driver) model.Result
}
