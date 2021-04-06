// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	s
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Start HTTP server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Allow only for Frontend
	switch r.Method {
	case "GET":
		fmt.Println("This is the C2C compass backend, go to link in order to use the compass")
	case "POST":
		{
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			prioritize(body)
		}
	}
}

func prioritize(body []byte) {
	fmt.Println(body)
}
