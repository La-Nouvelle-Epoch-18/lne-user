package user

import (
	"fmt"

	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/auth"
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/store"
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/types"
)

func NewService(s store.Store, a auth.Operator) *Service {
	return &Service{
		store: s,
		auth:  a,
	}
}

type Service struct {
	secret string
	store  store.Store
	auth   auth.Operator
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string      `json:"token"`
	User  *types.User `json:"user"`
}

func (s *Service) LoginUser(req *LoginUserRequest) (*LoginUserResponse, error) {
	switch {
	case req.Email == "":
		return nil, fmt.Errorf("email is empty")
	case req.Password == "":
		return nil, fmt.Errorf("password is empty")
	default:
		//
	}

	user, token, err := s.auth.AuthWithCredentials(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &LoginUserResponse{
		Token: token,
		User:  user,
	}, nil
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type VerifyTokenResponse struct{}

func (s *Service) VerifyToken(token string) error {
	return s.auth.VerifyToken(token)
}
