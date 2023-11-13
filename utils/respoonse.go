package utils

type Response struct{
	StatusCode int
	Message string
	Data interface{}
	Error error
}

func HttpResponse(statuscode int,message string,Data interface{},err error) Response{

	return Response{
		StatusCode: statuscode,
		Message: message,
		Data: Data,
		Error: err,
	}
		
	}
