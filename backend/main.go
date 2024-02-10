package main

import (
	"log"
	"net/http"

	"github.com/PranoSA/Cors_Test_Web/backend/router"
	tokenauthentication "github.com/PranoSA/Cors_Test_Web/backend/token_authentication"
)

func main() {

	// Start the Server with router.Router

	var app router.Application = router.Application{
		Router: router.Router,
		Auth:   tokenauthentication.TestCookieAuth{},
	}

	app.StartServer()

	server := http.Server{
		Addr:    ":8080",
		Handler: &app.Router,
	}

	log.Fatal(server.ListenAndServe())

}
