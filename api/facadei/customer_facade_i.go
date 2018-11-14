package facadei

import (
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/customer"
)

type CustomerFacadeI interface {
	CreateAccount(customer *customer.Customer) model.Result
}
