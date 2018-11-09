package facadei

import "taxi-tracker-api/api/model"

type CustomerFacadeI interface {
	CreateAccount(customer *model.Customer) model.Result
}
