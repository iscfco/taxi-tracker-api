package daoi

import "taxi-tracker-api/api/model"

type VehiclePositionHistoryDaoI interface {
	SavePositionInHistorical(vp *model.VehiclePosition) (string, error)
}
