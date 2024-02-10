package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	cors "github.com/PranoSA/Cors_Test_Web/backend/cors"
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
	Created_at     time.Time
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

	applicationId := params.ByName("id")

	_, err := uuid.Parse(applicationId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the application from the request
	corsTest := CorsTestRequest{}

	json.NewDecoder(r.Body).Decode(&corsTest)

	if corsTest.Origin == "" || corsTest.Endpoint == "" || corsTest.Methods == "" {
		log.Default().Print("Invalid Request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	corsTest.Owner, err = uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	// Insert the application into the database
	conn := a.Db

	row := conn.QueryRow(r.Context(), `INSERT INTO cors_tests (owner, application_id, origin, endpoint, methods, headers, authentication) VALUES ($1, $2, $3, $4, $5, $6, $7)
	 RETURNING id, owner, application_id, origin, endpoint, methods, headers, authentication, created_at`,
		corsTest.Owner, applicationId, corsTest.Origin, corsTest.Endpoint, corsTest.Methods, corsTest.Headers, corsTest.Authentication)

	var result CorsTestRequest

	err = row.Scan(&result.Id, &result.Owner, &result.ApplicationId, &result.Origin, &result.Endpoint, &result.Methods, &result.Headers, &result.Authentication, &result.Created_at)

	if err != nil {
		log.Default().Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	if corsTest.Origin == "" || corsTest.Endpoint == "" || corsTest.Methods == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	corsTest.Owner, err = uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Insert the application into the database
	conn := a.Db

	time_now := time.Now()

	row := conn.QueryRow(r.Context(), "UPDATE cors_tests SET origin = $1, endpoint = $2, methods = $3, headers = $4, authentication = $5, created_at = $6 WHERE id = $7 RETURNING created_at", corsTest.Origin, corsTest.Endpoint, corsTest.Methods, corsTest.Headers, corsTest.Authentication, time_now, testid)

	//var result CorsTestRequest

	err = row.Scan(&corsTest.Created_at)

	if err != nil {
		log.Default().Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	corsTest.Id = uuid.MustParse(testid)

	json.NewEncoder(w).Encode(corsTest)

}

func (a Application) DeleteCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	testid := params.ByName("id")

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	owneruuid, err := uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	id := params.ByName("id")

	_, err := uuid.Parse(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	owneruuid, err := uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get The Test from the database
	conn := a.Db

	row := conn.QueryRow(r.Context(), `SELECT id, owner, application_id, origin, endpoint, methods, headers, authentication, created_at FROM cors_tests WHERE id = $1 AND owner = $2`, id, owneruuid)

	var result CorsTestRequest

	err = row.Scan(&result.Id, &result.Owner, &result.ApplicationId, &result.Origin, &result.Endpoint, &result.Methods, &result.Headers, &result.Authentication, &result.Created_at)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func (a Application) RunCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	applicationid := params.ByName("id")

	_, err := uuid.Parse(applicationid)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	owneruuid, err := uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// First, Find ALl The Tests THat are Part of the Application
	// Then, Run Each Test

	// Get The Test from the database
	conn := a.Db

	rows, err := conn.Query(r.Context(), `SELECT id, owner, application_id, origin, endpoint, methods, headers, authentication FROM cors_tests WHERE application_id = $1 AND owner = $2`, applicationid, owneruuid)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Now Scan Into tests
	var tests []cors.CorsTestRequest = make([]cors.CorsTestRequest, 0)

	for rows.Next() {
		var corsTest cors.CorsTestRequest
		err = rows.Scan(&corsTest.Id, &corsTest.Owner, &corsTest.ApplicationId, &corsTest.Origin, &corsTest.Endpoint, &corsTest.Method, &corsTest.Header, &corsTest.Authentication)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Default().Print(corsTest.Id)
		tests = append(tests, corsTest)
	}

	log.Default().Print(tests)

	var genTests []cors.CorsTestRequest = make([]cors.CorsTestRequest, 0)

	// For Each Test, Run The Test
	for _, test := range tests {

		// For
		for _, method := range strings.Split(test.Method, ",") {
			//First Test Case With No Headers
			var test_result cors.CorsTestRequest = cors.CorsTestRequest{
				Owner:          owneruuid,
				ApplicationId:  uuid.MustParse(applicationid),
				Origin:         test.Origin,
				Endpoint:       test.Endpoint,
				Method:         method,
				Header:         "",
				Authentication: false,
				TestId:         test.Id,
			}
			genTests = append(genTests, test_result)

			for _, header := range strings.Split(test.Header, ",") {
				//Rest of the Test Cases With Headers
				var test cors.CorsTestRequest = cors.CorsTestRequest{
					Owner:          owneruuid,
					ApplicationId:  uuid.MustParse(applicationid),
					Origin:         test.Origin,
					Endpoint:       test.Endpoint,
					Method:         method,
					Header:         header,
					Authentication: false,
					TestId:         test.Id,
				}
				genTests = append(genTests, test)
				log.Default().Print("test id" + test.Id.String())

			}
		}

	}

	var testResults []cors.CorsTestRequest = make([]cors.CorsTestRequest, 0)

	for _, test := range genTests {
		//Run The Test
		var next_result cors.CorsTestRequest
		next_result.Errors = make([]string, 0)
		next_result.TestId = test.TestId
		_, err = test.RunCorsTest(&next_result)

		next_result.ApplicationId, err = uuid.Parse(applicationid)
		if err != nil {
			next_result.Okay = false
			next_result.Errors = append(next_result.Errors, err.Error())
		}
		testResults = append(testResults, next_result)
	}

	// Now Insert Into Database The Results
	for i, result := range testResults {
		// Insert the application into the database
		conn := a.Db
		log.Default().Print(result.ApplicationId)
		log.Default().Print(result.TestId)
		row := conn.QueryRow(r.Context(), `INSERT INTO cors_test_results 
		(owner, application_id, test_id, origin, endpoint, method, header, authentication, okay, errors, return_access_control_allow_origin, 
			return_access_control_allow_method, return_access_control_allow_headers, return_access_control_max_age, return_access_control_allow_credentials,
			 return_access_control_expose_header, time_generated, simple) 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING id`,
			result.Owner, result.ApplicationId, result.TestId, result.Origin, result.Endpoint, result.Method, result.Header,
			result.Authentication, result.Okay, result.Errors, result.Return_Access_Control_Allow_Origin, result.Return_Access_Control_Allow_Method,
			result.Return_Access_Control_Allow_Headers, result.Return_Access_Control_Max_Age, result.Return_Access_Control_Allow_Credentials,
			result.Return_Access_Control_Expose_Header, result.Time_Generated, result.Simple)

		next_id := uuid.UUID{}
		err = row.Scan(&next_id)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		genTests[i].Id = next_id
	}

	// Json Marshall All The Newly Created Tests
	json.NewEncoder(w).Encode(genTests)

	//Dispatch Results to the queue at a later stage
}

func (a Application) ListCorsTest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	_, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	owneruuid, err := uuid.Parse(owner.Subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Insert the application into the database
	conn := a.Db

	//row := conn.QueryRow(r.Context(), "SELECT * FROM cors_tests WHERE owner = $1 AND application_id = $2", owneruuid, id)

	rows, err := conn.Query(r.Context(), `SELECT id, owner, application_id, origin, endpoint, methods, headers, authentication
	FROM cors_tests 
	WHERE owner = $1 AND application_id = $2`, owneruuid, id)

	var result []CorsTestRequest = make([]CorsTestRequest, 0)

	for rows.Next() {
		var corsTest CorsTestRequest
		err = rows.Scan(&corsTest.Id, &corsTest.Owner, &corsTest.ApplicationId, &corsTest.Origin, &corsTest.Endpoint, &corsTest.Methods, &corsTest.Headers, &corsTest.Authentication)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result = append(result, corsTest)
	}

	json.NewEncoder(w).Encode(result)

}

func (a Application) GetCorsTestResults(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// This Will Be for the Entire Application For Now

	appid := params.ByName("id")

	_, err := uuid.Parse(appid)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the owner from the token
	owner, err := a.Auth.AuthenticateToken(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	owneruuid, err := uuid.Parse(owner.Subject)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get The Test from the database
	conn := a.Db

	var result []cors.CorsTestRequest = make([]cors.CorsTestRequest, 0)

	rows, err := conn.Query(r.Context(), `SELECT id, owner, application_id, origin, endpoint, methods, headers, authentication FROM cors_tests WHERE application_id = $1 AND owner = $2`, appid, owneruuid)

	for rows.Next() {
		var corsTest cors.CorsTestRequest
		err = rows.Scan(&corsTest.Id, &corsTest.Owner, &corsTest.ApplicationId, &corsTest.Origin, &corsTest.Endpoint, &corsTest.Method, &corsTest.Header, &corsTest.Authentication)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result = append(result, corsTest)
	}

	json.NewEncoder(w).Encode(result)

	/*rows, err := conn.Query(r.Context(), `SELECT id, owner, application_id, origin, endpoint, methods, headers, authentication, created_at FROM cors_tests WHERE application_id = $1 AND owner = $2`, appid, owneruuid)

	var result []CorsTestRequest = make([]CorsTestRequest, 0)

	for rows.Next() {
		var corsTest CorsTestRequest
		err = rows.Scan(&corsTest.Id, &corsTest.Owner, &corsTest.ApplicationId, &corsTest.Origin, &corsTest.Endpoint, &corsTest.Methods, &corsTest.Headers, &corsTest.Authentication, &corsTest.Created_at)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result = append(result, corsTest)
	}
	*/
}
