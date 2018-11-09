package psql

import "gbmchallenge/api/dbconn"

type TaxiServiceDao struct {
}

var qCreateService = `SELECT taxi_service_insert($1)`

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
		err = rows.Scan(&resCode)
		if err != nil {
			return resCode, err
		}
	}

	return resCode, err
}
