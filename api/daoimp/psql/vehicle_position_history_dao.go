package psql

import (
	"gbmchallenge/api/dbconn"
	"gbmchallenge/api/model"
)

type VehiclePositionHistoryDao struct {
}

var qSavePositionInHistorical = `
	INSERT INTO vehicle_position_history
		(vehicle_id,	latitude,	longitude	)
	VALUES
		($1, 			$2,			$3			)
	RETURNING vehicle_id`

func (VehiclePositionHistoryDao) SavePositionInHistorical(vp *model.VehiclePosition) (string, error) {
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var vehicleIdResult string
	err = db.QueryRow(
		qSavePositionInHistorical,
		vp.VehicleId,
		vp.Latitude,
		vp.Longitude,
	).Scan(&vehicleIdResult)
	if err != nil {
		return vehicleIdResult, err
	}
	return vehicleIdResult, nil
}
