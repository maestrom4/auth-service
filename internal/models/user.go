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
<<<<<<< HEAD
	UserRole          string `bson:"userRole"`
=======
	Role              string `bson:"userRole"`
>>>>>>> a077dc2c0f885525eab8209490b0d45ed10fa612
}
