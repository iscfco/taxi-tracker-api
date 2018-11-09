package psql

import (
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/dbconn"
	"taxi-tracker-api/api/model"
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
		res.ResCode, res.Msg, res.HttpCode = constants.SUCCESS_C, constants.SUCCESS_M, 200
	case -1:
		res.ResCode, res.Msg, res.HttpCode = constants.EUS001_C, constants.EUS001_M, 200
	}

	return res, nil
}

var qDriverGetByEmail = `
	SELECT C.id, C.first_name, C.last_name, C.email, C.password 
	FROM driver C 
	WHERE email = $1`

func (DriverDao) GetByEmail(email *string) (err error, c model.Driver) {
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return err, c
	}
	defer db.Close()

	stmt, err := db.Prepare(qDriverGetByEmail)
	if err != nil {
		return err, c
	}
	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		return err, c
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&c.Id, &c.FirstName, &c.LastName, &c.Email, &c.Password)
		if err != nil {
			return err, c
		}
	}
	return err, c
}
