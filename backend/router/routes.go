package router

import (
	"log"
	"net/http"

	tokenauthentication "github.com/PranoSA/Cors_Test_Web/backend/token_authentication"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type Application struct {
	Auth   tokenauthentication.TokenAuthentication
	Db     *pgxpool.Pool
	Router httprouter.Router
}

var Router httprouter.Router = httprouter.Router{}

// HealthCheck is a handler for the health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func OptionsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Default().Println("OPTIONS request")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusNoContent)
}

func (a *Application) StartServer() {
	a.Router.OPTIONS("/*path", OptionsHandler)
	//a.Router.OPTIONS("/api/v1/*path", OptionsHandler)
	a.Router.GET("/api/v1/application", CorsMiddlware(a.ListCorsTestApplications))
	a.Router.GET("/api/v1/application/:id", CorsMiddlware(a.GetCorsTestApplication))
	a.Router.POST("/api/v1/application", CorsMiddlware(a.CreateCorsTestApplication))
	a.Router.PUT("/api/v1/application/:id", CorsMiddlware(a.EditCorsTestApplication))
	a.Router.DELETE("/api/v1/application/:id", CorsMiddlware(a.DeleteCorsApplication))

	a.Router.GET("/api/v1/tests/:id", CorsMiddlware(a.ListCorsTest))
	a.Router.POST("/api/v1/test/:id", CorsMiddlware(a.CreateCorsTest))
	a.Router.PUT("/api/v1/test/:id", CorsMiddlware(a.EditCorsTest)) //id in this instance is the testid
	a.Router.DELETE("/api/v1/test/:id", CorsMiddlware(a.DeleteCorsTest))
	a.Router.GET("/api/v1/test/:id", CorsMiddlware(a.GetCorsTest))

	//Now Generate The Tests for the application

	a.Router.GET("/api/v1/results/:id", CorsMiddlware(a.GetCorsTestResults)) // Don't Really Need a Singular Version
	a.Router.POST("/api/v1/result/:id", CorsMiddlware(a.RunCorsTest))

	/*a.Router.PUT("/api/v1/application/{id}", a.EditCorsTestApplication)
	a.Router.DELETE("/api/v1/application/{id}", a.DeleteCorsApplication)

	a.Router.GET("/api/v1/application/{id}/test", a.GetCorsTest)
	a.Router.POST("/api/v1/application/{id}/test", a.CreateCorsTest)
	a.Router.PUT("/api/v1/application/{id}/test/{test_id}", a.EditCorsTest)
	a.Router.DELETE("/api/v1/application/{id}/test/{test_id}", a.DeleteCorsTest)

	a.Router.GET("/api/v1/application/{id}/test/{test_id}/result", a.GetCorsTestResults)
	a.Router.POST("/api/v1/application/{id}/test/{test_id}/result", a.CreateCorsTest)
	*/
}
