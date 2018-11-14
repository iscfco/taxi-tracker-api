package psql

import (
	"taxi-tracker-api/api/dbconn"
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/taxi_service"
)

type TaxiServiceDao struct {
}

var qCreateService = `SELECT taxi_service_insert($1, $2, $3)`

func (TaxiServiceDao) CreateService(customerId *string, usrPosition *model.UserPosition) (int, string, string, error) {
	var vehicleId, userName string
	var resCode int
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return resCode, vehicleId, userName, err
	}
	defer db.Close()

	stmt, err := db.Prepare(qCreateService)
	if err != nil {
		return resCode, vehicleId, userName, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(customerId)
	if err != nil {
		return resCode, vehicleId, userName, err
	}

	for rows.Next() {
		err = rows.Scan(&resCode, &vehicleId, &userName)
		if err != nil {
			return resCode, vehicleId, userName, err
		}
	}

	return resCode, vehicleId, userName, err
}


var qGetCustomerService = `
	SELECT customer_id, vehicle_id, driver_id
	FROM taxi_service
	WHERE customer_id = $1`

func (TaxiServiceDao) GetCustomrService(customerId *string) (taxi_service.TaxiService, error) {
	var taxiService taxi_service.TaxiService
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