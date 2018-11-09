package errorhandler

import (
	"taxi-tracker-api/api/model"
	"taxi-tracker-api/api/response/prebuilt"
	"log"
)

func HandleErr(err *error) model.Result {
	log.Println(*err)
	return prebuilt.GetInternalErr()
}
