package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/object88/bbrest/controllers"
	"github.com/object88/bbrest/handlers"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("127.0.0.1:27017/")

	if err != nil {
		fmt.Printf("Failed to create session to database: %s\n", err)
		panic(err)
	}

	fmt.Printf("Connected to mongo database.\n")
	return s
}

func main() {
	fmt.Printf("Starting server...\n")
	r := mux.NewRouter()

	s := getSession()

	pC := controllers.NewPhotoController(s, DatabaseName)
	handlers.AddPhotoHandler(pC, r)

	http.ListenAndServe(":8080", r)
}
