package dto

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

type Message struct {
	Message string `json:"message"`
}
