package facadeimp

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/daoi"
	"gbmchallenge/api/errorhandler"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/model"
	"gbmchallenge/api/response/prebuilt"
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
