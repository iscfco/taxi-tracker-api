package prebuilt

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/model"
)

func GetSuccess() model.Result {
	return model.Result{
		ResCode:  constants.SUCCESS_C,
		Msg:      constants.SUCCESS_M,
		HttpCode: 500,
	}
}

func GetInternalErr() model.Result {
	return model.Result{
		ResCode:  constants.ESR001_C,
		Msg:      constants.ESR001_M,
		HttpCode: 500,
	}
}
