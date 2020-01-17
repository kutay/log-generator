package main

import (
	"fmt"
	"net/http"
)

func initServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/generate", generate)

	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func generate(w http.ResponseWriter, req *http.Request) {
	interval := parseRequestParamInt(req, "interval_ms", 1000)
	limit := parseRequestParamInt(req, "limit", 10)
	format := parseRequestParamString(req, "format", "plain")
	messageLength := parseRequestParamInt(req, "message_length", 100)

	generateLogs(interval, limit, format, messageLength)
}