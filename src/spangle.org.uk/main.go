package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"spangle.org.uk/questions"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", questions.Index)
	router.HandleFunc("/questions", questions.QuestionIndex)
	router.HandleFunc("/questions/{questId}", questions.QuestionShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
