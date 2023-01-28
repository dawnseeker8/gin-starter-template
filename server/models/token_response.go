package models

type TokenResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	User    any    `json:"user"`
}
