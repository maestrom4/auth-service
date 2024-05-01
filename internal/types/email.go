package types

type EmailOpt struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Message   string `json:"message"`
	EmailFrom string `json:"emailFrom"`
	SmtpHosts string `json:"smtphosts"`
	SmtpPort  string `json:"smtpPort"`
	Body      string `json:"body"`
}
