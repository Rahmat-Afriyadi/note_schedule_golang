package helper


//Response is used for static shape json return
type Response struct {
	Code  int64        `json:"code"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(code int64, data interface{}) Response {
	res := Response{
		Code: code,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(code int64, data interface{}) Response {
	res := Response{
		Code: code,
		Data:    data,
	}
	return res
}
