package facadeimp

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/daoi"
	"gbmchallenge/api/errorhandler"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/model"
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
	vehicleId, err := f.daoTaxiService.CreateService(customerId)
	if err != nil {
		return errorhandler.HandleErr(&err)
	}

	res.HttpCode = 200
	if vehicleId == "" {
		res.ResCode, res.Msg = constants.ETS001_C, constants.ETS001_M
		return
	}
	res.ResCode, res.Msg = constants.SUCCESS_C, constants.SUCCESS_M
	return
}
