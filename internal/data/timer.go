package data

import (
	"time"
)

type Timer struct {
	Task      string `json:"task"`
	Duration  int    `json:"duration"`
	Message   string `json:"message"`
	StartTime string `json:"start_time"`
}

type FavoriteTimer struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Timer Timer  `json:"timer"`
}

func NewTimer(task string, duration int, message string) Timer {

	startTime := time.Now().Format(time.RFC3339)

	return Timer{
		Task:      task,
		Duration:  duration,
		Message:   message,
		StartTime: startTime,
	}
}

func NewFavoriteTimer(id uint32, name string, timer Timer) FavoriteTimer {

	return FavoriteTimer{
		Id:    id,
		Name:  name,
		Timer: timer,
	}
}
