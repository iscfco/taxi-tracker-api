package util

import (
	"gbmchallenge/api/constants"
	"gbmchallenge/api/model"
)

/*
func PrepareResponseWriter(w http.ResponseWriter, response response.RespI) http.ResponseWriter {
	statusCode := GetHttpStatusCode(response.GetCode())
	payload := []byte(response.GetPayload())
	w.WriteHeader(statusCode)
	w.Write(payload)
	printResp(&payload, &statusCode)
	return w
}*/

func GetServerErr() model.Result {
	return model.Result{
		ResultCode:     constants.ESR001_C,
		Message:        constants.ESR001_M,
		HttpStatusCode: 500,
	}
}
