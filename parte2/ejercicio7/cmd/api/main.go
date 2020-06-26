package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"../../pkg/storage/inmem"

	user "../../pkg"
	sample "../sample-data"

	"../../pkg/server"
)

func main() {
	withData := flag.Bool("withData", false, "initialize the api with some users")
	flag.Parse()

	var users map[string]*user.User
	if *withData {
		users = sample.Users
	}

	repo := inmem.NewUserRepository(users)
	s := server.New(repo)

	fmt.Println("The user server is on tap now: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", s.Router()))
}
