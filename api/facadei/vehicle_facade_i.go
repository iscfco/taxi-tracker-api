package facadei

import (
	"taxi-tracker-api/api/model"
)

type VehicleFacadeI interface {
	GetVehicleList() []model.Vehicle
	GetVehiclePosition(vehicleId *string) model.VehiclePosition
	UpdatePosition(vp *model.VehiclePosition)
}
