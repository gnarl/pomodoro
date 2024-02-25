package common

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/gnarl/pomodoro/internal/data"
)

func SetTimerCmdFlags(cmd *cobra.Command) {

	// Local flag definitions
	cmd.Flags().IntP("duration", "d", 35, "The duration the timer runs in minutes")
	cmd.Flags().StringP("task", "t", "", "The name of the task")
	cmd.Flags().StringP("message", "m", "Done!", "A message to display when the timer is done")
}

func GetTimerFromFlags(cmd *cobra.Command) *data.Timer {

	minutes, _ := cmd.Flags().GetInt("duration")
	message, _ := cmd.Flags().GetString("message")
	task, _ := cmd.Flags().GetString("task")
	timer := data.NewTimer(task, minutes, message)

	return &timer
}

func GetFavoriteTimerFromFlags(cmd *cobra.Command) *data.FavoriteTimer {

	timer := GetTimerFromFlags(cmd)
	name, _ := cmd.Flags().GetString("name")
	id := uuid.New().ID()
	favoriteTimer := data.NewFavoriteTimer(id, name, *timer)
	return &favoriteTimer
}
