package config

type EventListenerSettings struct {
	WorkerNumber int `mapstructure:"worker_number"`
	TaskLimit int `mapstructure:"task_limit"`
}
