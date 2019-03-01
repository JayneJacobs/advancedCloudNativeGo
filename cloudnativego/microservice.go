package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JayneJacobs/cloudnativego/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/hello", api.HelloHandleFunc)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/songs", api.SongsHandleFunc)
	http.HandleFunc("/api/songs/", api.SongsHandleFunc)
	http.ListenAndServe(port(), nil)

	fmt.Print("Something")

}
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to Cloud Native")
}
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
