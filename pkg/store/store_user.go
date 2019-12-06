package store

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/crypto"
	"github.com/La-Nouvelle-Epoch-18/lne-user/pkg/types"
)

var _ UserStore = &store{}

// UserStore .
type UserStore interface {
	// No validation has been seen in this corner.
	// Objects that uses the UserStore directly the
	// responsability to validate the parameters

	// Select user with username
	GetUser(string) (*types.User, error)

	// Select user by email
	GetUserByEmail(string) (*types.User, error)

	GetUserByID(string) (*types.User, error)

	// Insert user with fullname,username,email,password,accountType
	CreateUser(string, string, string, string) error

	// Update user with username and proto user object (easier this way)
	UpdateUser(string, *types.User) error

	// List all users. carefull with this /!\
	ListUsers() ([]*types.User, error)
}

const (
	USER_STUDENT = "student"
	USER_SCHOOL  = "school"
)

// CreateUser .
func (s *store) CreateUser(email, username, password, accountType string) error {
	if accountType == "" || (accountType != USER_STUDENT && accountType != USER_SCHOOL) {
		accountType = USER_STUDENT
	}

	genUUID := uuid.NewV4()

	user := types.User{
		ID:       genUUID.String(),
		Email:    email,
		Username: username,
		Password: crypto.Sha256(password),
		Type:     accountType,
	}

	// insert user
	if _, err := s.Engine.Insert(&user); err != nil {
		return fmt.Errorf("inserting user: %v", err)
	}

	return nil
}

// GetUser select a single user by it username .
func (s *store) GetUserByUsername(username string) (*types.User, error) {
	user := &types.User{}

	// query user
	has, err := s.Engine.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}

	// no user exist
	if !has {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// GetUserByEmail select user by it email
func (s *store) GetUserByEmail(email string) (*types.User, error) {
	user := &types.User{}
	has, err := s.Engine.Where("email = ?", email).Get(user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// GetUserByEmail select user by it email
func (s *store) GetUserByID(id string) (*types.User, error) {
	user := &types.User{}
	has, err := s.Engine.Where("id = ?", id).Get(user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// alias for get user by username
func (s *store) GetUser(username string) (*types.User, error) {
	return s.GetUserByUsername(username)
}

// UpdateUser by email
func (s *store) UpdateUser(username string, user *types.User) error {
	_, err := s.Engine.Where("username = ?", username).Update(user)
	if err != nil {
		return err
	}

	return nil
}

// ListUsers .
func (s *store) ListUsers() ([]*types.User, error) {
	var users []*types.User
	err := s.Engine.Find(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
