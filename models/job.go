package models

type job struct {
	Id          string `json:"id"`
	Status      uint8  `json:"status"`
	Attempts    uint   `json:"attempts"`
	MaxAttempts uint   `json:"max_attempts"`
	Error       string `json:"error"`
}
