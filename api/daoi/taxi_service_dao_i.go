package daoi

import (
	"taxi-tracker-api/api/model/taxi_service"
)

type TaxiServiceDaoI interface {
	CreateService(customerId *string) (string, error)
	GetCustomrService(customerId *string) (taxi_service.TaxiService, error)
}
