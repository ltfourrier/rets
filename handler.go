package rets

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorPayload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandlerHTTPBadRequest(writer http.ResponseWriter, request *http.Request) {
	payload := ErrorPayload{400, "invalid request"}
	HandlerHTTPError(writer, payload)
}

func HandlerHTTPNotFound(writer http.ResponseWriter, request *http.Request) {
	payload := ErrorPayload{404, "resource not found"}
	HandlerHTTPError(writer, payload)
}

func HandlerHTTPError(writer http.ResponseWriter, payload ErrorPayload) {
	writer.WriteHeader(payload.Code)
	data, err := json.Marshal(&payload)
	if err != nil {
		log.Println("JSON Marshal error: " + err.Error())
		return
	}
	writer.Write(data)
}
