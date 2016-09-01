package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/questions", QuestionIndex)
	router.HandleFunc("/questions/{questId}", QuestionShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func QuestionIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Question list!")
}

func QuestionShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["questId"]
	fmt.Fprintln(w, "Question show:", todoId)
}