package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type ApplicationTest struct {
	Owner        uuid.UUID
	Time_Created time.Time
	Time_Edited  time.Time
	Name         string
	Description  string
}

func (a Application) CreateCorsTestApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the application from the request
	application := ApplicationTest{}
	// Decode the application from the request
	// ...

	json.NewDecoder(r.Body).Decode(&application)

	if application.Name == "" || application.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	application.Owner = owner.(uuid.UUID)
	application.Time_Created = time.Now()
	application.Time_Edited = time.Now()

	// Insert the application into the database
	conn := a.Db

	_, err = conn.Exec(r.Context(), "INSERT INTO cors_test_applications (owner, time_created, time_edited, name, description) VALUES ($1, $2, $3, $4, $5)", application.Owner, application.Time_Created, application.Time_Edited, application.Name, application.Description)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(application)

}

func (a Application) EditCorsTestApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the application from the request
	application := ApplicationTest{}
	// Decode the application from the request
	// ...

	json.NewDecoder(r.Body).Decode(&application)

	if application.Name == "" || application.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	application.Owner = owner.(uuid.UUID)
	application.Time_Edited = time.Now()

	// Insert the application into the database
	conn := a.Db

	row, err := conn.Exec(r.Context(), "UPDATE cors_test_applications SET time_edited = $1, name = $2, description = $3 WHERE owner = $4 AND WHERE id = $5 ", application.Time_Edited, application.Name, application.Description, application.Owner, id)

	if row.RowsAffected() == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(application)

}

func (a Application) DeleteCorsApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Insert the application into the database
	conn := a.Db

	row, err := conn.Exec(r.Context(), "DELETE FROM cors_test_applications WHERE owner = $1 AND WHERE id = $2 ", owner, id)

	if row.RowsAffected() == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (a Application) GetCorsTestApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Insert the application into the database
	conn := a.Db

	row := conn.QueryRow(r.Context(), "SELECT * FROM cors_test_applications WHERE owner = $1 AND WHERE id = $2 ", owner, id)

	application := ApplicationTest{}

	err = row.Scan(&application.Owner, &application.Time_Created, &application.Time_Edited, &application.Name, &application.Description)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(application)

}

func (a Application) ListCorsTestApplications(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Insert the application into the database
	conn := a.Db

	rows, err := conn.Query(r.Context(), "SELECT * FROM cors_test_applications WHERE owner = $1 ", owner)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	applications := []ApplicationTest{}

	for rows.Next() {
		application := ApplicationTest{}
		err = rows.Scan(&application.Owner, &application.Time_Created, &application.Time_Edited, &application.Name, &application.Description)
		applications = append(applications, application)
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(applications)

}
