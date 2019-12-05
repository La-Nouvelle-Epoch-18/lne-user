package auth

import (
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/store"
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/types"
)

type Operator interface {
	AuthWithCredentials(email, password string) (*types.User, string, error)
	VerifyToken(token string) error
}

type operator struct {
	store store.Store
}

func NewOperator(s store.Store) *operator {
	return &operator{
		store: s,
	}
}

func (o *operator) AuthWithCredentials(email, password string) (*types.User, string, error) {
	user, err := o.store.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}

	return user, "TEST_TOKEN", nil
}

func (o *operator) VerifyToken(token string) error {
	return nil
}
