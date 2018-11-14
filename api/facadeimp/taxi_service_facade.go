package facadeimp

import (
	"encoding/json"
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/daoi"
	"taxi-tracker-api/api/errorhandler"
	"taxi-tracker-api/api/facadei"
	"taxi-tracker-api/api/helper"
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/model/pubsubtask/brokertasks"
	"taxi-tracker-api/api/model/pubsubtask/brokertasks/payload"
	"taxi-tracker-api/api/model/pubsubtask/messageformats"
	"taxi-tracker-api/api/model/pubsubtask/messageformats/content"
	"taxi-tracker-api/api/model/taxi_service"
	"taxi-tracker-api/api/response/prebuilt"
)

type taxiServiceFacade struct {
	daoTaxiService daoi.TaxiServiceDaoI
}

func NewTaxiServiceFacade(dao daoi.TaxiServiceDaoI) facadei.TaxiServiceFacadeI {
	return &taxiServiceFacade{
		daoTaxiService: dao,
	}
}

func (f *taxiServiceFacade) CreateService(customerId *string, userPosition *model.UserPosition) (
	res taxi_service.CreateServiceResp) {

	resCode, vehicleId, userName, err := f.daoTaxiService.CreateService(customerId, userPosition)
	if err != nil {
		res.Result = errorhandler.HandleErr(&err)
		return
	}

	if res.VehicleId = vehicleId; resCode == 1 {
		res.Result = prebuilt.GetSuccess()
		return
	}
	if res.VehicleId == "" {
		res.Result.ResCode, res.Result.Msg, res.Result.HttpCode = constants.ETS001_C, constants.ETS001_M, 200
		return
	}

	task := brokertasks.Task{
		TaskType: constants.Publish,
		Payload: payload.Publish{
			Topic: vehicleId,
			Message: messageformats.Message{
				Subject: constants.TaxiServiceRequest,
				Content: content.TaxiServiceRequest{
					UserName: userName,
					UserLatitude: userPosition.UserLatitude,
					UserLongitude: userPosition.UserLongitude,
				},
			},
		},
	}

	taskInBytes, _ := json.Marshal(task)
	taskInStr := string(taskInBytes)
	err = helper.WebSocketPublisher{}.Publish(&vehicleId, &taskInStr)
	if err != nil {
		errorhandler.HandleErr(&err)
	}
	res.Result = prebuilt.GetSuccess()
	return
}

func (f *taxiServiceFacade) GetService(customerId *string) (taxiService taxi_service.TaxiService, err error) {
	taxiService, err = f.daoTaxiService.GetCustomrService(customerId)
	if err != nil {
		return taxiService, err
	}

	return taxiService, err
}
