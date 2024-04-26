package types

type RegistrationResponse struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
}
