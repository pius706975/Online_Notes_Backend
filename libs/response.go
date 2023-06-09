package libs

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code        int         `json:"-"`
	Status      string      `json:"status"`
	IsError     bool        `json:"isError"`
	Data        interface{} `json:"data,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (res Response) Send(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error When Encode Response"))
	}
}

func Respond(data interface{}, code int, isError bool) *Response {

	if isError {
		return &Response{
			Code:        code,
			Status:      getStatus(code),
			IsError:     isError,
			Description: data,
		}

	}
	return &Response{
		Code:    code,
		Status:  getStatus(code),
		IsError: isError,
		Data:    data,
	}
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 404:
		desc = "Not found"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
