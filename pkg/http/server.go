package server

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	apiv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/api/v1"
)

var (
	corsHeaders = "Content-Type,Accept,Authorization"
	corsMethods = "GET,POST,DELETE"
)

// corsHandler
func corsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", corsHeaders)
	w.Header().Set("Access-Control-Allow-Methods", corsMethods)
}

func httpWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)

			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				corsHandler(w, r)

				return
			}
		}

		h.ServeHTTP(w, r)
	})
}

func NewServer(api *apiv1.Api, addr string) *http.Server {
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"*"})

	r.Use(handlers.CORS(originsOk, headersOk, methodsOk))

	r.HandleFunc("/v1", apiv1.TestHandle).Methods("GET")
	r.HandleFunc("/v1/user", api.HandleGetUser).Methods("GET")
	r.HandleFunc("/v1/users", api.HandleGetUsers).Methods("GET")

	r.HandleFunc("/v1/auth/login", api.HandleAuthUser).Methods("POST")
	r.HandleFunc("/v1/auth/logout", apiv1.TestHandle).Methods("POST")
	r.HandleFunc("/v1/auth/verify", api.HandleVerifyToken).Methods("POST")
	r.HandleFunc("/v1/auth/signup", api.HandleSignUp).Methods("POST")

	r.HandleFunc("/v1/readme", api.HandleGetReadme).Methods("GET", "OPTIONS")

	return &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
