package router

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type CorsTestRequest struct {
	Id             uuid.UUID
	Owner          uuid.UUID
	Origin         string
	ApplicationId  uuid.UUID
	Endpoint       string
	Methods        string
	Headers        string
	Authentication bool
}

type CorsTesDBInsert struct {
	Owner          uuid.UUID
	Origin         string
	Endpoint       string
	Methods        string
	Headers        string
	Authentication string
}

func (a Application) CreateCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	applicationId := params.ByName("applicationid")

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the application from the request
	corsTest := CorsTestRequest{}

	json.NewDecoder(r.Body).Decode(&corsTest)

	if corsTest.Origin == "" || corsTest.Endpoint == "" || corsTest.Methods == "" || corsTest.Headers == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	owner, ok := owner.(uuid.UUID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	corsTest.Owner = owner.(uuid.UUID)

	// Insert the application into the database
	conn := a.Db

	row := conn.QueryRow(r.Context(), "INSERT INTO cors_tests (owner, application_id, origin, endpoint, methods, headers, authentication) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", corsTest.Owner, applicationId, corsTest.Origin, corsTest.Endpoint, corsTest.Methods, corsTest.Headers, corsTest.Authentication)

	var result CorsTestRequest

	row.Scan(&result.Id, &result.Owner, &result.ApplicationId, &result.Origin, &result.Endpoint, &result.Methods, &result.Headers, &result.Authentication)

	json.NewEncoder(w).Encode(result)

}

func (a Application) EditCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	testid := params.ByName("id")

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the application from the request
	corsTest := CorsTestRequest{}

	json.NewDecoder(r.Body).Decode(&corsTest)

	if corsTest.Origin == "" || corsTest.Endpoint == "" || corsTest.Methods == "" || corsTest.Headers == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	owner, ok := owner.(uuid.UUID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	corsTest.Owner = owner.(uuid.UUID)

	// Insert the application into the database
	conn := a.Db

	row := conn.QueryRow(r.Context(), "UPDATE cors_tests SET origin = $1, endpoint = $2, methods = $3, headers = $4, authentication = $5 WHERE id = $6 RETURNING *", corsTest.Origin, corsTest.Endpoint, corsTest.Methods, corsTest.Headers, corsTest.Authentication, testid)

	var result CorsTestRequest

	err = row.Scan(&result.Id, &result.Owner, &result.Origin, &result.ApplicationId, &result.Endpoint, &result.Methods, &result.Headers, &result.Authentication)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func (a Application) DeleteCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	testid := params.ByName("id")

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var owneruuid uuid.UUID
	owneruuid, ok := owner.(uuid.UUID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the application into the database
	conn := a.Db

	_, err = conn.Exec(r.Context(), "DELETE FROM cors_tests WHERE id = $1 AND owner = $2", testid, owneruuid)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (a Application) GetCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
}

func (a Application) RunCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	//Dispatch Results to the queue at a later stage
}

func (a Application) ListCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (a Application) GetCorsTestResults(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
