package psql

import (
	"taxi-tracker-api/api/dbconn"
	"taxi-tracker-api/api/model/vehicle"
)

type VehiclePositionHistoryDao struct {
}

var qSavePositionInHistorical = `
	INSERT INTO vehicle_position_history
		(vehicle_id,	latitude,	longitude	)
	VALUES
		($1, 			$2,			$3			)
	RETURNING vehicle_id`

func (VehiclePositionHistoryDao) SavePositionInHistorical(vp *vehicle.VehiclePosition) (string, error) {
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
