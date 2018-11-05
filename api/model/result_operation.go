package model

type Result struct {
	ResultCode     string `json:"result_code"`
	Message        string `json:"message"`
	HttpStatusCode int    `json:"-"`
}
