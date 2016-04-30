package rets

import (
	"encoding/json"
	"log"
	"net/http"
)

// BasicPayload represents a simple code/message HTTP-like payload.
type BasicPayload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandlerError prints a HTTP 500 page and prints a given error in the
// application log.
func HandlerError(writer http.ResponseWriter, err error) {
	log.Println(err)
	HandlerHTTPInternalError(writer)
}

// HandlerHTTPOK prints a HTTP 200 page.
func HandlerHTTPOK(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{200, "success"})
}

// HandlerHTTPBadRequest prints a HTTP 400 page.
func HandlerHTTPBadRequest(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{400, "invalid request"})
}

// HandlerHTTPUnauthorized prints a HTTP 401 page.
func HandlerHTTPUnauthorized(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{401, "unauthorized"})
}

// HandlerHTTPNotFound prints a HTTP 404 page.
func HandlerHTTPNotFound(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{404, "resource not found"})
}

// HandlerHTTPInternalError prints a HTTP 500 page.
func HandlerHTTPInternalError(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{500, "internal server error"})
}

// HandlerBasic prints a HTTP page with the code and message from a given
// BasicPayload, and its JSON representation.
func HandlerBasic(writer http.ResponseWriter, payload BasicPayload) {
	data, err := json.Marshal(&payload)
	if err != nil {
		log.Println("JSON Marshal error: " + err.Error())
		writer.WriteHeader(500)
		return
	}
	if payload.Code >= 1000 {
		writer.WriteHeader(400)
	} else {
		writer.WriteHeader(payload.Code)
	}
	writer.Write(data)
}
