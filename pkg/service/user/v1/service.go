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
	store store.Store
	auth  auth.Operator
}

type CreateUserRequest struct {
	Email    string
	Username string
	Password string
	IsSchool bool
}

func (s *Service) CreateUser(req *CreateUserRequest) error {
	switch {
	case req.Email == "":
		return fmt.Errorf("email is empty")
	case req.Username == "":
		return fmt.Errorf("username is empty")
	case req.Password == "":
		return fmt.Errorf("password is empty")
	default:
		//
	}

	kind := store.USER_SCHOOL
	if req.IsSchool {
		kind = store.USER_STUDENT
	}
	return s.store.CreateUser(req.Email, req.Username, req.Password, kind)
}

type GetUsersRequest struct {
	IDs []string
}

type GetUsersResponse struct {
	Usernames []*types.User
}

func (s *Service) GetUsers(req *GetUsersRequest) {

}

type GetUsernamesRequest struct {
	IDs []string
}

type GetUsernamesResponse struct {
	Usernames []string
}

func (s *Service) GetUsernames(req *GetUsernamesRequest) {

}
