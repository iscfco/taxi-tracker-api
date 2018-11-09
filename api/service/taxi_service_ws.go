package service

import (
	"encoding/json"
	"fmt"
	"taxi-tracker-api/api/daoimp/psql"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/facadeimp"
	"github.com/gorilla/context"
	"net/http"
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

	fmt.Println("CustomerId extracted:", customerId)
	res := ws.taxiServiceFacade.CreateService(&customerId)

	payload, _ := json.Marshal(res)
	w.WriteHeader(res.HttpCode)
	w.Write(payload)
}