package daoi

import "taxi-tracker-api/api/model"

type TaxiServiceDaoI interface {
	CreateService(customerId *string) (string, error)
	GetCustomrService(customerId *string) (model.TaxiService, error)
}
