package models

type User struct {
	ID             string `bson:"_id,omitempty"`
	Username       string `bson:"username"`
	HashedPassword string `bson:"hashed_password"`
	CreatedAt      string
	LastLogin      string
	// FirstName      string `bson:"firstname"`
	// LastName       string `bson:"lastname"`
	// CreateAt       string `bson:"createAt"`
	// LastLogin      string `bson:"lastLogin"`
}
