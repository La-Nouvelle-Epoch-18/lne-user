package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	authv1 "github.com/La-Nouvelle-Epoch-18/lne-user/pkg/service/auth/v1"
)

func getTokenFromHeader(header string) (string, error) {
	elems := strings.Split(header, " ")
	if len(elems) != 2 {
		return "", fmt.Errorf("invalid token format")
	}
	if elems[0] != "bearer" {
		return "", fmt.Errorf("authorization header should start with 'bearer'")
	}
	return elems[1], nil
}

func (a *Api) HandleAuthUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req *authv1.LoginUserRequest
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
		return
	}

	resp, err := a.auth.LoginUser(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
		return
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
		return
	}
}

func (a *Api) HandleVerifyToken(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Authorization")
	token, err := getTokenFromHeader(header)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"error":"%s"}`, header)
		return
	}

	err = a.auth.VerifyToken(token)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"valid":"false","error":"%s"}`, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"valid":"true"}`)
}
