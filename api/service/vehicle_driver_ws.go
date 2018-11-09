package service

import (
	"encoding/json"
	"fmt"
	"taxi-tracker-api/api/daoimp/psql"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/facadeimp"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

type vehicleDriverWS struct {
	vehicleDriverFacade facadei.VehicleDriverFacadeI
}

func NewVehicleDriverWS() vehicleDriverWS {
	dao := psql.VehicleDriverDao{}
	return vehicleDriverWS{
		vehicleDriverFacade: facadeimp.NewVehicleDriverFacade(dao),
	}
}

func (ws *vehicleDriverWS) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	driverId := context.Get(r, "userId").(string)
	defer context.Clear(r)

	vehicleId := mux.Vars(r)["vehicleId"]
	fmt.Println("CustomerId extracted:", driverId)

	res := ws.vehicleDriverFacade.Create(&vehicleId, &driverId, )

	payload, _ := json.Marshal(res)
	w.WriteHeader(res.HttpCode)
	w.Write(payload)
}
