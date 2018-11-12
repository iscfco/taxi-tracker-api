package facadeimp

import (
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/daoi"
	"taxi-tracker-api/api/errorhandler"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/model"
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

func (f *taxiServiceFacade) CreateService(customerId *string) (res model.Result) {
	customerIdRes, err := f.daoTaxiService.CreateService(customerId)
	if err != nil {
		return errorhandler.HandleErr(&err)
	}

	if customerIdRes == "" {
		res.ResCode, res.Msg, res.HttpCode = constants.ETS001_C, constants.ETS001_M, 200
		return
	}
	return prebuilt.GetSuccess()
}

func (f *taxiServiceFacade) GetService(customerId *string) (taxiService model.TaxiService, err error) {
	taxiService, err = f.daoTaxiService.GetCustomrService(customerId)
	if err != nil {
		return taxiService, err
	}

	return taxiService, err
}