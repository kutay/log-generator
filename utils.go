package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
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

func parseRequestParamString(req *http.Request, paramName string, defaultValue string) string {
	reqValue := req.URL.Query().Get(paramName)
	if reqValue != "" {
		return reqValue
	}

	return defaultValue
}
