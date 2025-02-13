package graphql

import (
	cfg "auth-service/internal/config"
	gql "auth-service/internal/types"
	"auth-service/pkg/mongodb"
	"auth-service/utils"
	"errors"
	"fmt"
	"net/smtp"

	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

type Resolver struct {
	UserRepository *mongodb.UserRepository
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)
	if !ok {
		return nil, errors.New("id is not a valid string")
	}

	userRepository := mongodb.NewUserRepository(cfg.GetDBCollection(cfg.CollectionUser))
	user, err := userRepository.GetUserByID(p.Context, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type EmailSender interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

type SmtpClient struct{}

func (s *SmtpClient) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

func (r *Resolver) AddUserResolver(p graphql.ResolveParams) (interface{}, error) {
	email, _ := p.Args["email"].(string)
	password, _ := p.Args["password"].(string)

	existingUser, _ := r.UserRepository.GetUserByEmail(p.Context, email)
	if existingUser != nil {
		return nil, errors.New("email already taken")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	// Generate verification token
	verificationToken, err := utils.GenerateVerificationToken()
	if err != nil {
		return nil, err
	}

	user, err := r.UserRepository.AddUser(p.Context, email, hashedPassword, verificationToken)
	if err != nil {
		return nil, err
	}

	verificationLink := fmt.Sprintf("%s/verify?token=%s", cfg.BaseUrl, verificationToken)
	bodyAndLink := fmt.Sprintf(`<!DOCTYPE html>
	<html>
	<head>
		<title>Email Verification</title>
	</head>
	<body>
		<h1>Verify Your Email Address</h1>
		<p>Hello,</p>
		<p>Thank you for registering with us. To complete your registration, please verify your email address by clicking the link below:</p>
		<a href="%s">Verify Email</a>
		<p>If you did not request this verification, please ignore this email.</p>
		<p>Thank you!</p>
		<p>The YourService Team</p>
	</body>
	</html>`, verificationLink)

	emailData := gql.EmailOpt{
		Email:     user.Email,
		Password:  cfg.EmailPass,
		Token:     verificationToken,
		Message:   cfg.Message,
		EmailFrom: cfg.EmailFrom,
		SmtpHosts: cfg.SmtpHosts,
		SmtpPort:  cfg.SmtpPort,
		Body:      bodyAndLink,
	}
	// Send verification email
	smtpClient := &SmtpClient{}
	err = utils.SendVerificationEmail(emailData, smtpClient)
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("failed to obtain user ID after creation")
	}
	user.IsVerified = false

	return user, nil
}

func (r *Resolver) UpdateUserResolver(p graphql.ResolveParams) (interface{}, error) {
	id, idOk := p.Args["_id"].(string)
	if !idOk {
		return nil, errors.New("invalid or missing '_id' argument")
	}

	name, nameOk := p.Args["name"].(string)
	if !nameOk {
		name = ""
	}

	email, emailOk := p.Args["email"].(string)
	if !emailOk {
		email = ""
	}

	updatedUser, err := r.UserRepository.UpdateUser(p.Context, id, name, email)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}
	return updatedUser, nil
}

// DeleteUserResolver deletes a user from the system
func (r *Resolver) DeleteUserResolver(p graphql.ResolveParams) (interface{}, error) {
	id, idOk := p.Args["_id"].(string)
	if !idOk {
		return nil, errors.New("invalid or missing '_id' argument")
	}

	err := r.UserRepository.DeleteUser(p.Context, id)
	if err != nil {
		return nil, fmt.Errorf("error deleting user: %v", err)
	}
	return true, nil
}

func (r *Resolver) LoginResolver(p graphql.ResolveParams) (interface{}, error) {
	email, _ := p.Args["email"].(string)
	password, _ := p.Args["password"].(string)

	userID, userIDOk := p.Context.Value(string(cfg.UserIDKey)).(string)
	log.Debugf("Retrieved userID in resolver: %s, Exists: %t", userID, userIDOk)

	if userIDOk && userID != "" {
		return gql.LoginResponse{
			Message:    "Already authenticated via token",
			UserId:     userID,
			IsLoggedIn: true,
		}, nil
	}

	userRepository := mongodb.NewUserRepository(cfg.GetDBCollection(cfg.CollectionUser))
	user, err := userRepository.GetUserByEmail(p.Context, email)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}
	if user == nil || user.ID == "" {
		return nil, errors.New("username or password is incorrect or user ID is missing")
	}
	log.Debugln("user: ", user)
	log.Debugln("userid: ", user.ID)
	isPasswordCorrect := utils.CheckPasswordHash(password, user.HashedPassword)
	if !isPasswordCorrect {
		return nil, errors.New("email or password is incorrect")
	}

	token, err := utils.CreateToken(user.ID, cfg.JwtSecretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	return gql.LoginResponse{
		Token:      token,
		Email:      user.Email,
		UserId:     user.ID,
		Message:    "Successfully logged in",
		IsLoggedIn: true,
	}, nil
}
