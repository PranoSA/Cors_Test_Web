package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type ApplicationTest struct {
	Id           uuid.UUID
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

	varownerstring := owner.Subject

	application.Owner, err = uuid.Parse(varownerstring)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	application.Time_Created = time.Now()
	application.Time_Edited = time.Now()

	// Insert the application into the database
	conn := a.Db

	//_, err = conn.Exec(r.Context(), "INSERT INTO cors_test_applications (owner, time_created, time_edited, name, description) VALUES ($1, $2, $3, $4, $5)", application.Owner, application.Time_Created, application.Time_Edited, application.Name, application.Description)

	row := conn.QueryRow(r.Context(), "INSERT INTO cors_test_applications (owner, time_created, time_edited, name, description) VALUES ($1, $2, $3, $4, $5) RETURNING id, time_created, time_edited", application.Owner, application.Time_Created, application.Time_Edited, application.Name, application.Description)

	err = row.Scan(&application.Id, &application.Time_Created, &application.Time_Edited)

	if err != nil {
		log.Default().Print(err)
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

	application.Owner, err = uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	application.Time_Edited = time.Now()

	// Insert the application into the database
	conn := a.Db

	//row, err := conn.Exec(r.Context(), "UPDATE cors_test_applications SET time_edited = $1, name = $2, description = $3 WHERE owner = $4 AND WHERE id = $5 ", application.Time_Edited, application.Name, application.Description, application.Owner, id)

	row := conn.QueryRow(r.Context(), "UPDATE cors_test_applications SET time_edited = $1, name = $2, description = $3 WHERE owner = $4 AND id = $5 RETURNING time_edited", application.Time_Edited, application.Name, application.Description, application.Owner, id)

	err = row.Scan(&application.Time_Edited)

	if err != nil {
		log.Default().Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	application.Id, err = uuid.Parse(id)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(application)

}

func (a Application) DeleteCorsApplication(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	/*iduuid, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Insert the application into the database
	conn := a.Db

	owneruuid, err := uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	row := conn.QueryRow(r.Context(), "DELETE FROM cors_test_applications WHERE owner = $1 AND id = $2 RETURNING id", owneruuid, id)
	var result uuid.UUID

	err = row.Scan(&result)

	if err != nil {
		log.Default().Print(err)
		log.Default().Print(id)
		log.Default().Print(owneruuid)
		w.WriteHeader(http.StatusNotFound)
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

	owneruuid, err := uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	row := conn.QueryRow(r.Context(), "SELECT owner, time_created, time_edited, name, description FROM cors_test_applications WHERE owner = $1 AND id = $2 ", owneruuid, id)

	if err != nil {
		log.Default().Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	application := ApplicationTest{}

	err = row.Scan(&application.Owner, &application.Time_Created, &application.Time_Edited, &application.Name, &application.Description)

	if err != nil {
		log.Default().Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	application.Id, err = uuid.Parse(id)

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

	owner_string := owner.Subject
	owneruuid, err := uuid.Parse(owner_string)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert the application into the database
	conn := a.Db

	rows, err := conn.Query(r.Context(), "SELECT id, owner, time_created, time_edited, name, description FROM cors_test_applications WHERE owner = $1 ", owneruuid)

	if err != nil {
		log.Default().Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	applications := []ApplicationTest{}

	for rows.Next() {
		application := ApplicationTest{}
		err = rows.Scan(&application.Id, &application.Owner, &application.Time_Created, &application.Time_Edited, &application.Name, &application.Description)
		applications = append(applications, application)
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(applications)

}
