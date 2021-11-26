package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	api "github.com/paulbobba/nwb/api"
	// sw "./TradeSkills/ApiCodeGen"
)

func main() {
	log.Printf("Server started")

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
