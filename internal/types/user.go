package types

type RegistrationResponse struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}