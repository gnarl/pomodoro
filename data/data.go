package data

import (
	"encoding/json"
	"os"
	"time"
)

var dataFileName = ".pomodoro.json"

type Timer struct {
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Message   string `json:"message"`
	StartTime string `json:"start_time"`
}

func NewTimer(name string, duration int, message string) Timer {

	startTime := time.Now().Format(time.RFC3339)

	return Timer{
		Name:      name,
		Duration:  duration,
		Message:   message,
		StartTime: startTime,
	}
}

func AppendTimer(timer Timer) {

	file, err := os.OpenFile(dataFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(timer)
	if err != nil {
		panic(err)
	}
}

func ReadTimers() []Timer {

	file, err := os.OpenFile(dataFileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var timers []Timer
	for {
		var timer Timer
		if err := decoder.Decode(&timer); err != nil {
			break
		}
		timers = append(timers, timer)
	}

	return timers
}
