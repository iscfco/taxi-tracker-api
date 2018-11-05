package service

import (
	"encoding/json"
	"gbmchallenge/api/daoimp/psql"
	"gbmchallenge/api/facadei"
	"gbmchallenge/api/facadeimp"
	"gbmchallenge/api/model"
	"github.com/gorilla/mux"
	"net/http"
)

type vehicleWS struct {
	vehicleFacade facadei.VehicleFacadeI
}

func NewVehicleWS() vehicleWS {
	vDao := psql.VehicleDao{}
	vphDao := psql.VehiclePositionHistoryDao{}

	return vehicleWS{
		vehicleFacade: facadeimp.NewVehicleFacade(vDao, vphDao),
	}
}

func (ws *vehicleWS) GetVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rJson := ws.vehicleFacade.GetVehicleList()
	payload, _ := json.Marshal(rJson)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func (ws *vehicleWS) GetVehiclePositionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	varsFromRequest := mux.Vars(r)
	vehicleId := varsFromRequest["vehicleId"]
	rJson := ws.vehicleFacade.GetVehiclePosition(&vehicleId)
	payload, _ := json.Marshal(rJson)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func (ws *vehicleWS) UpdateVehiclePositionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	var vehiclePosition model.VehiclePosition
	err := json.NewDecoder(r.Body).Decode(&vehiclePosition)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding body"))
		return
	}
	ws.vehicleFacade.UpdatePosition(&vehiclePosition)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
