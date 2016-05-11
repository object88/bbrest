package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/object88/bbrest/controllers"
	"github.com/object88/bbrest/handlers"
	"github.com/zenazn/goji/web"
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
	s := getSession()

	mux := web.New()

	pC := controllers.NewPhotoController(s, DatabaseName)
	pH := handlers.CreatePhotoHandler(pC)
	pH.AddPhotoHandler(mux)

	handlers.AddMiscellaneousHandler(mux)

	http.ListenAndServe(":1337", mux)
}
