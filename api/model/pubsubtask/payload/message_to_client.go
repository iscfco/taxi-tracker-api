package payload

type MessageToClient struct {
	Subject string      `json:"subject"`
	Content interface{} `json:"content"`
}
