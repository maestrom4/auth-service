package graphql

import (
	cfg "auth-service/internal/config"
	"auth-service/internal/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var RegistrationMutx = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// username, _ := p.Args["username"].(string)
		// email, _ := p.Args["email"].(string)
		// password, _ := p.Args["password"].(string)

		// hashedPassword, err := hashPassword(password)
		// if err != nil {
		// 	return nil, err
		// }

		resolver, ok := p.Context.Value("resolver").(*Resolver)
		if !ok {
			return nil, errors.New("could not get the resolver from the context")
		}

		// now := time.Now().Format(time.RFC3339)
		user, err := resolver.AddUserResolver(p)
		if err != nil {
			return nil, err
		}

		userObj, ok := user.(*models.User)
		if !ok || userObj.ID == "" {
			return nil, errors.New("user ID is unavailable")
		}

		token, err := createToken(userObj.ID)
		if err != nil {
			return nil, errors.New("could not create token")
		}
		return token, nil
	},
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func createToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(cfg.SaltPassKey))
	return tokenString, err
}
