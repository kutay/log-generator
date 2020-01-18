package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/generate", generate)

	http.Handle("/metrics", promhttp.Handler())

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

	go func() {
		generateLogs(interval, limit, format, messageLength)
	}()

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Generating %d logs of length %d every %dms in %s format", limit, messageLength, interval, format)

}
