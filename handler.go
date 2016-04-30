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
	HandlerBasicHTTP(writer, 500)
}

// HandlerBasicHTTP automatically prints a page from the given HTTP error code.
func HandlerBasicHTTP(writer http.ResponseWriter, code int) {
	messages := map[int]string{
		100: "Continue",
		101: "Switching Protocols",
		200: "OK",
		201: "Created",
		202: "Accepted",
		203: "Non-Authoritative Information",
		204: "No Content",
		205: "Reset Content",
		206: "Partial Content",
		300: "Multiple Choices",
		301: "Moved Permanently",
		302: "Found",
		303: "See Other",
		304: "Not Modified",
		305: "Use Proxy",
		306: "Switch Proxy",
		307: "Temporary Redirect",
		308: "Permanent Redirect",
		400: "Bad Request",
		401: "Unauthorized",
		402: "Payment Required",
		403: "Forbidden",
		404: "Not Found",
		405: "Method Not Allowed",
		406: "Not Acceptable",
		407: "Proxy Authentication Required",
		408: "Request Timeout",
		409: "Conflict",
		410: "Gone",
		411: "Length Required",
		412: "Precondition Failed",
		413: "Payload Too Large",
		414: "URI Too Long",
		415: "Unsupported Media Type",
		416: "Range Not Satisfiable",
		417: "Expectation Failed",
		421: "Misdirected Request",
		426: "Upgrade Required",
		428: "Precondition Required",
		429: "Too Many Requests",
		431: "Request Header Fields Too Large",
		451: "Unavailable For Legal Reasons",
		500: "Internal Server Error",
		501: "Not Implemented",
		502: "Bad Gateway",
		503: "Service Unavailable",
		504: "Gateway Timeout",
		505: "HTTP Version Not Supported",
		506: "Variant Also Negotiates",
		510: "Not Extended",
		511: "Network Authentication Required",
	}

	message, exists := messages[code]
	if exists {
		HandlerBasic(writer, BasicPayload{code, message})
	}
	HandlerBasic(writer, BasicPayload{501, "HTTP Code Not Implemented"})
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
