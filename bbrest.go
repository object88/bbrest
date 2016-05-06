package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("127.0.0.1:27017/")

	if err != nil {
		fmt.Printf("Failed to create session to database: %s", err)
		panic(err)
	}

	fmt.Printf("Connected to mongo database\n")
	return s
}

func main() {
	fmt.Printf("Starting server...\n")
	r := mux.NewRouter()

	s := getSession()

	pC := NewPhotoController(s)
	pR := PhotoRouter{}
	pR.AddRouter(r, pC)

	http.ListenAndServe(":8080", r)
}
