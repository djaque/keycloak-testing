package server

import (
	"encoding/json"
	"net/http"

	gopher "github.com/djaque/keycloak-testing/rest-sample/pkg"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const BEARER_SCHEMA = "Bearer "
const API_KEY = "ff965bc8-c0b3-47a8-a1c8-4eec7b950e05"

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
	r.HandleFunc("/gophers/{ID}", a.fetchGopher).Methods(http.MethodGet)
	r.HandleFunc("/gophers/{ID}", a.checkGopher).Methods(http.MethodPost)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) verifyHeader(w http.ResponseWriter, r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Error("Authorization missing")
		return false
	}
	token := authHeader[len(BEARER_SCHEMA):]
	if token == "" {
		log.Error("Token missing")
		return false
	}

	if token != API_KEY {
		log.Error("Token invalid")
		return false
	}

	log.Info("Token OK")
	return true
}

func (a *api) fetchGophers(w http.ResponseWriter, r *http.Request) {
	log.Println("get all gophers")

	if !a.verifyHeader(w, r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	gophers, _ := a.repository.FetchGophers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gophers)
}

func (a *api) fetchGopher(w http.ResponseWriter, r *http.Request) {
	if !a.verifyHeader(w, r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Info("Search for one gopher")
	vars := mux.Vars(r)
	log.Printf("Gopher search for %s", vars["ID"])
	gopher, err := a.repository.FetchGopherByName(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("Gopher %s not found", vars["ID"])
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}

	log.Info("Gopher %s found", vars["ID"])
	json.NewEncoder(w).Encode(gopher)
}

type pass struct {
	Password string `json:"password"`
}

func (a *api) checkGopher(w http.ResponseWriter, r *http.Request) {
	if !a.verifyHeader(w, r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	log.Printf("Gopher for pswd search for %s", vars["ID"])
	gopher, err := a.repository.FetchGopherByName(vars["ID"])
	if err != nil {
		log.Printf("Gopher %s not found", vars["ID"])
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var input pass

	err = decoder.Decode(&input)
	if err != nil {
		log.Printf("input body invalid %+v", r.Body)
		w.WriteHeader(http.StatusBadRequest) // We use not found for simplicity
		json.NewEncoder(w).Encode("input not valid")
		return
	}

	log.Info("Gopher passtocheck:", input.Password)

	if input.Password != gopher.Password {
		log.Println("Password missmatch")
		w.WriteHeader(http.StatusBadRequest) // We use not found for simplicity
		json.NewEncoder(w).Encode("Password invalid")
		return

	}
	log.Println("Password success")
	r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gopher)
}
