package facadei

import "taxi-tracker-api/api/model"

type TaxiServiceFacadeI interface {
	CreateService(customerId *string) model.Result
	GetService(customerId *string) (taxiService model.TaxiService, err error)
}
