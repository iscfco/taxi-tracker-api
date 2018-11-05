package daoi

import "gbmchallenge/api/model"

type VehiclePositionHistoryDaoI interface {
	SavePositionInHistorical(vp *model.VehiclePosition) (string, error)
}
