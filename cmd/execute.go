package cmd

import (
	"fmt"
	"github.com/dipesh23-apt/golang_api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Execute() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/user/fetch", controller.GetallUsers).Methods("POST")
	r.HandleFunc("/api/v1/user/create", controller.CreateUser).Methods("POST")

	r.HandleFunc("/api/v1/user/{id}", controller.DeleteUser).Methods("DELETE")
	http.Handle("/", r)
	fmt.Println("Running on server 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
