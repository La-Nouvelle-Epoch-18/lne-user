package api

import (
	"fmt"
	"net/http"
)

func (a *Api) HandleAuthUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"test":"world"}`)
}
