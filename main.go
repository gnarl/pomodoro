/*
Copyright © 2024 Chuck Fouts

*/
package main

import (
	"github.com/gnarl/pomodoro/cmd"
	"github.com/gnarl/pomodoro/utils"

	"log/slog"
)

func main() {
	utils.InitJsonLogger(slog.LevelDebug)
	log := utils.GetLogger()
	log.Debug("Starting pomodoro")

	cmd.Execute()
}
