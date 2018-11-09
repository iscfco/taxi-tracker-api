package facadei

import "gbmchallenge/api/model"

type TaxiServiceFacadeI interface {
	CreateService(customerId *string) model.Result
}
