package PubSubTask

type Task struct {
	TaskType string      `json:"task_type"`
	Payload  interface{} `json:"payload"`
}
