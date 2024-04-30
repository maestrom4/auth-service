package config

import (
	"os"
)

var (
	Username       string
	Password       string
	Database       string
	MongoUrl       string
	SaltPassKey    string
	JwtSecretKey   string
	CollectionUser string
	MongoUsersPort string
	LogLevel       string
	ExpiryHour     string
	AppName        string
	Message        string
	EmailFrom      string
	SmtpHosts      string
	SmtpPort       string
	EmailPass      string
)

func ConfigInit() {
	Username = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	Password = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	Database = os.Getenv("MONGO_INITDB_DATABASE")
	MongoUrl = os.Getenv("MONGO_URI")
	SaltPassKey = os.Getenv("SALT_PASS_KEY")
	JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	CollectionUser = os.Getenv("MONGO_INITDB_COLLECTION")
	MongoUsersPort = os.Getenv("MONGO_USERS_PORT")
	LogLevel = os.Getenv("LOG_LEVEL")
	ExpiryHour = os.Getenv("EXPIRY_HOUR")
	AppName = os.Getenv("APP_NAME")
	Message = os.Getenv("SMTP_SUBJECT")
	EmailFrom = os.Getenv("SMTP_FROM")
	SmtpHosts = os.Getenv("SMTP_HOSTS")
	SmtpPort = os.Getenv("SMTP_PORT")
	EmailPass = os.Getenv("SMTP_PASS")
}
