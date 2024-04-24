package graphql

import (
	cfg "auth-service/internal/config"
	"auth-service/pkg/mongodb"
	"auth-service/utils"
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
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

func (r *Resolver) AddUserResolver(p graphql.ResolveParams) (interface{}, error) {
	username, _ := p.Args["username"].(string)
	password, _ := p.Args["password"].(string)

	existingUser, _ := r.UserRepository.GetUserByUsername(p.Context, username)
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := r.UserRepository.AddUser(p.Context, username, hashedPassword)
	if err != nil {
		return nil, err
	}

	// Ensure that ID is being set properly
	if user.ID == "" {
		return nil, errors.New("failed to obtain user ID after creation")
	}

	return user, nil
}

// func (r *Resolver) AddUserResolver(p graphql.ResolveParams) (interface{}, error) {
// 	name, _ := p.Args["name"].(string)
// 	email, _ := p.Args["email"].(string)
// 	hashedPassword, _ := p.Args["hashed_password"].(string)
// 	log.Debugln("hashedPassword: ", hashedPassword)
// 	user, err := r.UserRepository.AddUser(p.Context, name, email, hashedPassword)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Debugln("user: ", user)
// 	return user, nil
// }

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
