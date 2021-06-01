package controller

import (
	"encoding/json"
	"fmt"
	"github.com/dipesh23-apt/golang_api/models"
	"github.com/dipesh23-apt/golang_api/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "apllication/json")
		var b models.User
		param := mux.Vars(r)
		b, err := repo.GetUserfromDB(param["id"])
		if err != nil {
			fmt.Fprintf(w, "Invalid ID,does not exist")
			return
		}
		json.NewEncoder(w).Encode(b)
	}
}
func GetallUsers() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var x models.Body
		json.NewDecoder(r.Body).Decode(&x)
		c, err := repo.GetallUsersfromDB(x.Id)
		if err != nil {
			fmt.Fprintf(w, "Some id data cannot be found !!")
			return
		}
		json.NewEncoder(w).Encode(c)
	}
}

func CreateUser() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var us models.User
		err := json.NewDecoder(r.Body).Decode(&us)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		d, err := repo.CreateUserinDB(us)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(d)
	}
}
func DeleteUser() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		param := mux.Vars(r)
		err := repo.DeleteUserfromDB(param["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Successfully deleted the user with id %s", param["id"])
	}
}
