package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type EventLog struct {
	Datetime string
	Message  string
	Severity string
}

var (
	logGenerated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "log_generator_generated_log_total",
		Help: "The total number of generated logs",
	})
)

func generateLogs(interval int, limit int, format string, messageLength int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < limit; i++ {
		logString := generateLog(format, messageLength, limit, interval)
		fmt.Println(logString)
		logGenerated.Inc()

		if interval > 0 {
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
	}
}

func generateLog(format string, messageLength int, nbLogs int, interval int) string {
	now := time.Now().Format(time.RFC3339Nano)

	if format == "plain" {
		return fmt.Sprintf("%s - INFO - %s", now, randSeq(messageLength))
	} else if format == "json" {
		msg, err := json.Marshal(EventLog{now, randSeq(messageLength), "INFO"})
		if err != nil {
			fmt.Println(err)
		}

		return string(msg)
	} else {
		panic("unhandled format. Only accepts plain and json.")
	}
}
