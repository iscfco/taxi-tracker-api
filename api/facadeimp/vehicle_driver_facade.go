package facadeimp

import (
	"taxi-tracker-api/api/daoi"
	"taxi-tracker-api/api/errorhandler"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/response/prebuilt"
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
