package facadeimp

import (
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/daoi"
	"taxi-tracker-api/api/errorhandler"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/model/taxi_service"
	"taxi-tracker-api/api/response/prebuilt"
)

type taxiServiceFacade struct {
	daoTaxiService daoi.TaxiServiceDaoI
}

func NewTaxiServiceFacade(dao daoi.TaxiServiceDaoI) facadei.TaxiServiceFacadeI {
	return &taxiServiceFacade{
		daoTaxiService: dao,
	}
}

func (f *taxiServiceFacade) CreateService(customerId *string) (res taxi_service.CreateServiceResp) {
	var err error
	res.VehicleId, err = f.daoTaxiService.CreateService(customerId)
	if err != nil {
		res.Result = errorhandler.HandleErr(&err)
		return
	}

	if res.VehicleId == "" {
		res.Result.ResCode, res.Result.Msg, res.Result.HttpCode = constants.ETS001_C, constants.ETS001_M, 200
		return
	}
	res.Result = prebuilt.GetSuccess()
	return
}

func (f *taxiServiceFacade) GetService(customerId *string) (taxiService taxi_service.TaxiService, err error) {
	taxiService, err = f.daoTaxiService.GetCustomrService(customerId)
	if err != nil {
		return taxiService, err
	}

	return taxiService, err
}