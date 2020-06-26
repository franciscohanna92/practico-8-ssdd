package server

import (
	"encoding/json"
	"net/http"

	user "../../pkg"
	"github.com/gorilla/mux"
)

type api struct {
	router     http.Handler
	repository user.UserRepository
}

type Server interface {
	Router() http.Handler
}

func New(repo user.UserRepository) Server {
	a := &api{repository: repo}

	r := mux.NewRouter()
	r.HandleFunc("/users", a.fetchUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{Id:[a-zA-Z0-9_]+}", a.fetchUser).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := a.repository.FetchUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (a *api) fetchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := a.repository.FetchUserById(vars["Id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("User Not found")
		return
	}

	json.NewEncoder(w).Encode(user)
}
