package api

import (
	"fmt"
	"net/http"

	authv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/service/auth/v1"
	userv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/service/user/v1"
)

func TestHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"test":"world"}`)
}

func New(auth *authv1.Service, user *userv1.Service) *Api {
	return &Api{
		auth: auth,
		user: user,
	}
}

type Api struct {
	auth *authv1.Service
	user *userv1.Service
}
