package facadei

import "gbmchallenge/api/model"

type CustomerFacadeI interface {
	CreateAccount(customer *model.Customer) model.Result
}
