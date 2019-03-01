package api

import (
	"fmt"
	"net/http"
)

// HelloHandleFunc is here
func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello Cloud Native Go.")
}
