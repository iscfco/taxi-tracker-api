package service

import (
	"encoding/json"
	"gbmchallenge/api/daoimp/psql"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/facadeimp"
	"gbmchallenge/api/model"
	"net/http"
)

type customerSessionWS struct {
	sessionFacade facadei.CustomerSessionFacadeI
}

func NewCustomerSessionWS() customerSessionWS {
	dao := psql.CustomerDao{}
	return customerSessionWS{
		sessionFacade: facadeimp.NewCustomerSessionFacade(dao),
	}
}

func (ws *customerSessionWS) CustomerLogin(w http.ResponseWriter, r *http.Request) {
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
