// Command proof serves the rail proof: mockup frame and generated widget side by side.
package main

import (
	"log"
	"net/http"
	"os"

	"carrel.dev/typecase/internal/proof"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}
	mux, err := proof.Mux(".")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("typecase proof on http://localhost:%s/compare", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
