

type Application = {
    Id: string;
    Owner : string;
    Time_Created: string;
    Time_Edited: string;
    Name: string;
    Description: string;
}

export type {
    Application,
    CorsTests,
    CorsTestResults
}


/*
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
*/

type CorsTests = {
    Id: string;
    Owner: string;
    Origin: string;
    ApplicationId: string;
    Endpoint: string;
    Methods: string;
    Headers: string;
    Authentication: boolean;
    Created_at: string;
}

/*
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
*/

type CorsTestResults = {
    Id: string;
    Owner: string;
    ApplicationId: string;
    TestId: string;
    Simple: boolean;
    Origin: string;
    Endpoint: string;
    Method: string;
    Header: string;
    Authentication: boolean;
    Okay: boolean;
    Errors: string[];
    Return_Access_Control_Allow_Origin: string;
    Return_Access_Control_Allow_Method: string;
    Return_Access_Control_Allow_Headers: string;
    Return_Access_Control_Max_Age: string;
    Return_Access_Control_Allow_Credentials: string;
    Return_Access_Control_Expose_Header: string;
    Time_Generated: string;
}