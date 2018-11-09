package facadeimp

import (
	"gbmchallenge/api/daoi"
	"gbmchallenge/api/errorhandler"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/model"
	"gbmchallenge/api/response/prebuilt"
)

type vehicleDriverFacade struct {
	vehicleDriverFacadeDao daoi.VehicleDriverDaoI
}

func NewVehicleDriverFacade(dao daoi.VehicleDriverDaoI) facadei.VehicleDriverFacadeI {
	return &vehicleDriverFacade{
		vehicleDriverFacadeDao: dao,
	}
}

func (f *vehicleDriverFacade) Create(vehicleId, driverId *string) model.Result {
	err := f.vehicleDriverFacadeDao.Create(vehicleId, driverId)
	if err != nil {
		return errorhandler.HandleErr(&err)
	}
	return prebuilt.GetSuccess()
}
