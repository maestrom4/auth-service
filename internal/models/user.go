package models

type User struct {
	ID                string `bson:"_id,omitempty"`
	Username          string `bson:"username"`
	Email             string `bson: "email"`
	HashedPassword    string `bson:"hashed_password"`
	IsVerified        bool   `bson:"isVerified"`
	VerificationToken string `bson:"verificationToken"`
	CreatedAt         string `bson:"createAt"`
	LastLogin         string `bson:"lastLogin"`
	UserRole          string `bson:"userRole"`
}
