package psql

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/dbconn"
	"gbmchallenge/api/model"
)

type CustomerDao struct {
}

var qCustomerInsert = `SELECT * FROM customer_insert($1, $2, $3, $4, $5)`

func (CustomerDao) CreateAccount(c *model.Customer) (res model.Result, err error) {
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return res, err
	}
	defer db.Close()

	var resCode int8
	err = db.QueryRow(qCustomerInsert, c.Id, c.FirstName, c.LastName, c.Email, c.Password).Scan(&resCode)
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

var qGetByEmail = `
	SELECT C.id, C.first_name, C.last_name, C.email, C.password 
	FROM customer C 
	WHERE email = $1`

func (CustomerDao) GetByEmail(email *string) (err error, c model.Customer){
	db, err := dbconn.GetPsqlDBConn()
	if err != nil {
		return err, c
	}
	defer db.Close()

	stmt, err := db.Prepare(qGetByEmail)
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