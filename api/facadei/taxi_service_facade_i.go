package facadei

import (
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/taxi_service"
)

type TaxiServiceFacadeI interface {
	CreateService(customerId *string, userPosition *model.UserPosition) taxi_service.CreateServiceResp
	GetService(customerId *string) (taxiService taxi_service.TaxiService, err error)
}
