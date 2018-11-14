package daoi

import (
	"taxi-tracker-api/api/model/vehicle"
)

type VehicleDaoI interface {
	GetVehicleList() (vehicles []vehicle.Vehicle, err error)
	GetVehiclePosition(vehicleId *string) (vp vehicle.VehiclePosition, err error)
	UpdatePosition(vp *vehicle.VehiclePosition) (string, error)
}
