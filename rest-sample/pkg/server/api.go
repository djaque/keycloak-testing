package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	gopher "github.com/djaque/compose-keycloak/rest-sample/pkg"

	"github.com/gorilla/mux"
)

type api struct {
	router     http.Handler
	repository gopher.GopherRepository
}

type Server interface {
	Router() http.Handler
}

func New(repo gopher.GopherRepository) Server {
	a := &api{repository: repo}

	r := mux.NewRouter()
	r.HandleFunc("/gophers", a.fetchGophers).Methods(http.MethodGet)
	r.HandleFunc("/gophers/{ID:[a-zA-Z0-9_]+}", a.fetchGopher).Methods(http.MethodGet)
	r.HandleFunc("/gophers/{ID:[a-zA-Z0-9_]+}", a.checkGopher).Methods(http.MethodPost)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchGophers(w http.ResponseWriter, r *http.Request) {
	gophers, _ := a.repository.FetchGophers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gophers)
}

func (a *api) fetchGopher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("Gopher search for %s \n", vars["ID"])
	gopher, err := a.repository.FetchGopherByName(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}

	fmt.Printf("Gopher %s found\n", vars["ID"])
	json.NewEncoder(w).Encode(gopher)
}

type pass struct {
	Password string `json:"password"`
}

func (a *api) checkGopher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("Gopher search for %s \n", vars["ID"])
	gopher, err := a.repository.FetchGopherByName(vars["ID"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var input pass

	err = decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // We use not found for simplicity
		json.NewEncoder(w).Encode("Password not found on input")
		return
	}

	fmt.Printf("Gopher passtocheck %s \n", input.Password)

	if input.Password != gopher.Password {
		w.WriteHeader(http.StatusBadRequest) // We use not found for simplicity
		json.NewEncoder(w).Encode("Password invalid")
		return

	}
	r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gopher)
}
