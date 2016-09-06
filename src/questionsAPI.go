package main

import (
	"fmt"
	"log"
	"net/http"
	"spangle.org.uk"

	"github.com/gorilla/mux"
	"encoding/json"
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
	questions := spangle_org_uk.Questions{
		spangle_org_uk.Question{Id: 1, Text: "Test question?"},
		spangle_org_uk.Question{Id: 2, Text: "Testing again?"},
	}

	json.NewEncoder(w).Encode(questions)
}

func QuestionShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questId := vars["questId"]
	fmt.Fprintln(w, "Question show:", questId)
}