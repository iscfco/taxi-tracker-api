package facadeimp

import (
	"encoding/json"
	"fmt"
	"log"
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/daoi"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/helper"
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/PubSubTask"
	"taxi-tracker-api/api/model/PubSubTask/payload"
)

type vehicleFacadeImp struct {
	vehicleDao                daoi.VehicleDaoI
	vehiclePositionHistoryDao daoi.VehiclePositionHistoryDaoI
}

func NewVehicleFacade(v daoi.VehicleDaoI, vph daoi.VehiclePositionHistoryDaoI) facadei.VehicleFacadeI {
	return &vehicleFacadeImp{
		vehicleDao:                v,
		vehiclePositionHistoryDao: vph,
	}
}

func (f *vehicleFacadeImp) GetVehicleList() []model.Vehicle {
	v, err := f.vehicleDao.GetVehicleList()
	if err != nil {
		// handle
		fmt.Println(err)
	}
	return v
}

func (f *vehicleFacadeImp) GetVehiclePosition(vehicleId *string) model.VehiclePosition {
	v, err := f.vehicleDao.GetVehiclePosition(vehicleId)
	if err != nil {
		// handle
		fmt.Println(err)
	}
	return v
}

func (f *vehicleFacadeImp) UpdatePosition(vp *model.VehiclePosition) {
	vehicleId, err := f.vehicleDao.UpdatePosition(vp)
	if err != nil {
		// handle
		fmt.Println(err)
	}
	if vehicleId == "" {
		return
	}

	vehicleId, err = f.vehiclePositionHistoryDao.SavePositionInHistorical(vp)
	if err != nil {
		// handle
		fmt.Println(err)
	}
	if vehicleId == "" {
		return
	}
}

func (f *vehicleFacadeImp) UpdatePositionV2(vp *model.VehiclePosition) {
	vehicleId, err := f.vehicleDao.UpdatePosition(vp)
	if err != nil {
		// handle
		fmt.Println(err)
	}
	if vehicleId == "" {
		return
	}

	vehicleId, err = f.vehiclePositionHistoryDao.SavePositionInHistorical(vp)
	if err != nil {
		// handle
		fmt.Println(err)
	}
	if vehicleId == "" {
		return
	}

	task := PubSubTask.Task{
		TaskType: constants.Publish,
		Payload: payload.Publish{
			Topic: vp.VehicleId,
			Message: payload.MessageToClient{
				Subject: constants.VehiclePositionUpdate,
				Content: *vp,
			},
		},
	}

	taskInBytes, _ := json.Marshal(task)
	taskInStr := string(taskInBytes)
	err = helper.WebSocketPublisher{}.Publish(&vp.VehicleId, &taskInStr)
	if err != nil {
		log.Println(err)
	}
}
