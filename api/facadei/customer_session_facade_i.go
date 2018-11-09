package facadei

import "taxi-tracker-api/api/model"

type CustomerSessionFacadeI interface {
	Authorize(user *model.User) (model.CustomerSession)
}
