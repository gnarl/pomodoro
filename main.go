/*
Copyright © 2024 Chuck Fouts

*/
package main

import (
	//"github.com/gnarl/pomodoro/cmd"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gnarl/pomodoro/internal/tui"
	log "github.com/gnarl/pomodoro/internal/utils"
)

func main() {

	log.Logger.Debug("Starting pomodoro")
	//cmd.Execute()

	model := tui.NewModel()

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
