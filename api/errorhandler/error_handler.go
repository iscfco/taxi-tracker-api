package errorhandler

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/model"
	"log"
)

func HandleErr(err *error) model.Result {
	log.Println(*err)
	return model.Result{
		ResCode:  constants.ESR001_C,
		Msg:      constants.ESR001_M,
		HttpCode: 500,
	}
}
