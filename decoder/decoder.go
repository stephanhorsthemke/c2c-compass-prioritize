package decoder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Questions struct {
	Knowledge int    `json:"knowledge"`
	Position  string `json:"position"`
}

func GetQuestions(r *http.Request) Questions {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Could not read body")
	}

	var questions Questions
	err = json.Unmarshal(body, &questions)
	if err != nil {
		log.Fatalf("unmarshaling of body did not work: %s", err)
	}

	return questions
}
