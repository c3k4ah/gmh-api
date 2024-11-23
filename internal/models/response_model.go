package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type AppartementResponse struct {
	Appartement Appartement `json:"appartement"`
}

type AppartementsResponse struct {
	Appartements []Appartement `json:"appartements"`
}

type SuccessMessage struct {
	Message string `json:"message"`
}
