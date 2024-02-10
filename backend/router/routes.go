package router

import (
	"net/http"

	tokenauthentication "github.com/PranoSA/Cors_Test_Web/backend/token_authentication"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type Application struct {
	Auth   tokenauthentication.TokenAuthentication
	Db     *pgxpool.Conn
	Router httprouter.Router
}

var Router httprouter.Router = httprouter.Router{}

// HealthCheck is a handler for the health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func (a *Application) StartServer() {
	a.Router.GET("/api/v1/application", a.GetCorsTestApplication)
	a.Router.GET("/api/v1/application/:id", a.GetCorsTestApplication)
	a.Router.POST("/api/v1/application", a.CreateCorsTestApplication)
	a.Router.PUT("/api/v1/application/:id", a.EditCorsTestApplication)
	a.Router.DELETE("/api/v1/application/:id", a.DeleteCorsApplication)

	a.Router.GET("/api/v1/application/:id/test", a.GetCorsTest)
	a.Router.POST("/api/v1/application/:id/test", a.CreateCorsTest)
	a.Router.PUT("/api/v1/application/:id/test/:test_id", a.EditCorsTest)
	a.Router.DELETE("/api/v1/application/:id/test/:test_id", a.DeleteCorsTest)

	a.Router.GET("/api/v1/application/:id/test/:test_id/result", a.GetCorsTestResults)
	a.Router.POST("/api/v1/application/:id/test/:test_id/result", a.CreateCorsTest)
}
