package defs

type Error struct {
	Error string	`json:"error"`
	ErrorCode string	`json:"errorcode"`
}

type ErrorResponse struct {
	HttpSC int
	Error	Error
}

var(
	ErrorRequestBodyParseFailed=ErrorResponse{HttpSC:400,Error:Error{"Requst body invaild","001"}}
	ErrorAuthFailed=ErrorResponse{HttpSC:401,Error:Error{"Auth failed","002"}}
	ErrorDBError=ErrorResponse{HttpSC:500,Error:Error{"DB Auth Failed","003"}}
	ErrorHttpInternal=ErrorResponse{
		HttpSC: 500,
		Error:  Error{"Server Internal service error","004"},
	}
	)