package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type EventLog struct {
	Datetime string
	Message  string
}

func generateLogs(interval int, limit int, format string, messageLength int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < limit; i++ {
		now := time.Now().Format(time.RFC3339Nano)

		if format == "plain" {
			fmt.Println(now, randSeq(messageLength))
		} else if format == "json" {
			msg, err := json.Marshal(EventLog{now, randSeq(messageLength)})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(msg))
		}

		if interval > 0 {
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
	}
}
