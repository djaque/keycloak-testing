package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	gopher "github.com/djaque/keycloak-testing/rest-sample/pkg"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/Nerzal/gocloak/v9"
)

const (
	BEARER_SCHEMA       = "Bearer "
	API_KEY             = "ff965bc8-c0b3-47a8-a1c8-4eec7b950e05"
	KEY_CLOAK_TOKEN_URL = "http://keycloak:8080/auth/realms/master/protocol/openid-connect/token"
	KEY_CLOAK_ADMIN_URL = "http://keycloak:8080/auth/admin/realms/master/users"
	KEY_CLOAK_USER_INFO = "http://keycloak:8080/auth/realms/master/protocol/openid-connect/userinfo"
	ADMIN_USERNAME      = "admin"
	ADMIN_PASS          = "password"
	ADMIN_CLIENT_ID     = "admin-cli"
	CLIENT_ID           = "sample-app"
	REALM               = "master"
)

type api struct {
	router     http.Handler
	repository gopher.GopherRepository
	client     gocloak.GoCloak
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

	r.HandleFunc("/api/register", a.ApiRegister).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/login", a.ApiLogin).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/userinfo", a.ApiUserInfo).Methods(http.MethodGet, http.MethodOptions)

	client := gocloak.NewClient("http://keycloak:8080")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, ADMIN_USERNAME, ADMIN_PASS, REALM)
	if err != nil {
		panic("Something wrong with the credentials or url")
	}
	log.Println(token)
	a.client = client
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

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (a *api) ApiRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("Register user")

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var input CreateUser

	err := decoder.Decode(&input)
	if err != nil {
		log.Printf("input body invalid %+v", r.Body)
		w.WriteHeader(http.StatusBadRequest) // We use not found for simplicity
		json.NewEncoder(w).Encode("input not valid")
		return
	}

	log.Printf("%+v", a.registerKeycloak(input))

	log.Println("Create Success")
	r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(input)

}

type token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

func (a *api) getTokenKeycloak(username, password, client_id string) token {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("grant_type", "password")
	data.Add("client_id", client_id)
	encodedData := data.Encode()
	log.Println(encodedData)
	req, err := http.NewRequest("POST", KEY_CLOAK_TOKEN_URL, strings.NewReader(encodedData))
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
	}
	defer response.Body.Close()

	/*
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
	*/
	//Read the response body
	decoder := json.NewDecoder(response.Body)
	var input token
	err = decoder.Decode(&input)
	if err != nil {
		log.Printf("input body token invalid %+v", response.Body)
	}
	log.Printf("Login info: %+v", input)
	return input
}

/*
{
    "email": "danyjaqueherrera@gmail.com",
    "emailVerified": true,
    "firstName": "Dny",
    "lastName": "Jax",
    "enabled": true,
 	"credentials":[{
 		"type": "password",
 		"value": "1234567890",
 		"temporary": false
 	}]
}
*/
type Credential struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary bool   `json:"temporary"`
}
type KeyCloakUser struct {
	Email         string       `json:"email"`
	EmailVerified bool         `json:"emailVerified"`
	FirstName     string       `json:"firstName"`
	LastName      string       `json:"lastName"`
	Enabled       bool         `json:"enabled"`
	Credentials   []Credential `json:"credentials"`
}

func (a *api) registerKeycloak(user CreateUser) string {
	//convert CreateUser to KeyCloakUser
	keycloakUser := &KeyCloakUser{
		Email:         user.Email,
		EmailVerified: true,
		Enabled:       true,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Credentials: []Credential{
			{
				Type:      "password",
				Value:     user.Password,
				Temporary: false,
			},
		},
	}
	//Encode the data
	postBody, _ := json.Marshal(keycloakUser)
	responseBody := bytes.NewBuffer(postBody)
	newToken := a.getTokenKeycloak(ADMIN_USERNAME, ADMIN_PASS, ADMIN_CLIENT_ID)
	//Leverage Go's HTTP Post function to make request
	req, _ := http.NewRequest("POST", KEY_CLOAK_ADMIN_URL, responseBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+newToken.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)

	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

	return "OK"
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (a *api) ApiLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Login user")
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var input LoginUser

	err := decoder.Decode(&input)
	if err != nil {
		log.Printf("input body invalid %+v", r.Body)
		w.WriteHeader(http.StatusBadRequest) // We use not found for simplicity
		json.NewEncoder(w).Encode("input not valid")
		return
	}

	access := a.getTokenKeycloak(input.Email, input.Password, CLIENT_ID)
	log.Println("Token:" + access.AccessToken)

	log.Println("Login Success")
	r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(access)

}

type keycloakUserInfo struct {
	Sub               string `json:"sub"`
	EmailVerified     bool   `json:"email_verified"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	Email             string `json:"email"`
}

func (a *api) ApiUserInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("UserInfo user")
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Error("Authorization missing")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	token := authHeader[len(BEARER_SCHEMA):]
	if token == "" {
		log.Error("Token missing")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Println("Token:" + token)
	ctx := context.Background()
	userInfo, err := a.client.GetUserInfo(ctx, token, REALM)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("User info: %+v", userInfo)
	w.Header().Set("Content-Type", "application/json")

	log.Println("UserInfo Success")
	r.Body.Close()
	json.NewEncoder(w).Encode(userInfo)
}
