package facadei

import (
	"gbmchallenge/api/model"
)

type VehicleFacadeI interface {
	GetVehicleList() []model.Vehicle
	GetVehiclePosition(vehicleId *string) model.VehiclePosition
	UpdatePosition(vp *model.VehiclePosition)
}
