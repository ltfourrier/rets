package rets

import (
	"encoding/json"
	"log"
	"net/http"
)

type BasicPayload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandlerError(writer http.ResponseWriter, err error) {
	log.Println(err)
	HandlerHTTPInternalError(writer)
}

func HandlerHTTPOK(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{200, "success"})
}

func HandlerHTTPBadRequest(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{400, "invalid request"})
}

func HandlerHTTPUnauthorized(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{401, "unauthorized"})
}

func HandlerHTTPNotFound(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{404, "resource not found"})
}

func HandlerHTTPInternalError(writer http.ResponseWriter) {
	HandlerBasic(writer, BasicPayload{500, "internal server error"})
}

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
