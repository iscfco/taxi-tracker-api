package psql

import "gbmchallenge/api/dbconn"

type TaxiServiceDao struct {
}

var qCreateService = `
	INSERT INTO taxi_service (customer_id, vehicle_id, driver_id)
		SELECT $1, vehicle_id, driver_id FROM vehicle_driver LIMIT 1
	RETURNING vehicle_id`

func (TaxiServiceDao) CreateService(customerId *string) (string, error) {
	var resCode string
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return resCode, err
	}
	defer db.Close()

	stmt, err := db.Prepare(qCreateService)
	if err != nil {
		return resCode, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(customerId)
	if err != nil {
		return resCode, err
	}

	for rows.Next() {
		err = rows.Scan(resCode)
		if err != nil {
			return resCode, err
		}
	}

	return resCode, err
}
