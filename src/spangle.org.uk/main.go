package main

import (
	"log"
	"net/http"
	"spangle.org.uk/questions"
)

func main() {

	router := questions.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
