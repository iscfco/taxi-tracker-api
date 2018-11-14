package payload

type Publish struct {
	TaskId  string      `json:"task_id"`
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
}
