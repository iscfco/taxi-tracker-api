package service

import (
	"encoding/json"
	"net/http"
	"taxi-tracker-api/api/daoimp/psql"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/facadeimp"
	"taxi-tracker-api/api/model/driver"
)

type driverWS struct {
	driverFacade facadei.DriverFacadeI
}

func NewDriverWS() driverWS {
	dao := psql.DriverDao{}
	return driverWS{
		driverFacade: facadeimp.NewDriverFacade(dao),
	}
}

func (ws *driverWS) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	driver := driver.Driver{}
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding body"))
		return
	}
	res := ws.driverFacade.CreateAccount(&driver)
	payload, _ := json.Marshal(res)
	w.WriteHeader(res.HttpCode)
	w.Write(payload)
}
