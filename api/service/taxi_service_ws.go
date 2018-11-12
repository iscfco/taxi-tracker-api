package service

import (
	"encoding/json"
	"github.com/gorilla/context"
	"net/http"
	"taxi-tracker-api/api/daoimp/psql"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/facadeimp"
)

type taxiServiceWS struct {
	taxiServiceFacade facadei.TaxiServiceFacadeI
}

func NewTaxiServiceWS() taxiServiceWS{
	dao := psql.TaxiServiceDao{}
	return taxiServiceWS{
		taxiServiceFacade: facadeimp.NewTaxiServiceFacade(dao),
	}
}

func (ws *taxiServiceWS) CreateService(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	customerId := context.Get(r, "userId").(string)
	defer context.Clear(r)

	res := ws.taxiServiceFacade.CreateService(&customerId)

	payload, _ := json.Marshal(res)
	w.WriteHeader(res.HttpCode)
	w.Write(payload)
}

func (ws *taxiServiceWS) GetCustomerService(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	customerId := context.Get(r, "userId").(string)
	defer context.Clear(r)

	res, err := ws.taxiServiceFacade.GetService(&customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal error"))
	}

	payload, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}