package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	userv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/service/user/v1"
)

func (a *Api) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req *userv1.CreateUserRequest
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
		return
	}

	err = a.user.CreateUser(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"success":true}`)
}

func (a *Api) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	token, err := getTokenFromHeader(header)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"error":"%s"}`, header)
		return
	}

	user, err := a.user.GetUser(token)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
		return
	}

	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (a *Api) HandleGetReadme(w http.ResponseWriter, r *http.Request) {

}
