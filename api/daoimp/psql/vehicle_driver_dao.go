package psql

import "gbmchallenge/api/dbconn"

type VehicleDriverDao struct {
}

var qCreateVehicleDriver = `SELECT vehicle_driver_insert($1, $2)`

func (VehicleDriverDao)Create(vehicleId, driverId *string) (err error){
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(qCreateVehicleDriver, vehicleId, driverId)
	if err != nil {
		return err
	}

	return nil
}
