package service

import (
	"encoding/json"
	"taxi-tracker-api/api/daoimp/psql"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/facadeimp"
	"taxi-tracker-api/api/model"
	"net/http"
)

type driverSessionWS struct {
	sessionFacade facadei.DriverSessionFacadeI
}

func NewDriverSessionWS() driverSessionWS {
	dao := psql.DriverDao{}
	return driverSessionWS{
		sessionFacade: facadeimp.NewDriverSessionFacade(dao),
	}
}

func (ws *driverSessionWS) DriverLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := model.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding body"))
		return
	}

	session := ws.sessionFacade.Authorize(&user)
	payload, _ := json.Marshal(session)
	w.WriteHeader(session.Res.HttpCode)
	w.Write(payload)
	return
}

// If the accessToken & refreshToken live is long, create a logout and revoked tokens db
