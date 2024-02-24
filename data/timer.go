package data

import (
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
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

func SetTimerCmdFlags(cmd *cobra.Command) {

	// Local flag definitions
	cmd.Flags().IntP("duration", "d", 35, "The duration the timer runs in minutes")
	cmd.Flags().StringP("task", "t", "", "The name of the task")
	cmd.Flags().StringP("message", "m", "Done!", "A message to display when the timer is done")
}

func GetTimerFromFlags(cmd *cobra.Command) *Timer {

	minutes, _ := cmd.Flags().GetInt("duration")
	message, _ := cmd.Flags().GetString("message")
	task, _ := cmd.Flags().GetString("task")
	timer := NewTimer(task, minutes, message)

	return &timer
}

func GetFavoriteTimerFromFlags(cmd *cobra.Command) *FavoriteTimer {
	timer := GetTimerFromFlags(cmd)
	name, _ := cmd.Flags().GetString("name")
	id := uuid.New().ID()
	favoriteTimer := NewFavoriteTimer(id, name, *timer)
	return &favoriteTimer
}
