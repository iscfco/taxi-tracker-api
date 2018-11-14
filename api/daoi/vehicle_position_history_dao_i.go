package daoi

import (
	"taxi-tracker-api/api/model/vehicle"
)

type VehiclePositionHistoryDaoI interface {
	SavePositionInHistorical(vp *vehicle.VehiclePosition) (string, error)
}
