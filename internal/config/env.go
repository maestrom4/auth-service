package config

import "os"

var (
	Username    string
	Password    string
	Database    string
	MongoUrl    string
	SaltPassKey string
)

func ConfigInit() {
	Username = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	Password = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	Database = os.Getenv("MONGO_INITDB_DATABASE")
	MongoUrl = os.Getenv("MONGO_URI")
	SaltPassKey = os.Getenv("SALT_PASS_KEY")
}
