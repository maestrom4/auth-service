package types

type RegistrationResponse struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}

type LoginResponse struct {
	Token      string `json:"token"`
	Username   string `json:"username"`
	Message    string `json:"message"`
	IsLoggedIn bool   `json: "isLoggedIn"`
	UserId     string `json: "userId"`
}
