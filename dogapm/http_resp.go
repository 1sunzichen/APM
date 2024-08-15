package dogapm

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    any    `json:"body"`
}

type httpStatus struct {
}

var HttpStatus = &httpStatus{}

const (
	jsonContentType = "application/json"
)

func (h *httpStatus) Ok(w http.ResponseWriter) {
	status := &Status{
		Code:    http.StatusOK,
		Message: "success",
		Body:    nil,
	}
	data, _ := json.Marshal(status)
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (h *httpStatus) OkBody(w http.ResponseWriter, body any) {
	status := &Status{
		Code:    http.StatusOK,
		Message: "success",
		Body:    body,
	}
	data, _ := json.Marshal(status)
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *httpStatus) Fail(w http.ResponseWriter, message string, body any) {
	status := &Status{
		Code:    http.StatusBadRequest,
		Message: message,
		Body:    body,
	}
	data, _ := json.Marshal(status)
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)
}
func (h *httpStatus) Error(w http.ResponseWriter, message string, body any) {
	status := &Status{
		Code:    http.StatusInternalServerError,
		Message: message,
		Body:    body,
	}
	data, _ := json.Marshal(status)
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(data)
}
