package user

import (
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
	store store.Store
	auth  auth.Operator
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type LoginUserResponse struct {
	Token string
	User  *types.User
}

func (s *Service) LoginUser(req *LoginUserRequest) (*LoginUserResponse, error) {
	user, token, err := s.auth.AuthWithCredentials(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &LoginUserResponse{
		Token: token,
		User:  user,
	}, nil
}

type ValidateTokenRequest struct {
	Token string
}

type ValidateTokenResponse struct{}

func (s *Service) ValidateToken(req *ValidateTokenRequest) error {
	return s.auth.VerifyToken(req.Token)
}
