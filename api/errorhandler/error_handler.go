package errorhandler

import (
	"gbmchallenge/api/model"
	"gbmchallenge/api/response/prebuilt"
	"log"
)

func HandleErr(err *error) model.Result {
	log.Println(*err)
	return prebuilt.GetInternalErr()
}
