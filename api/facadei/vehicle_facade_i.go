package facadei

import (
	"taxi-tracker-api/api/model/vehicle"
)

type VehicleFacadeI interface {
	GetVehicleList() []vehicle.Vehicle
	GetVehiclePosition(vehicleId *string) vehicle.VehiclePosition
	UpdatePosition(vp *vehicle.VehiclePosition)
	UpdatePositionV2(vp *vehicle.VehiclePosition)
}
