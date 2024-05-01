package models

type User struct {
	ID             string `bson:"_id,omitempty"`
	Username       string `bson:"username"`
	Email          string `bson: "email"`
	HashedPassword string `bson:"hashed_password"`
	IsVerified     bool   `bson:"isVerified"`
	CreatedAt      string `bson:"createAt"`
	LastLogin      string `bson:"lastLogin"`
	// FirstName      string `bson:"firstname"`
	// LastName       string `bson:"lastname"`
	// CreateAt       string `bson:"createAt"`
	// LastLogin      string `bson:"lastLogin"`
}
