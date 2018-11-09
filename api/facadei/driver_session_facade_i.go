package facadei

import "taxi-tracker-api/api/model"

type DriverSessionFacadeI interface {
	Authorize(user *model.User) model.DriverSession
}
