package web

import (
	"fmt"
	"log"
	"net/http"
)

func Connect() http.ResponseWriter {
	var webInterface http.ResponseWriter
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		webInterface = w
		fmt.Fprintf(w, "Va te faire enculer Ylango")
	})
	log.Fatal(http.ListenAndServe(":8000", nil))

	return webInterface
}
