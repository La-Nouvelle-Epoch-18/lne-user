package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	apiv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/api/v1"
)

func NewServer(api *apiv1.Api, addr string) *http.Server {
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r.HandleFunc("/v1", apiv1.TestHandle).Methods("GET")
	r.HandleFunc("/v1/user", api.HandleGetUser).Methods("GET")
	r.HandleFunc("/v1/users", api.HandleGetUsers).Methods("GET")

	r.HandleFunc("/v1/auth/login", api.HandleAuthUser).Methods("POST")
	r.HandleFunc("/v1/auth/logout", apiv1.TestHandle).Methods("POST")
	r.HandleFunc("/v1/auth/verify", api.HandleVerifyToken).Methods("POST")
	r.HandleFunc("/v1/auth/signup", api.HandleSignUp).Methods("POST")

	r.HandleFunc("/v1/readme", api.HandleGetReadme).Methods("GET")

	return &http.Server{
		Handler:      handlers.CORS(originsOk, headersOk, methodsOk)(r),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
