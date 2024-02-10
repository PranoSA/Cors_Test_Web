package corstest

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Should I define The Struct Here ....

type CorsTestRequest struct {
	Owner                                   uuid.UUID
	ApplicationId                           uuid.UUID
	ResultsId                               uuid.UUID
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

func (c CorsTestRequest) RunCorsTest(ctr *CorsTestRequest) (*CorsTestRequest, error) {
	//How Should I Run The Test

	if isSimple(ctr.Method, ctr.Header, ctr.Authentication) {
		c.Simple = true
	} else {
		c.Simple = false
	}

	if c.Simple {
		//
		req, err := http.NewRequest("GET", ctr.Endpoint, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Origin", ctr.Origin)
		req.Header.Set("Access-Control-Request-Method", ctr.Method)
		req.Header.Set("Access-Control-Request-Headers", ctr.Header)
		req.Method = ctr.Method

		client := http.Client{}
		client.Do(req)
		if err != nil {
			return nil, err
		}

		// Get The Results
		allowedOrigin := req.Header.Get("Access-Control-Allow-Origin")
		//If The Origin Exists
		if allowedOrigin != "" {
			c.Return_Access_Control_Allow_Origin = allowedOrigin
		}
		return &c, nil
	}

	req, err := http.NewRequest("OPTIONS", ctr.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Origin", ctr.Origin)
	req.Header.Set("Access-Control-Request-Method", ctr.Method)
	req.Header.Set("Access-Control-Request-Headers", ctr.Header)

	// Run the test

	// Just Make The HTTP Requests
	client := http.Client{}
	if err != nil {
		return nil, err
	}

	req.Header.Set("Origin", ctr.Origin)
	req.Header.Set("Access-Control-Request-Method", ctr.Method)
	req.Header.Set("Access-Control-Request-Headers", ctr.Header)
	req.Method = "OPTIONS"

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Get The Results

	allowedHeaders := resp.Header.Get("Access-Control-Allow-Headers")
	//If The Header Exists
	if allowedHeaders != "" {
		c.Return_Access_Control_Allow_Headers = allowedHeaders
	}

	allowedMethods := resp.Header.Get("Access-Control-Allow-Methods")
	//If The Method Exists
	if allowedMethods != "" {
		c.Return_Access_Control_Allow_Method = allowedMethods
	}

	allowedOrigin := resp.Header.Get("Access-Control-Allow-Origin")

	//If The Origin Exists
	if allowedOrigin != "" {
		c.Return_Access_Control_Allow_Origin = allowedOrigin
	}

	// SHould I check for "*" or the actual origin

	allowedCredentials := resp.Header.Get("Access-Control-Allow-Credentials")
	//If The Credentials Exists
	if allowedCredentials != "" {
		c.Return_Access_Control_Allow_Credentials = allowedCredentials
	}

	exposeHeader := resp.Header.Get("Access-Control-Expose-Headers")
	//If The Header Exists
	if exposeHeader != "" {
		c.Return_Access_Control_Expose_Header = exposeHeader
	}

	// Run the test
	// ...
	// Return the results
	return &c, nil
}
