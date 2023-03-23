package main

import (
	"log"
	"net/http"

	"github.com/FidelMonteiro04/web_form.go/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
