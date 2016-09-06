package questions

import (
	"net/http"
	"spangle.org.uk/questions/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func QuestionIndex(w http.ResponseWriter, r *http.Request) {
	questionList := questions.Questions{
		questions.Question{Id: 1, Text: "Test question?"},
		questions.Question{Id: 2, Text: "Testing again?"},
	}

	json.NewEncoder(w).Encode(questionList)
}

func QuestionShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questId := vars["questId"]
	fmt.Fprintln(w, "Question show:", questId)
}
