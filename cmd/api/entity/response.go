package entity

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func CreateResponse(message string, value interface{}, err error) Response {
	var response Response

	response.Message = message
	if err != nil {
		response.Error = value
	} else {
		response.Data = value
	}

	return response
}
