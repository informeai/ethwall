package dto

//Status return errors with message
type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
