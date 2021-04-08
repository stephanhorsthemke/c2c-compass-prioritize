// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stephanhorsthemke/c2c-compass-prioritize/decoder"
	"github.com/stephanhorsthemke/c2c-compass-prioritize/encoder"
	"github.com/stephanhorsthemke/c2c-compass-prioritize/gsheet"
	"github.com/stephanhorsthemke/c2c-compass-prioritize/prioritizer"
)

func init() {
	gsheet.UpdateLinks()
}

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
			handlePost(w, r)
		}
	case "OPTION":
		fmt.Fprintf(w, "")
	default:
		fmt.Fprintf(w, "Unsupported http method, use GET or POST")
	}

}

func handlePost(w http.ResponseWriter, r *http.Request) {

	questions := decoder.GetQuestions(r)
	fmt.Printf("knowledge: %v, position: %v  \n", questions.Knowledge, questions.Position)
	links := gsheet.GetLinks()
	resultLinks := prioritizer.Prioritize(questions, links)
	resultJson := encoder.LinksToJson(resultLinks)
	fmt.Fprintf(w, "%s", resultJson)

}
