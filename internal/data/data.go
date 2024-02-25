package data

import (
	"encoding/json"
	"errors"
	"os"
)

const timerHistoryFileName = ".pomodoro_history.json"
const favoritesFileName = ".pomodoro_favorites.json"

func fileExists(filePath string) (bool, error) {

	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

func AppendTimer(timer *Timer) {

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

	exists, err := fileExists(timerHistoryFileName)
	if err != nil {
		panic(err)
	}
	if !exists {
		return []Timer{}
	}

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

	for _, fav := range favorites {
		err = encoder.Encode(fav)
		if err != nil {
			panic(err)
		}
	}
}

func ReadFavorites() []FavoriteTimer {

	exists, err := fileExists(favoritesFileName)
	if err != nil {
		panic(err)
	}
	if !exists {
		return []FavoriteTimer{}
	}

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
