package psql

import (
	"taxi-tracker-api/api/dbconn"
	"taxi-tracker-api/api/model"
)

type TaxiServiceDao struct {
}

var qCreateService = `SELECT taxi_service_insert($1)`

func (TaxiServiceDao) CreateService(customerId *string) (string, error) {
	var vehicleId string
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return vehicleId, err
	}
	defer db.Close()

	stmt, err := db.Prepare(qCreateService)
	if err != nil {
		return vehicleId, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(customerId)
	if err != nil {
		return vehicleId, err
	}

	for rows.Next() {
		err = rows.Scan(&vehicleId)
		if err != nil {
			return vehicleId, err
		}
	}

	return vehicleId, err
}


var qGetCustomerService = `
	SELECT customer_id, vehicle_id, driver_id
	FROM taxi_service
	WHERE customer_id = $1`

func (TaxiServiceDao) GetCustomrService(customerId *string) (model.TaxiService, error) {
	var taxiService model.TaxiService
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return taxiService, err
	}
	defer db.Close()

	stmt, err := db.Prepare(qGetCustomerService)
	if err != nil {
		return taxiService, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(customerId)
	if err != nil {
		return taxiService, err
	}

	for rows.Next() {
		err = rows.Scan(&taxiService.CustomerId, &taxiService.VehicleId, &taxiService.DriverId)
		if err != nil {
			return taxiService, err
		}
	}

	return taxiService, err
}