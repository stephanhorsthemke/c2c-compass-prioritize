// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stephanhorsthemke/c2c-compass-prioritize/gsheet"
)

func main() {

	gsheet.GetGoogleSheet()

	log.Print("starting server...")
	/*http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	*/

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello !\n")
}
