package prebuilt

import (
	"taxi-tracker-api/api/constants"
	"taxi-tracker-api/api/model"
)

func GetSuccess() model.Result {
	return model.Result{
		ResCode:  constants.SUCCESS_C,
		Msg:      constants.SUCCESS_M,
		HttpCode: 200,
	}
}

func GetInternalErr() model.Result {
	return model.Result{
		ResCode:  constants.ESR001_C,
		Msg:      constants.ESR001_M,
		HttpCode: 500,
	}
}
