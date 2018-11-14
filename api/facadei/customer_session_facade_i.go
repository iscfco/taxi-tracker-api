package facadei

import (
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/customer"
)

type CustomerSessionFacadeI interface {
	Authorize(user *model.User) (customer.CustomerSession)
}
