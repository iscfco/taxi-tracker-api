package psql

import (
	"taxi-tracker-api/api/dbconn"
	"taxi-tracker-api/api/model"
)

type VehicleDao struct {
}

var qGetVehicleList = `
	SELECT * 
	FROM vehicle`

func (VehicleDao) GetVehicleList() (vehicles []model.Vehicle, err error) {
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return vehicles, err
	}
	defer db.Close()

	stmt, err := db.Prepare(qGetVehicleList)
	if err != nil {
		return vehicles, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return vehicles, err
	}
	defer rows.Close()

	var v model.Vehicle
	for rows.Next() {
		err = rows.Scan(&v.Id, &v.Make, &v.Model, &v.Year, &v.Latitude, &v.Longitude)
		if err != nil {
			return vehicles, err
		}
		vehicles = append(vehicles, v)
	}
	return vehicles, nil
}

var qGetVehiclePosition = `
	SELECT latitude, longitude 
	FROM vehicle
	WHERE id = $1`

func (VehicleDao) GetVehiclePosition(vehicleId *string) (vp model.VehiclePosition, err error) {
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return vp, err
	}
	defer db.Close()

	stmt, err := db.Prepare(qGetVehiclePosition)
	if err != nil {
		return vp, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(vehicleId)
	if err != nil {
		return vp, err
	}
	defer rows.Close()

	for rows.Next() {
		vp.VehicleId = *vehicleId
		err = rows.Scan(&vp.Latitude, &vp.Longitude)
		if err != nil {
			return vp, err
		}
	}

	return vp, nil
}

var qUpdatePosition = `
	UPDATE vehicle
	SET latitude = $1,
		longitude = $2
	WHERE id = $3
	RETURNING id`

func (VehicleDao) UpdatePosition(vp *model.VehiclePosition) (string, error) {
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var vehicleIdResult string
	err = db.QueryRow(
		qUpdatePosition,
		vp.Latitude,
		vp.Longitude,
		vp.VehicleId,
	).Scan(&vehicleIdResult)
	if err != nil {
		return vehicleIdResult, err
	}
	return vehicleIdResult, nil
}
