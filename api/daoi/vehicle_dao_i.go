package daoi

import (
	"taxi-tracker-api/api/model"
)

type VehicleDaoI interface {
	GetVehicleList() (vehicles []model.Vehicle, err error)
	GetVehiclePosition(vehicleId *string) (vp model.VehiclePosition, err error)
	UpdatePosition(vp *model.VehiclePosition) (string, error)
}
