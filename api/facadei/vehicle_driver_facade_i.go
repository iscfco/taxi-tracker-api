package facadei

import "taxi-tracker-api/api/model"

type VehicleDriverFacadeI interface {
	Create(vehicleId, driverId *string) model.Result
}
