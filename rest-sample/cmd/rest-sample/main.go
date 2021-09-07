package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/djaque/compose-keycloak/rest-sample/pkg/storage/inmem"

	sample "github.com/djaque/compose-keycloak/rest-sample/cmd/sample-data"
	gopher "github.com/djaque/compose-keycloak/rest-sample/pkg"

	"github.com/djaque/compose-keycloak/rest-sample/pkg/server"
)

func main() {
	withData := flag.Bool("withData", false, "initialize the api with some gophers")
	flag.Parse()

	var gophers map[string]*gopher.Gopher
	if *withData {
		gophers = sample.Gophers
	}

	repo := inmem.NewGopherRepository(gophers)
	s := server.New(repo)

	fmt.Println("The gopher server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
