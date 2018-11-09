package facadei

import "gbmchallenge/api/model"

type VehicleDriverFacadeI interface {
	Create(vehicleId, driverId *string) model.Result
}
