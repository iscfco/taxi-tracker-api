package service

import (
	"encoding/json"
	"net/http"
	"taxi-tracker-api/api/daoimp/psql"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/facadeimp"
	"taxi-tracker-api/api/model/customer"
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
	customer := customer.Customer{}

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
