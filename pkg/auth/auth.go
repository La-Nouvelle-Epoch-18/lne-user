package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/crypto"
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/store"
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/types"
)

type Operator interface {
	AuthWithCredentials(email, password string) (*types.User, string, error)
	VerifyToken(token string) error
}

type operator struct {
	store  store.Store
	secret []byte
}

func NewOperator(secret string, s store.Store) *operator {
	return &operator{
		store:  s,
		secret: []byte(secret),
	}
}

func (o *operator) AuthWithCredentials(email, password string) (*types.User, string, error) {
	user, err := o.store.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if user.Password != crypto.Sha256(password) {
		return nil, "", fmt.Errorf("incorrect password")
	}

	user.Password = ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":       user.ID,
		"user_email":    user.Email,
		"user_username": user.Username,
		"user_type":     user.Type,
		"nbf":           time.Now().Unix(),
	})

	tokenStr, err := token.SignedString(o.secret)
	if err != nil {
		return nil, "", err
	}

	return user, tokenStr, nil
}

func (o *operator) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return o.secret, nil
	})
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	}
	return fmt.Errorf("invalid token: %v", err)
}
