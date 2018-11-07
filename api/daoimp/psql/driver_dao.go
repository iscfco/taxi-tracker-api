package psql

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/dbconn"
	"gbmchallenge/api/model"
)

type DriverDao struct {
}

var qDriverInsert = `SELECT * FROM driver_insert($1, $2, $3, $4, $5)`

func (DriverDao) CreateAccount(d *model.Driver) (res model.Result, err error) {
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return res, err
	}
	defer db.Close()

	var resCode int8
	err = db.QueryRow(qDriverInsert, d.Id, d.FirstName, d.LastName, d.Email, d.Password).Scan(&resCode)
	if err != nil {
		return res, err
	}

	switch resCode {
	case 0:
		res.ResultCode, res.Message, res.HttpStatusCode = constants.SUCCESS_C, constants.SUCCESS_M, 200
	case -1:
		res.ResultCode, res.Message, res.HttpStatusCode = constants.EUS001_C, constants.EUS001_M, 200
	}

	return res, nil
}
