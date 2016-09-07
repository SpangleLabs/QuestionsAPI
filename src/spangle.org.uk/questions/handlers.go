package questions

import (
	"net/http"
	"spangle.org.uk/questions/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/jinzhu/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func QuestionIndex(w http.ResponseWriter, r *http.Request) {
	questionList := questions.Questions{
		questions.Question{Text: "Test question?"},
		questions.Question{Text: "Testing again?"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(questionList); err != nil {
		panic(err)
	}
}

func QuestionShow(w http.ResponseWriter, r *http.Request) {
	// Get input variables
	vars := mux.Vars(r)
	questId := vars["questId"]

	// Connect to database
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database.")
	}
	defer db.Close()

	// Get question
	var question questions.Question
	db.First(&question, questId)

	// Format and write output
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(question); err != nil {
		panic(err)
	}
}
