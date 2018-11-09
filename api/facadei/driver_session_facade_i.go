package facadei

import "gbmchallenge/api/model"

type DriverSessionFacadeI interface {
	Authorize(user *model.User) model.DriverSession
}
