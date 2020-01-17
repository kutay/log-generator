package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println("Starting log-generator on port 8090")

	generateLogs(100, 10)

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

func parseRequestParamInt(req *http.Request, paramName string, defaultValue int) int {
	reqValue := req.URL.Query().Get(paramName)
	if reqValue != "" {
		i, err := strconv.Atoi(reqValue)

		if err != nil {
			fmt.Println(err)
		} else {
			return i
		}
	}

	return defaultValue
}

func generate(w http.ResponseWriter, req *http.Request) {
	interval := parseRequestParamInt(req, "interval_ms", 1000)
	limit := parseRequestParamInt(req, "limit", 10)

	generateLogs(interval, limit)
}

func generateLogs(interval int, limit int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < limit; i++ {
		fmt.Println(randSeq(100))
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
