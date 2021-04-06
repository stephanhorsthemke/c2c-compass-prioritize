// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stephanhorsthemke/c2c-compass-prioritize/decoder"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Start HTTP server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTION")

	// Allow only for Frontend
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "This is the C2C compass backend, go to <link to frontend> in order to use the compass")
	case "POST":
		{
			prioritize(w, r)
		}
	case "OPTION":
		fmt.Fprintf(w, "")
	default:
		fmt.Fprintf(w, "Unsupported http method, use GET or POST")
	}

}

func prioritize(w http.ResponseWriter, r *http.Request) {

	var questions decoder.Questions = decoder.GetQuestions(r)
	fmt.Fprintf(w, "knowledge type %T, value: %v \n", questions.Knowledge, questions.Knowledge)
	fmt.Printf("knowledge: %v, position: %v  \n", questions.Knowledge, questions.Position)
}
