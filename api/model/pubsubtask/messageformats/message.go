package messageformats

type Message struct {
	Subject string      `json:"subject"`
	Content interface{} `json:"content"`
}
