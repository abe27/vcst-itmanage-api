package models

type Response struct {
	Success bool        `json:"success,omitempty"`
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
