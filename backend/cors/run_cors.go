package cors

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Should I define The Struct Here ....

type CorsTestRequest struct {
	Id                                      uuid.UUID
	Owner                                   uuid.UUID
	ApplicationId                           uuid.UUID
	TestId                                  uuid.UUID
	Simple                                  bool
	Origin                                  string
	Endpoint                                string
	Method                                  string
	Header                                  string
	Authentication                          bool
	Okay                                    bool
	Errors                                  []string
	Return_Access_Control_Allow_Origin      string
	Return_Access_Control_Allow_Method      string
	Return_Access_Control_Allow_Headers     string
	Return_Access_Control_Max_Age           string
	Return_Access_Control_Allow_Credentials string
	Return_Access_Control_Expose_Header     string

	Time_Generated time.Time
}

/*type CorsTestRequest struct {
	Owner          uuid.UUID
	ApplicationId  uuid.UUID
	Origin         string
	Endpoint       string
	Method         string
	Header         string
	Authentication string
}*/

/*

Accept, Accept Language, Content-Language, Content-Type(refer below for more conditions )
*/

func isSimple(method string, header string, authentication bool) bool {
	// How Should I Check If The Method Is Simple

	//What Methods Are Simple
	if method != "GET" && method != "HEAD" && method != "POST" {
		return false
	}

	if authentication {
		return false
	}

	// What Headers Are Simple
	if header != "" {
		return false
	}

	return true
}

func RunCorsTest(ctr *CorsTestRequest) error {
	//How Should I Run The Test

	if isSimple(ctr.Method, ctr.Header, ctr.Authentication) {
		ctr.Simple = true
	} else {
		ctr.Simple = false
	}

	if ctr.Simple {
		//
		req, err := http.NewRequest("GET", ctr.Endpoint, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Origin", ctr.Origin)
		req.Header.Set("Access-Control-Request-Method", ctr.Method)
		req.Header.Set("Access-Control-Request-Headers", ctr.Header)
		req.Method = ctr.Method

		client := http.Client{}
		client.Do(req)
		if err != nil {
			return err
		}

		// Get The Results
		allowedOrigin := req.Header.Get("Access-Control-Allow-Origin")
		//If The Origin Exists
		if allowedOrigin != "" {
			ctr.Return_Access_Control_Allow_Origin = allowedOrigin
		}
		return nil
	}

	req, err := http.NewRequest("OPTIONS", ctr.Endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Origin", ctr.Origin)
	req.Header.Set("Access-Control-Request-Method", ctr.Method)
	req.Header.Set("Access-Control-Request-Headers", ctr.Header)
	req.Method = "OPTIONS"

	// Run the test

	// Just Make The HTTP Requests
	client := http.Client{}
	if err != nil {
		return err
	}

	req.Header.Set("Origin", ctr.Origin)
	log.Default().Println(ctr.Origin)
	log.Default().Println(ctr.Endpoint)
	req.Header.Set("Access-Control-Request-Method", ctr.Method)
	req.Header.Set("Access-Control-Request-Headers", ctr.Header)

	if ctr.Authentication {
		// Add The Authentication Header
		//req.Header.Set("Access-Control-Request-Authorization")
	}
	req.Method = "OPTIONS"

	resp, err := client.Do(req)
	if err != nil {
		log.Default().Println(ctr.Endpoint)
		log.Default().Println("Bad Connection to " + ctr.Endpoint + err.Error())
		return err
	}

	// Get The Results

	allowedHeaders := resp.Header.Get("Access-Control-Allow-Headers")
	//If The Header Exists
	if allowedHeaders != "" {
		ctr.Return_Access_Control_Allow_Headers = allowedHeaders
	}

	allowedMethods := resp.Header.Get("Access-Control-Allow-Methods")
	allowedmethods := strings.Split(allowedMethods, ",")

	log.Default().Println("allowed methods :" + allowedMethods)

	//If The Method Exists
	for _, method := range allowedmethods {
		if strings.TrimSpace(method) == ctr.Method {
			ctr.Return_Access_Control_Allow_Method = allowedMethods
			break
		}
	}

	//method := strings.TrimSpace(ctr.Return_Access_Control_Allow_Method)

	var seeString []byte = []byte(ctr.Return_Access_Control_Allow_Method)

	log.Default().Println(seeString)

	log.Default().Println("allowed methods :" + ctr.Return_Access_Control_Allow_Method)

	log.Default().Println(ctr.Return_Access_Control_Allow_Method == "")

	if ctr.Return_Access_Control_Allow_Method == "" {
		log.Default().Println("The Method Is Not Allowed")
		ctr.Okay = false
		ctr.Errors = append(ctr.Errors, "The Method Is Not Allowed")
	}

	allowedOriginBool := resp.Header.Values("Access-Control-Allow-Origin")
	if allowedOriginBool == nil {
		ctr.Okay = false
		ctr.Errors = append(ctr.Errors, "The Origin Is Not Allowed")
	}

	allowedOrigin := resp.Header.Get("Access-Control-Allow-Origin")

	//If The Origin Exists
	if allowedOrigin != "" {
		ctr.Return_Access_Control_Allow_Origin = allowedOrigin
	}

	if ctr.Return_Access_Control_Allow_Origin != "*" && ctr.Return_Access_Control_Allow_Origin != ctr.Origin {
		ctr.Okay = false
		ctr.Errors = append(ctr.Errors, "The Origin Is Not Allowed")
	}

	if ctr.Return_Access_Control_Allow_Origin != "*" && ctr.Return_Access_Control_Allow_Origin != ctr.Origin {
		ctr.Okay = false
		ctr.Errors = append(ctr.Errors, "The Origin Is Not Allowed")
	}

	// SHould I check for "*" or the actual origin

	allowedCredentials := resp.Header.Get("Access-Control-Allow-Credentials")
	//If The Credentials Exists
	if allowedCredentials != "" {
		ctr.Return_Access_Control_Allow_Credentials = allowedCredentials
	}

	exposeHeader := resp.Header.Get("Access-Control-Expose-Headers")
	//If The Header Exists
	if exposeHeader != "" {
		ctr.Return_Access_Control_Expose_Header = exposeHeader
	}

	// Run the test
	// ...
	// Return the results
	return nil
}
