package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/PranoSA/Cors_Test_Web/backend/router"
	tokenauthentication "github.com/PranoSA/Cors_Test_Web/backend/token_authentication"
	"github.com/jackc/pgx/v5/pgxpool"
)

var port int
var dbUser string
var dbPassword string
var dbName string
var dbHost string
var dbPort int

func main() {
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.StringVar(&dbUser, "dbUser", "cors_tester", "Database User")
	flag.StringVar(&dbPassword, "dbPassword", "test", "Database Password")
	flag.StringVar(&dbName, "dbName", "cors_test", "Database Name")
	flag.StringVar(&dbHost, "dbHost", "localhost", "Database Host")
	flag.IntVar(&dbPort, "dbPort", 5432, "Database Port")
	flag.Parse()

	var db_connection string = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Create A New Connection Pool
	conn, err := pgxpool.New(context.Background(), db_connection)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Connected to Database, Successful Ping\n")

	defer conn.Close()

	// Start the Server with router.Router

	var app router.Application = router.Application{
		Router: router.Router,
		Auth:   tokenauthentication.TestCookieAuth{},
		Db:     conn,
	}

	app.StartServer()

	server := http.Server{
		Addr:    ":8080",
		Handler: &app.Router,
	}

	log.Fatal(server.ListenAndServe())

}
