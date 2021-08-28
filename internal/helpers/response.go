package helpers

type ResponseStruct struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func Response(message string, data interface{}, statusCode int) ResponseStruct {
	return ResponseStruct{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
