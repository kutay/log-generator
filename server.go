package main

import (
	"fmt"
	"net/http"

	"github.com/kutay/log-generator/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initServer(serverPort int) {
	docs.SwaggerInfo.Title = "LogGenerator API"
	docs.SwaggerInfo.Version = "0.0.2"

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/generate", generate)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
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

// @Summary Start log generation
// @Description starts log generation with the provided parameters
// @Param interval_ms query int false "interval in ms between each log" default(1000)
// @Param limit query int false "number of logs to generate" default(10)
// @Param message_length query int false "length of the log message field" default(100)
// @Param format query string false "format : plain / json" Enums(plain,json) default(plain)
// @Success 202
// @Failure 400
// @Router /generate [get]
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

// @Summary Get metrics
// @Description get metrics in Prometheus format
// @Success 200
// @Router /metrics [get]
func metrics() {
	// I don't know yet how to document external handlers
}
