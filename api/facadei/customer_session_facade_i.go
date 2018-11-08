package facadei

import "gbmchallenge/api/model"

type CustomerSessionFacadeI interface {
	Authorize(user *model.User) (model.Session)
}
