package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	apiv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/api/v1"
)

func NewServer(api *apiv1.Api, addr string) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/v1", apiv1.TestHandle).Methods("GET")
	r.HandleFunc("/v1/user", apiv1.TestHandle).Methods("POST")
	r.HandleFunc("/v1/auth/login", api.HandleAuthUser).Methods("POST")
	r.HandleFunc("/v1/auth/logout", apiv1.TestHandle).Methods("POST")
	r.HandleFunc("/v1/auth/signup", api.HandleSignUp).Methods("POST")

	return &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
