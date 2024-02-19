package data

import (
	"encoding/json"
	"os"
	"time"
)

var timerHistoryFileName = ".pomodoro_history.json"
var favoritesFileName = ".pomodoro_history.json"

type Timer struct {
	Task      string `json:"task"`
	Duration  int    `json:"duration"`
	Message   string `json:"message"`
	StartTime string `json:"start_time"`
}

type FavoriteTimer struct {
	Id    int    `json:"id"`
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

func NewFavoriteTimer(id int, name string, timer Timer) FavoriteTimer {

	return FavoriteTimer{
		Id:    id,
		Name:  name,
		Timer: timer,
	}
}

func AppendTimer(timer Timer) {

	file, err := os.OpenFile(timerHistoryFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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

	file, err := os.OpenFile(timerHistoryFileName, os.O_RDONLY, 0644)
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

func WriteFavorites(favorites []FavoriteTimer) {

	file, err := os.OpenFile(favoritesFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(favorites)
	if err != nil {
		panic(err)
	}
}

func ReadFavorites() []FavoriteTimer {

	file, err := os.OpenFile(favoritesFileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var favorites []FavoriteTimer
	for {
		var favorite FavoriteTimer
		if err := decoder.Decode(&favorite); err != nil {
			break
		}
		favorites = append(favorites, favorite)
	}

	return favorites
}
