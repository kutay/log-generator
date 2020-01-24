package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	port := 8090
	envPort := os.Getenv("PORT")

	if envPort != "" {
		i, err := strconv.Atoi(envPort)
		if err == nil {
			port = i
		}
	}

	fmt.Println("Starting log-generator on port", port)

	initServer(port)
}
