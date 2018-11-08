package service

import (
	"encoding/json"
	"gbmchallenge/api/daoimp/psql"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/facadeimp"
	"gbmchallenge/api/model"
	"net/http"
)

type customerWS struct {
	customerFacade facadei.CustomerFacadeI
}

func NewCustomerWS() customerWS {
	c := psql.CustomerDao{}

	return customerWS{
		customerFacade: facadeimp.NewCustomerFacade(c),
	}
}

func (ws *customerWS) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customer := model.Customer{}

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding body"))
		return
	}
	res := ws.customerFacade.CreateAccount(&customer)
	payload, _ := json.Marshal(res)
	w.WriteHeader(res.HttpCode)
	w.Write(payload)
}
