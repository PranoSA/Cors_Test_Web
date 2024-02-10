/*
 
 type ApplicationTest struct {
 Owner        uuid.UUID
 Time_Created time.Time
 Time_Edited  time.Time
 Name         string
 Description  string
 }
 
 cors_test_applications 
 */
CREATE TABLE cors_test_applications(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner UUID NOT NULL,
    time_created TIMESTAMP NOT NULL,
    time_edited TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);
/*
 cors_tests -> Can Have Multiple Headers Seperated By A Comma
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
 
 
 row := conn.QueryRow(r.Context(), "INSERT INTO cors_tests (owner, application_id, origin, endpoint, methods, headers, authentication) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", corsTest.Owner, applicationId, corsTest.Origin, corsTest.Endpoint, corsTest.Methods, corsTest.Headers, corsTest.Authentication)
 
 
 */
CREATE TABLE cors_tests(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner UUID NOT NULL,
    application_id UUID NOT NULL,
    origin VARCHAR(255) NOT NULL,
    endpoint VARCHAR(255) NOT NULL,
    methods VARCHAR(255) NOT NULL,
    headers VARCHAR(255) NOT NULL,
    authentication BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (application_id) REFERENCES cors_test_applications(id)
);
/*
 
 Cors Test Results 
 
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
 s
 
 
 */
CREATE TABLE IF NOT EXISTS cors_test_results(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner UUID NOT NULL,
    application_id UUID NOT NULL,
    test_id UUID NOT NULL,
    simple BOOLEAN NOT NULL,
    origin VARCHAR(255) NOT NULL,
    endpoint VARCHAR(255) NOT NULL,
    method VARCHAR(255) NOT NULL,
    header VARCHAR(255) NOT NULL,
    authentication BOOLEAN NOT NULL,
    okay BOOLEAN NOT NULL,
    errors TEXT [] NOT NULL,
    return_access_control_allow_origin VARCHAR(255) NOT NULL,
    return_access_control_allow_method VARCHAR(255) NOT NULL,
    return_access_control_allow_headers VARCHAR(255) NOT NULL,
    return_access_control_max_age VARCHAR(255) NOT NULL,
    return_access_control_allow_credentials VARCHAR(255) NOT NULL,
    return_access_control_expose_header VARCHAR(255) NOT NULL,
    time_generated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (application_id) REFERENCES cors_test_applications(id),
    FOREIGN KEY (test_id) REFERENCES cors_tests(id)
);